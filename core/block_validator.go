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
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/Project-InitVerse/chain/accounts/abi"
	"github.com/Project-InitVerse/chain/common"
	"github.com/Project-InitVerse/chain/consensus"
	"github.com/Project-InitVerse/chain/consensus/inihash/systemcontract"
	"github.com/Project-InitVerse/chain/core/state"
	"github.com/Project-InitVerse/chain/core/types"
	"github.com/Project-InitVerse/chain/core/vm"
	"github.com/Project-InitVerse/chain/crypto"
	"github.com/Project-InitVerse/chain/params"
	"github.com/Project-InitVerse/chain/trie"
	"github.com/ethereum/go-ethereum/log"
)

type BlockValidatorOption func(*BlockValidator) *BlockValidator

func EnableRemoteVerifyManager(remoteValidator *remoteVerifyManager) BlockValidatorOption {
	return func(bv *BlockValidator) *BlockValidator {
		bv.remoteValidator = remoteValidator
		return bv
	}
}

// BlockValidator is responsible for validating block headers, uncles and
// processed state.
//
// BlockValidator implements Validator.
type BlockValidator struct {
	config          *params.ChainConfig // Chain configuration options
	bc              *BlockChain         // Canonical block chain
	remoteValidator *remoteVerifyManager
}

// NewBlockValidator returns a new block validator which is safe for re-use
func NewBlockValidator(config *params.ChainConfig, blockchain *BlockChain, opts ...BlockValidatorOption) *BlockValidator {
	validator := &BlockValidator{
		config: config,
		bc:     blockchain,
	}

	for _, opt := range opts {
		validator = opt(validator)
	}

	return validator
}

// ValidateBody validates the given block's uncles and verifies the block
// header's transaction and uncle roots. The headers are assumed to be already
// validated at this point.
func (v *BlockValidator) ValidateBody(block *types.Block) error {
	// Check whether the block is already imported.
	if v.bc.HasBlockAndState(block.Hash(), block.NumberU64()) {
		return ErrKnownBlock
	}
	// Header validity is known at this point. Here we verify that uncles, transactions
	// and withdrawals given in the block body match the header.
	header := block.Header()
	if err := v.bc.engine.VerifyUncles(v.bc, block); err != nil {
		return err
	}
	if hash := types.CalcUncleHash(block.Uncles()); hash != header.UncleHash {
		return fmt.Errorf("uncle root hash mismatch (header value %x, calculated %x)", header.UncleHash, hash)
	}

	validateFuns := []func() error{
		func() error {
			if hash := types.DeriveSha(block.Transactions(), trie.NewStackTrie(nil)); hash != header.TxHash {
				return fmt.Errorf("transaction root hash mismatch: have %x, want %x", hash, header.TxHash)
			}
			return nil
		},
		func() error {
			// Withdrawals are present after the Shanghai fork.
			if header.WithdrawalsHash != nil {
				// Withdrawals list must be present in body after Shanghai.
				if block.Withdrawals() == nil {
					return errors.New("missing withdrawals in block body")
				}
				if hash := types.DeriveSha(block.Withdrawals(), trie.NewStackTrie(nil)); hash != *header.WithdrawalsHash {
					return fmt.Errorf("withdrawals root hash mismatch (header value %x, calculated %x)", *header.WithdrawalsHash, hash)
				}
			} else if block.Withdrawals() != nil { // Withdrawals turn into empty from nil when BlockBody has Sidecars
				// Withdrawals are not allowed prior to shanghai fork
				return errors.New("withdrawals present in block body")
			}
			// Blob transactions may be present after the Cancun fork.
			var blobs int
			for i, tx := range block.Transactions() {
				// Count the number of blobs to validate against the header's blobGasUsed
				blobs += len(tx.BlobHashes())

				// If the tx is a blob tx, it must NOT have a sidecar attached to be valid in a block.
				if tx.BlobTxSidecar() != nil {
					return fmt.Errorf("unexpected blob sidecar in transaction at index %d", i)
				}

				// The individual checks for blob validity (version-check + not empty)
				// happens in state transition.
			}

			// Check blob gas usage.
			if header.BlobGasUsed != nil {
				if want := *header.BlobGasUsed / params.BlobTxBlobGasPerBlob; uint64(blobs) != want { // div because the header is surely good vs the body might be bloated
					return fmt.Errorf("blob gas used mismatch (header %v, calculated %v)", *header.BlobGasUsed, blobs*params.BlobTxBlobGasPerBlob)
				}
			} else {
				if blobs > 0 {
					return errors.New("data blobs present in block body")
				}
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
		func() error {
			if v.remoteValidator != nil && !v.remoteValidator.AncestorVerified(block.Header()) {
				return fmt.Errorf("%w, number: %s, hash: %s", ErrAncestorHasNotBeenVerified, block.Number(), block.Hash())
			}
			return nil
		},
		func() error {
			// Validate block signature
			if err := v.validateBlockSignature(block); err != nil {
				return err
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

// validateBlockSignature validates the block signature
func (v *BlockValidator) validateBlockSignature(block *types.Block) error {
	header := block.Header()

	log.Info("validateBlockSignature called", "number", header.Number, "isEinsteinSignatureActive", v.bc.Config().IsEinsteinSignatureActive(header.Number))

	// Check if Einstein signature validation is active
	if !v.bc.Config().IsEinsteinSignatureActive(header.Number) {
		// For blocks before Einstein signature activation, skip signature validation
		log.Info("Einstein signature not yet active, skipping signature validation")
		return nil
	}

	// Check if V, R, S exist
	if header.V == nil || header.R == nil || header.S == nil {
		log.Warn("Block signature missing", "v", header.V, "r", header.R, "s", header.S)
		return errors.New("block signature missing (V, R, or S is nil)")
	}

	// Get signer from header
	signer := header.Signer
	if signer == nil || *signer == (common.Address{}) {
		return errors.New("block signer missing")
	}

	//log.Info("validateBlockSignature", "number", header.Number, "signer", signer, "coinbase", header.Coinbase)

	// Ensure header.Extra has enough space for extraSeal (65 bytes)
	// EncodeSigHeader assumes header.Extra has at least 65 bytes
	// extraSeal := 65
	// if len(header.Extra) < extraSeal {
	// 	return errors.New("block extra data too short for signature validation")
	// }

	// Calculate the block header hash
	hash := v.bc.Engine().SealHash(header)
	//log.Info("Validating block signature", "number", header.Number, "hash", hash, "signer", signer)

	// SignData does keccak256(data) internally, so we need to do the same here
	signingHash := crypto.Keccak256(hash.Bytes())

	// Combine V, R, S into signature (65 bytes: V(1) + R(32) + S(32))
	signature := make([]byte, 65)
	signature[64] = byte(header.V.Uint64())
	rBytes := header.R.Bytes()
	sBytes := header.S.Bytes()
	// R and S are big-endian, so we need to pad to 32 bytes
	copy(signature[32-len(rBytes):32], rBytes)
	copy(signature[64-len(sBytes):], sBytes)

	// Recover the public key from the signature
	pubkey, err := crypto.Ecrecover(signingHash, signature)
	if err != nil {
		return fmt.Errorf("failed to recover public key from signature: %v", err)
	}

	// Derive the address from the public key
	var recoveredAddr common.Address
	copy(recoveredAddr[:], crypto.Keccak256(pubkey[1:])[12:])
	// Check if the recovered address matches the signer
	if recoveredAddr != *signer {
		//log.Info("Signature validation failed: recovered address does not match signer", "recovered", recoveredAddr, "signer", signer)
		return errors.New("block signature invalid: recovered address does not match signer")
	}

	// Get coinbase (recipient)
	recipient := header.Coinbase

	// Check if signer is the recipient
	if *signer == recipient {
		//	log.Info("Signature validation passed: signer is recipient", "signer", signer)
		return nil
	}

	// Check delegate from contract at 0x000000000000000000000000000000000000C009
	delegateAddr, err := v.getDelegateAddress(recipient, header)
	if err != nil {
		log.Info("Failed to get delegate address", "err", err)
		return errors.New("block signature invalid: signer not authorized")
	}

	//log.Debug("Signature validation: checking delegate", "signer", signer, "delegate", delegateAddr, "recipient", recipient)
	// Check if signer matches the delegate address
	if *signer == delegateAddr {
		log.Info("Signature validation passed: signer is delegate of recipient", "signer", signer, "delegate", delegateAddr)
		return nil
	}

	//log.Info("Signature validation failed: signer not authorized", "signer", signer, "delegate", delegateAddr)
	return errors.New("block signature invalid: signer not authorized")
}

// getDelegateAddress retrieves the delegate address for a given recipient from the contract
func (v *BlockValidator) getDelegateAddress(recipient common.Address, header *types.Header) (common.Address, error) {
	// Get the parent block to access its state
	parent := v.bc.GetBlock(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil {
		return common.Address{}, errors.New("parent block not found")
	}

	// Get the state at the parent block
	stateDB, err := v.bc.StateAt(parent.Root())
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get state at parent block: %v", err)
	}

	// Parse the ABI from systemcontract package
	delegateABI, err := abi.JSON(strings.NewReader(systemcontract.MinerDelegateABI))
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to parse delegate ABI: %v", err)
	}

	// Pack the call data - using minter_delegate method
	delegateMethod := "minter_delegate"
	data, err := delegateABI.Pack(delegateMethod, recipient)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to pack delegate call: %v", err)
	}

	// Create a minimal chain context
	chainContext := &minimalChainContext{bc: v.bc}

	// Create EVM block context
	blockContext := NewEVMBlockContext(header, chainContext, &header.Coinbase)

	// Create EVM instance
	evm := vm.NewEVM(blockContext, stateDB, v.bc.Config(), vm.Config{})
	// Contract address from systemcontract package
	contractAddr := systemcontract.MinerDelegateContractAddr
	// Call the contract
	ret, _, err := evm.StaticCall(vm.AccountRef(recipient), contractAddr, data, math.MaxUint64)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to call delegate contract: %v", err)
	}
	// Unpack the result
	var delegatedAddr common.Address
	if err := delegateABI.UnpackIntoInterface(&delegatedAddr, delegateMethod, ret); err != nil {
		return common.Address{}, fmt.Errorf("failed to unpack delegate result: %v", err)
	}
	log.Info("getDelegateAddress result", "recipient", recipient, "delegate", delegatedAddr)
	return delegatedAddr, nil
}

// minimalChainContext implements ChainContext for EVM execution
type minimalChainContext struct {
	bc *BlockChain
}

func (m *minimalChainContext) Engine() consensus.Engine {
	return m.bc.Engine()
}

func (m *minimalChainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	return m.bc.GetHeader(hash, number)
}

func (m *minimalChainContext) Config() *params.ChainConfig {
	return m.bc.Config()
}

// ValidateState validates the various changes that happen after a state transition,
// such as amount of used gas, the receipt roots and the state root itself.
func (v *BlockValidator) ValidateState(block *types.Block, statedb *state.StateDB, res *ProcessResult, stateless bool) error {
	if res == nil {
		return errors.New("nil ProcessResult value")
	}
	header := block.Header()
	if block.GasUsed() != res.GasUsed {
		return fmt.Errorf("invalid gas used (remote: %d local: %d)", block.GasUsed(), res.GasUsed)
	}
	// Validate the received block's bloom with the one derived from the generated receipts.
	// For valid blocks this should always validate to true.
	validateFuns := []func() error{
		func() error {
			rbloom := types.CreateBloom(res.Receipts)
			if rbloom != header.Bloom {
				return fmt.Errorf("invalid bloom (remote: %x  local: %x)", header.Bloom, rbloom)
			}
			return nil
		},
	}
	// In stateless mode, return early because the receipt and state root are not
	// provided through the witness, rather the cross validator needs to return it.
	if !stateless {
		validateFuns = append(validateFuns, func() error {
			// The receipt Trie's root (R = (Tr [[H1, R1], ... [Hn, Rn]]))
			receiptSha := types.DeriveSha(res.Receipts, trie.NewStackTrie(nil))
			if receiptSha != header.ReceiptHash {
				return fmt.Errorf("invalid receipt root hash (remote: %x local: %x)", header.ReceiptHash, receiptSha)
			}

			// Validate the parsed requests match the expected header value.
			return v.bc.engine.VerifyRequests(block.Header(), res.Requests)
		})
		validateFuns = append(validateFuns, func() error {
			// Validate the state root against the received state root and throw
			// an error if they don't match.
			if root := statedb.IntermediateRoot(v.config.IsEIP158(header.Number)); header.Root != root {
				return fmt.Errorf("invalid merkle root (remote: %x local: %x) dberr: %w", header.Root, root, statedb.Error())
			}
			return nil
		})
	}

	validateRes := make(chan error, len(validateFuns))
	for _, f := range validateFuns {
		tmpFunc := f
		go func() {
			validateRes <- tmpFunc()
		}()
	}

	var err error
	for i := 0; i < len(validateFuns); i++ {
		r := <-validateRes
		if r != nil && err == nil {
			err = r
		}
	}
	return err
}

func (v *BlockValidator) RemoteVerifyManager() *remoteVerifyManager {
	return v.remoteValidator
}

// CalcGasLimit computes the gas limit of the next block after parent. It aims
// to keep the baseline gas close to the provided target, and increase it towards
// the target if the baseline gas is lower.
func CalcGasLimit(parentGasLimit, desiredLimit uint64) uint64 {
	// change GasLimitBoundDivisor to 1024 from 256 from lorentz hard fork, but no need hard fork control here.
	delta := parentGasLimit/params.GasLimitBoundDivisor - 1
	limit := parentGasLimit
	if desiredLimit < params.MinGasLimit {
		desiredLimit = params.MinGasLimit
	}
	// If we're outside our allowed gas range, we try to hone towards them
	if limit < desiredLimit {
		limit = parentGasLimit + delta
		if limit > desiredLimit {
			limit = desiredLimit
		}
		return limit
	}
	if limit > desiredLimit {
		limit = parentGasLimit - delta
		if limit < desiredLimit {
			limit = desiredLimit
		}
	}
	return limit
}
