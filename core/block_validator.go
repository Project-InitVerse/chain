// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"PureChain/consensus"
	"PureChain/core/state"
	"PureChain/core/types"
	"PureChain/params"
	"PureChain/trie"
	"fmt"
)

// BlockValidator is responsible for validating block headers, uncles and
// processed state.
//
// BlockValidator implements Validator.
type BlockValidator struct {
	config *params.ChainConfig // Chain configuration options
	bc     *BlockChain         // Canonical block chain
	engine consensus.Engine    // Consensus engine used for validating
}

// NewBlockValidator returns a new block validator which is safe for re-use
func NewBlockValidator(config *params.ChainConfig, blockchain *BlockChain, engine consensus.Engine) *BlockValidator {
	validator := &BlockValidator{
		config: config,
		engine: engine,
		bc:     blockchain,
	}
	return validator
}

// ValidateBody validates the given block's uncles and verifies the block
// header's transaction and uncle roots. The headers are assumed to be already
// validated at this point.
func (v *BlockValidator) ValidateBody(block *types.Block) error {
	// Check whether the block's known, and if not, that it's linkable
	if v.bc.HasBlockAndState(block.Hash(), block.NumberU64()) {
		return ErrKnownBlock
	}
	// Header validity is known at this point, check the uncles and transactions
	header := block.Header()
	if err := v.engine.VerifyUncles(v.bc, block); err != nil {
		return err
	}
	if hash := types.CalcUncleHash(block.Uncles()); hash != header.UncleHash {
		return fmt.Errorf("uncle root hash mismatch: have %x, want %x", hash, header.UncleHash)
	}

	validateFuns := []func() error{
		func() error {
			if v.bc.HasBlockAndState(block.Hash(), block.NumberU64()) {
				return ErrKnownBlock
			}
			return nil
		},
		func() error {
			if hash := types.DeriveSha(block.Transactions(), trie.NewStackTrie(nil)); hash != header.TxHash {
				return fmt.Errorf("transaction root hash mismatch: have %x, want %x", hash, header.TxHash)
			}
			return nil
		},
		func() error {
			if !v.bc.HasBlockAndState(block.ParentHash(), block.NumberU64()-1) {
				if !v.bc.HasBlock(block.ParentHash(), block.NumberU64()-1) {
					return consensus.ErrUnknownAncestor
				}
				return consensus.ErrPrunedAncestor
			}
			return nil
		},
	}
	validateRes := make(chan error, len(validateFuns))
	for _, f := range validateFuns {
		tmpFunc := f
		go func() {
			validateRes <- tmpFunc()
		}()
	}
	for i := 0; i < len(validateFuns); i++ {
		r := <-validateRes
		if r != nil {
			return r
		}
	}
	return nil
}

// ValidateState validates the various changes that happen after a state
// transition, such as amount of used gas, the receipt roots and the state root
// itself. ValidateState returns a database batch if the validation was a success
// otherwise nil and an error is returned.
func (v *BlockValidator) ValidateState(block *types.Block, statedb *state.StateDB, receipts types.Receipts, usedGas uint64) error {
	header := block.Header()
	if block.GasUsed() != usedGas {
		return fmt.Errorf("invalid gas used (remote: %d local: %d)", block.GasUsed(), usedGas)
	}
	// Validate the received block's bloom with the one derived from the generated receipts.
	// For valid blocks this should always validate to true.
	validateFuns := []func() error{
		func() error {
			rbloom := types.CreateBloom(receipts)
			if rbloom != header.Bloom {
				return fmt.Errorf("invalid bloom (remote: %x  local: %x)", header.Bloom, rbloom)
			}
			return nil
		},
		func() error {
			receiptSha := types.DeriveSha(receipts, trie.NewStackTrie(nil))
			if receiptSha != header.ReceiptHash {
				return fmt.Errorf("invalid receipt root hash (remote: %x local: %x)", header.ReceiptHash, receiptSha)
			} else {
				return nil
			}
		},
		func() error {
			if root := statedb.IntermediateRoot(v.config.IsEIP158(header.Number)); header.Root != root {
				transaction_str := ""
				for _, oneTrx := range block.Transactions() {
					txJson, err := oneTrx.MarshalJSON()
					if err == nil {
						transaction_str += string(txJson) + "\n"
					} else {
						transaction_str += "marshal json failed tx hash " + string(oneTrx.Hash().String()) + "\n"
					}

				}
				//err_str := fmt.Errorf("invalid merkle root block number%v blockVal:%v transaction %v", header.Number.String(), header.Coinbase.String(), transaction_str)
				//log.Error("invalid merkle root block", "error", err_str)
				//statedb.IterativeDump(true, true, true, json.NewEncoder(os.Stdout))
				return fmt.Errorf("invalid merkle root (remote: %x local: %x)", header.Root, root)
			} else {
				return nil
			}
		},
	}
	validateRes := make(chan error, len(validateFuns))
	for _, f := range validateFuns {
		tmpFunc := f
		go func() {
			validateRes <- tmpFunc()
		}()
	}
	for i := 0; i < len(validateFuns); i++ {
		r := <-validateRes
		if r != nil {
			return r
		}
	}
	return nil
}

// CalcGasLimit computes the gas limit of the next block after parent. It aims
// to keep the baseline gas above the provided floor, and increase it towards the
// ceil if the blocks are full. If the ceil is exceeded, it will always decrease
// the gas allowance.
func CalcGasLimit(parent *types.Block, gasFloor, gasCeil uint64) uint64 {
	// contrib = (parentGasUsed * 3 / 2) / 256
	contrib := (parent.GasUsed() + parent.GasUsed()/2) / params.GasLimitBoundDivisor

	// decay = parentGasLimit / 256 -1
	decay := parent.GasLimit()/params.GasLimitBoundDivisor - 1

	/*
		strategy: gasLimit of block-to-mine is set based on parent's
		gasUsed value.  if parentGasUsed > parentGasLimit * (2/3) then we
		increase it, otherwise lower it (or leave it unchanged if it's right
		at that usage) the amount increased/decreased depends on how far away
		from parentGasLimit * (2/3) parentGasUsed is.
	*/
	limit := parent.GasLimit() - decay + contrib
	if limit < params.MinGasLimit {
		limit = params.MinGasLimit
	}
	// If we're outside our allowed gas range, we try to hone towards them
	if limit < gasFloor {
		limit = parent.GasLimit() + decay
		if limit > gasFloor {
			limit = gasFloor
		}
	} else if limit > gasCeil {
		limit = parent.GasLimit() - decay
		if limit < gasCeil {
			limit = gasCeil
		}
	}
	return limit
}
