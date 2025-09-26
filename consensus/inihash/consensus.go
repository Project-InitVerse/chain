// Copyright 2017 The go-ethereum Authors
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

package inihash

import (
	"errors"
	"fmt"
	"github.com/Project-InitVerse/chain/consensus/inihash/systemcontract"
	"github.com/Project-InitVerse/chain/consensus/inihash/vmcaller"
	"github.com/Project-InitVerse/chain/core"
	"github.com/Project-InitVerse/chain/core/tracing"
	"github.com/Project-InitVerse/chain/core/vm"
	"github.com/Project-InitVerse/chain/crypto/versaHash"
	"github.com/Project-InitVerse/chain/log"
	"github.com/holiman/uint256"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
	"runtime"
	"time"

	"github.com/Project-InitVerse/chain/common"
	"github.com/Project-InitVerse/chain/common/gopool"
	"github.com/Project-InitVerse/chain/consensus"
	"github.com/Project-InitVerse/chain/core/state"
	"github.com/Project-InitVerse/chain/core/types"
	"github.com/Project-InitVerse/chain/params"
	"github.com/Project-InitVerse/chain/rlp"
	"github.com/Project-InitVerse/chain/trie"
	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/crypto/sha3"
)

// Inihash proof-of-work protocol constants.
var (
	BaseBlockReward               = big.NewInt(0).Mul(big.NewInt(145833333), big.NewInt(1e+11)) // Block reward in wei for successfully mining a block
	maxUncles                     = 0                                                           // Maximum number of uncles allowed in a single block
	allowedFutureBlockTimeSeconds = int64(140)                                                  // Max seconds from current time allowed for blocks, before they're considered future blocks

)

type blacklistDirection uint

const (
	DirectionFrom blacklistDirection = iota
	DirectionTo
	DirectionBoth
)
const gasLimitBoundDivisorBeforeLorentz uint64 = 256 // The bound divisor of the gas limit, used in update calculations before lorentz hard fork.

// Various error messages to mark blocks invalid. These should be private to
// prevent engine specific errors from being referenced in the remainder of the
// codebase, inherently breaking if the engine is swapped out. Please put common
// error types into the consensus package.
var (
	errOlderBlockTime    = errors.New("timestamp older than parent")
	errTooManyUncles     = errors.New("too many uncles")
	errDuplicateUncle    = errors.New("duplicate uncle")
	errUncleIsAncestor   = errors.New("uncle is ancestor")
	errDanglingUncle     = errors.New("uncle's parent is not ancestor")
	errInvalidDifficulty = errors.New("non-positive difficulty")
	errInvalidMixDigest  = errors.New("invalid mix digest")
	errInvalidPoW        = errors.New("invalid proof-of-work")
)

// Author implements consensus.Engine, returning the header's coinbase as the
// proof-of-work verified author of the block.
func (inihash *Inihash) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func CalBlockReward(blockNumber uint64, multi uint64, forkBlock int64, forkMulti uint64) *big.Int {
	if forkBlock == -1 {
		epoch := (blockNumber/20160)*20160 + 20160
		rate := math.Pow(math.E, float64(epoch)*float64(-0.00000012096))
		rateStr := fmt.Sprintf("%.9f", rate)
		rateDecimal, _ := decimal.NewFromString(rateStr)

		rewardTmp := decimal.NewFromBigInt(BaseBlockReward, 0).Mul(rateDecimal).Mul(decimal.NewFromUint64(multi))

		return rewardTmp.BigInt()
	} else {
		convertBlock := blockNumber
		isMulit := false
		if blockNumber > uint64(forkBlock) {
			isMulit = true
			convertBlock = (blockNumber-uint64(forkBlock))/forkMulti + uint64(forkBlock) + 1
		}

		epoch := (convertBlock/20160)*20160 + 20160
		rate := math.Pow(math.E, float64(epoch)*float64(-0.00000012096))
		rateStr := fmt.Sprintf("%.9f", rate)
		rateDecimal, _ := decimal.NewFromString(rateStr)
		if isMulit {
			rewardTmp := decimal.NewFromBigInt(BaseBlockReward, 0).Mul(rateDecimal).Mul(decimal.NewFromUint64(multi)).Div(decimal.NewFromUint64(forkMulti))
			return rewardTmp.BigInt()
		} else {
			rewardTmp := decimal.NewFromBigInt(BaseBlockReward, 0).Mul(rateDecimal).Mul(decimal.NewFromUint64(multi))
			return rewardTmp.BigInt()
		}

	}

}

// VerifyHeader checks whether a header conforms to the consensus rules of the
// stock Ethereum inihash engine.
func (inihash *Inihash) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header) error {
	// If we're running a full engine faking, accept any input as valid
	if inihash.config.PowMode == ModeFullFake {
		return nil
	}
	// Short circuit if the header is known, or its parent not
	number := header.Number.Uint64()
	if chain.GetHeader(header.Hash(), number) != nil {
		return nil
	}
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	// Sanity checks passed, do a proper verification
	return inihash.verifyHeader(chain, header, parent, false, true, time.Now().Unix())
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers
// concurrently. The method returns a quit channel to abort the operations and
// a results channel to retrieve the async verifications.
func (inihash *Inihash) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header) (chan<- struct{}, <-chan error) {
	// If we're running a full engine faking, accept any input as valid
	if inihash.config.PowMode == ModeFullFake || len(headers) == 0 {
		abort, results := make(chan struct{}), make(chan error, len(headers))
		for i := 0; i < len(headers); i++ {
			results <- nil
		}
		return abort, results
	}
	// Spawn as many workers as allowed threads
	workers := runtime.GOMAXPROCS(0)
	if len(headers) < workers {
		workers = len(headers)
	}
	// Create a task channel and spawn the verifiers
	var (
		inputs  = make(chan int)
		done    = make(chan int, workers)
		errors  = make([]error, len(headers))
		abort   = make(chan struct{})
		unixNow = time.Now().Unix()
	)
	for i := 0; i < workers; i++ {
		gopool.Submit(func() {
			for index := range inputs {
				errors[index] = inihash.verifyHeaderWorker(chain, headers, index, unixNow)
				done <- index
			}
		})
	}

	errorsOut := make(chan error, len(headers))
	gopool.Submit(func() {
		defer close(inputs)
		var (
			in, out = 0, 0
			checked = make([]bool, len(headers))
			inputs  = inputs
		)
		for {
			select {
			case inputs <- in:
				if in++; in == len(headers) {
					// Reached end of headers. Stop sending to workers.
					inputs = nil
				}
			case index := <-done:
				for checked[index] = true; checked[out]; out++ {
					errorsOut <- errors[out]
					if out == len(headers)-1 {
						return
					}
				}
			case <-abort:
				return
			}
		}
	})
	return abort, errorsOut
}

func (inihash *Inihash) verifyHeaderWorker(chain consensus.ChainHeaderReader, headers []*types.Header, index int, unixNow int64) error {
	var parent *types.Header
	if index == 0 {
		parent = chain.GetHeader(headers[0].ParentHash, headers[0].Number.Uint64()-1)
	} else if headers[index-1].Hash() == headers[index].ParentHash {
		parent = headers[index-1]
	}
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	return inihash.verifyHeader(chain, headers[index], parent, false, true, unixNow)
}

// VerifyUncles verifies that the given block's uncles conform to the consensus
// rules of the stock Ethereum inihash engine.
func (inihash *Inihash) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	// If we're running a full engine faking, accept any input as valid
	if inihash.config.PowMode == ModeFullFake {
		return nil
	}
	// Verify that there are at most 2 uncles included in this block
	if len(block.Uncles()) > maxUncles {
		return errTooManyUncles
	}
	if len(block.Uncles()) == 0 {
		return nil
	}
	// Gather the set of past uncles and ancestors
	uncles, ancestors := mapset.NewSet[common.Hash](), make(map[common.Hash]*types.Header)

	number, parent := block.NumberU64()-1, block.ParentHash()
	for i := 0; i < 7; i++ {
		ancestorHeader := chain.GetHeader(parent, number)
		if ancestorHeader == nil {
			break
		}
		ancestors[parent] = ancestorHeader
		// If the ancestor doesn't have any uncles, we don't have to iterate them
		if ancestorHeader.UncleHash != types.EmptyUncleHash {
			// Need to add those uncles to the blacklist too
			ancestor := chain.GetBlock(parent, number)
			if ancestor == nil {
				break
			}
			for _, uncle := range ancestor.Uncles() {
				uncles.Add(uncle.Hash())
			}
		}
		parent, number = ancestorHeader.ParentHash, number-1
	}
	ancestors[block.Hash()] = block.Header()
	uncles.Add(block.Hash())

	// Verify each of the uncles that it's recent, but not an ancestor
	for _, uncle := range block.Uncles() {
		// Make sure every uncle is rewarded only once
		hash := uncle.Hash()
		if uncles.Contains(hash) {
			return errDuplicateUncle
		}
		uncles.Add(hash)

		// Make sure the uncle has a valid ancestry
		if ancestors[hash] != nil {
			return errUncleIsAncestor
		}
		if ancestors[uncle.ParentHash] == nil || uncle.ParentHash == block.ParentHash() {
			return errDanglingUncle
		}
		if err := inihash.verifyHeader(chain, uncle, ancestors[uncle.ParentHash], true, true, time.Now().Unix()); err != nil {
			return err
		}
	}
	return nil
}

// verifyHeader checks whether a header conforms to the consensus rules of the
// stock Ethereum inihash engine.
// See YP section 4.3.4. "Block Header Validity"
func (inihash *Inihash) verifyHeader(chain consensus.ChainHeaderReader, header, parent *types.Header, uncle bool, seal bool, unixNow int64) error {
	// Ensure that the header's extra-data section is of a reasonable size
	if uint64(len(header.Extra)) > params.MaximumExtraDataSize {
		return fmt.Errorf("extra-data too long: %d > %d", len(header.Extra), params.MaximumExtraDataSize)
	}

	// Verify the header's timestamp
	if !uncle {
		if header.Time > uint64(unixNow+allowedFutureBlockTimeSeconds) {
			return consensus.ErrFutureBlock
		}
	}
	if header.Time <= parent.Time {
		return errOlderBlockTime
	}
	// Verify the block's difficulty based on its timestamp and parent's difficulty
	expected := CalcDifficulty(chain.Config(), header.Time, parent)

	if expected.Cmp(header.Difficulty) != 0 {
		return fmt.Errorf("invalid difficulty: have %v, want %v", header.Difficulty, expected)
	}
	// Verify that the gas limit is <= 2^63-1
	cap := uint64(0x7fffffffffffffff)
	if header.GasLimit > cap {
		return fmt.Errorf("invalid gasLimit: have %v, max %v", header.GasLimit, cap)
	}
	// Verify that the gasUsed is <= gasLimit
	if header.GasUsed > header.GasLimit {
		return fmt.Errorf("invalid gasUsed: have %d, gasLimit %d", header.GasUsed, header.GasLimit)
	}

	// Verify that the gas limit remains within allowed bounds
	diff := int64(parent.GasLimit) - int64(header.GasLimit)
	if diff < 0 {
		diff *= -1
	}

	gasLimitBoundDivisor := gasLimitBoundDivisorBeforeLorentz
	if chain.Config().IsLorentz(header.Number, header.Time) {
		gasLimitBoundDivisor = params.GasLimitBoundDivisor
	}

	limit := parent.GasLimit / gasLimitBoundDivisor

	if uint64(diff) >= limit || header.GasLimit < params.MinGasLimit {
		return fmt.Errorf("invalid gas limit: have %d, want %d += %d", header.GasLimit, parent.GasLimit, limit)
	}
	// Verify that the block number is parent's +1
	if diff := new(big.Int).Sub(header.Number, parent.Number); diff.Cmp(big.NewInt(1)) != 0 {
		return consensus.ErrInvalidNumber
	}
	// Verify the engine specific seal securing the block
	//if seal {
	//	if err := inihash.verifySeal(chain, header, false); err != nil {
	//		return err
	//	}
	//}

	return nil
}
func (inihash *Inihash) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	return CalcDifficulty(chain.Config(), time, parent)
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns
// the difficulty that a new block should have when created at time
// given the parent block's time and difficulty.
func CalcDifficulty(config *params.ChainConfig, time uint64, parent *types.Header) *big.Int {
	//next := new(big.Int).Add(parent.Number, big1)
	if config.ChainID == nil {
		return calcDifficulty(time, parent)
	} else if config.ChainID.Int64() == 7234 {
		if config.IsNewTon(parent.Number) {
			return calcDifficultyNewTon(time, parent)
		} else if parent.Number.Int64() >= params.TestnetForkBlockNumber {
			return calcDifficultyNew(time, parent)
		} else {
			return calcDifficulty(time, parent)
		}

	} else if config.ChainID.Int64() == 7233 {
		//mainnet
		if config.IsNewTon(parent.Number) {
			return calcDifficultyNewTon(time, parent)
		} else if parent.Number.Int64() >= params.MainnetForkBlockNumber {
			return calcDifficultyNew(time, parent)
		} else {
			return calcDifficulty(time, parent)
		}
	}
	return calcDifficulty(time, parent)
}

// Some weird constants to avoid constant memory allocs for them.
var (
	expDiffPeriod = big.NewInt(100000)
	big1          = big.NewInt(1)
	big2          = big.NewInt(2)
	big5          = big.NewInt(5)
	big6          = big.NewInt(6)
	big9          = big.NewInt(9)
	big10         = big.NewInt(10)
	bigMinus599   = big.NewInt(-599)
	bigMinus99    = big.NewInt(-99)
)

// calcDifficultyHomestead is the difficulty adjustment algorithm. It returns
// the difficulty that a new block should have when created at time given the
// parent block's time and difficulty. The calculation uses the Homestead rules.
func calcDifficulty(time uint64, parent *types.Header) *big.Int {
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-2.md
	// algorithm:
	// diff = (parent_diff +
	//         (parent_diff / 12288 * max(6 - (block_timestamp - parent_timestamp) // 5, -599))
	//        )
	bigTime := new(big.Int).SetUint64(time)
	bigParentTime := new(big.Int).SetUint64(parent.Time)
	// holds intermediate values to make the algo easier to read & audit
	x := new(big.Int)
	y := new(big.Int)
	// 1 - (block_timestamp - parent_timestamp) // 5
	x.Sub(bigTime, bigParentTime)
	x.Div(x, big5)
	x.Sub(big6, x)
	// max(1 - (block_timestamp - parent_timestamp) // 5, -599)
	if x.Cmp(bigMinus599) < 0 {
		x.Set(bigMinus599)
	}
	// (parent_diff + parent_diff // 12288 * max(1 - (block_timestamp - parent_timestamp) // 5, -599))
	y.Div(parent.Difficulty, params.DifficultyBoundDivisor)
	x.Mul(y, x)
	x.Add(parent.Difficulty, x)

	// minimum difficulty can ever be (before exponential factor)
	if x.Cmp(params.MinimumDifficulty) < 0 {
		x.Set(params.MinimumDifficulty)
	}

	return x
}

// calcDifficultyHomestead is the difficulty adjustment algorithm. It returns
// the difficulty that a new block should have when created at time given the
// parent block's time and difficulty. The calculation uses the Homestead rules.
func calcDifficultyNew(time uint64, parent *types.Header) *big.Int {
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-2.md
	// algorithm:
	// diff = (parent_diff +
	//         (parent_diff / 2048 * max(1 - (block_timestamp - parent_timestamp) // 10, -99))
	//        )
	bigTime := new(big.Int).SetUint64(time)
	bigParentTime := new(big.Int).SetUint64(parent.Time)
	// holds intermediate values to make the algo easier to read & audit
	x := new(big.Int)
	y := new(big.Int)
	// 1 - (block_timestamp - parent_timestamp) // 10
	x.Sub(bigTime, bigParentTime)
	x.Div(x, big10)
	x.Sub(big1, x)
	// max(1 - (block_timestamp - parent_timestamp) // 5, -599)
	if x.Cmp(bigMinus99) < 0 {
		x.Set(bigMinus99)
	}
	// (parent_diff + parent_diff // 12288 * max(1 - (block_timestamp - parent_timestamp) // 5, -599))
	y.Div(parent.Difficulty, params.NewDifficultyBoundDivisor)
	x.Mul(y, x)
	x.Add(parent.Difficulty, x)

	// minimum difficulty can ever be (before exponential factor)
	if x.Cmp(params.MinimumDifficulty) < 0 {
		x.Set(params.MinimumDifficulty)
	}

	return x
}
func calcDifficultyNewTon(time uint64, parent *types.Header) *big.Int {
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-2.md
	// algorithm:
	// diff = (parent_diff +
	//         (parent_diff / 2048 * max(1 - (block_timestamp - parent_timestamp) // 10, -99))
	//        )
	bigTime := new(big.Int).SetUint64(time)
	bigParentTime := new(big.Int).SetUint64(parent.Time)
	// holds intermediate values to make the algo easier to read & audit
	x := new(big.Int)
	y := new(big.Int)
	// 5 - (block_timestamp - parent_timestamp) // 2
	x.Sub(bigTime, bigParentTime)
	x.Div(x, big2)
	x.Sub(big5, x)
	// max(5 - (block_timestamp - parent_timestamp) // 5, -99)
	if x.Cmp(bigMinus99) < 0 {
		x.Set(bigMinus99)
	}
	// (parent_diff + parent_diff // 2048 * max(5 - (block_timestamp - parent_timestamp) // 2, -99))
	y.Div(parent.Difficulty, params.NewDifficultyBoundDivisor)
	x.Mul(y, x)
	x.Add(parent.Difficulty, x)

	// minimum difficulty can ever be (before exponential factor)
	if x.Cmp(params.MinimumDifficulty) < 0 {
		x.Set(params.MinimumDifficulty)
	}

	return x
}

// Exported for fuzzing
var BaseDifficultyCalulator = calcDifficulty

// verifySeal checks whether a block satisfies the PoW difficulty requirements,
// either using the usual inihash cache for it, or alternatively using a full DAG
// to make remote mining fast.
func (inihash *Inihash) verifySeal(chain consensus.ChainHeaderReader, header *types.Header, fulldag bool, bigUtil types.BigInt) error {
	// If we're running a fake PoW, accept any seal as valid
	if inihash.config.PowMode == ModeFake || inihash.config.PowMode == ModeFullFake {
		time.Sleep(inihash.fakeDelay)
		if inihash.fakeFail == header.Number.Uint64() {
			return errInvalidPoW
		}
		return nil
	}
	// If we're running a shared PoW, delegate verification to it
	if inihash.shared != nil {
		return inihash.shared.verifySeal(chain, header, fulldag, bigUtil)
	}
	// Ensure that we have a valid difficulty for the block
	if header.Difficulty.Sign() <= 0 {
		return errInvalidDifficulty
	}
	// Recompute the digest and PoW values
	//number := header.Number.Uint64()

	var (
		//digest []byte
		result []byte
	)
	// If fast-but-heavy PoW verification was requested, use an inihash dataset
	headerHash := inihash.SealHash(header).Bytes()
	result = versaHash.VersaHash(headerHash, header.Nonce[:], header.ExtraNonce[:])

	target := new(big.Int).Div(two256, header.Difficulty)

	if bigUtil.SetBytes(result).Cmp(target) > 0 {
		return errInvalidPoW
	}
	return nil
}

// Prepare implements consensus.Engine, initializing the difficulty field of a
// header to conform to the inihash protocol. The changes are done inline.
func (inihash *Inihash) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	//todo set real address
	header.TeamAddress = common.Address{}
	header.TeamRate = 0
	header.Difficulty = CalcDifficulty(chain.Config(), header.Time, parent)
	return nil
}

// Finalize implements consensus.Engine, accumulating the block and uncle rewards,
// setting the final state on the header
func (inihash *Inihash) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state vm.StateDB, txs *[]*types.Transaction,
	uncles []*types.Header, _ []*types.Withdrawal, receipts *[]*types.Receipt, systemTxs *[]*types.Transaction, usedGas *uint64, tracer *tracing.Hooks) error {
	// Accumulate any block and uncle rewards and commit the final state root
	accumulateRewards(chain.Config(), state, header, uncles)

	bigUtil := common.NewBig.Copy()

	if chain.Config().IsNewTon(big.NewInt(0).Sub(header.Number, common.Big3)) {

		abi := systemcontract.GetInteractiveABI()[systemcontract.AddressListContractName]
		data, err := abi.Pack(systemcontract.ValidMethod, header.Coinbase)

		if err != nil {
			log.Error("Can't pack data ", "method", systemcontract.ValidMethod, "err", err)

		} else {
			msg := core.NewMessage(systemcontract.AddressListContractAddr, &systemcontract.AddressListContractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, nil, false)

			result, err := vmcaller.ExecuteMsg(msg, state, header, newMinimalChainContext(inihash), inihash.chainConfig)
			if err != nil {

			} else {
				ret, err := abi.Unpack(systemcontract.ValidMethod, result)
				if err != nil {
					return err
				}
				if len(ret) != 1 {
					return errors.New("invalid params length")
				}
				bigUtil.SetUint64(ret[0].(*big.Int).Uint64())
			}
		}

	}
	if err := inihash.verifySeal(chain, header, false, bigUtil); err != nil {
		return err
	}

	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	return nil
}

// FinalizeAndAssemble implements consensus.Engine, accumulating the block and
// uncle rewards, setting the final state and assembling the block.
func (inihash *Inihash) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB,
	body *types.Body, receipts []*types.Receipt, tracer *tracing.Hooks) (*types.Block, []*types.Receipt, error) {

	if body.Transactions == nil {
		body.Transactions = make([]*types.Transaction, 0)
	}
	if receipts == nil {
		receipts = make([]*types.Receipt, 0)
	}
	accumulateRewards(chain.Config(), state, header, make([]*types.Header, 0))

	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))

	// Header seems complete, assemble into a block and return
	return types.NewBlock(header, body, receipts, trie.NewStackTrie(nil)), receipts, nil
}

func (inihash *Inihash) ValidateTx(tx *types.Transaction, header *types.Header, parentState *state.StateDB) error {
	if header.Number.Int64() > 10 && inihash.chainConfig.IsNewTon(big.NewInt(0).Sub(header.Number, common.Big3)) {
		signer := types.MakeSigner(inihash.chainConfig, new(big.Int).SetUint64(header.Number.Uint64()), header.Time)
		from, err := types.Sender(signer, tx)
		if err != nil {
			return err
		}
		m, err := inihash.getBlacklist(header, parentState)
		if err != nil {
			log.Error("can't get blacklist", "err", err)
			return err
		}
		if d, exist := m[from]; exist && (d != DirectionTo) {
			return errors.New("address denied")
		}
		if to := tx.To(); to != nil {
			if d, exist := m[*to]; exist && (d != DirectionFrom) {
				return errors.New("address denied")
			}
		}
	}
	return nil
}

func (inihash *Inihash) getBlacklist(header *types.Header, parentState *state.StateDB) (map[common.Address]blacklistDirection, error) {

	if v, ok := inihash.blacklists.Get(header.ParentHash); ok {
		return v.(map[common.Address]blacklistDirection), nil
	}

	inihash.blLock.Lock()
	defer inihash.blLock.Unlock()
	if v, ok := inihash.blacklists.Get(header.ParentHash); ok {
		return v.(map[common.Address]blacklistDirection), nil
	}

	abi := systemcontract.GetInteractiveABI()[systemcontract.AddressListContractName]
	get := func(method string) ([]common.Address, error) {
		data, err := abi.Pack(method)
		if err != nil {
			log.Error("Can't pack data ", "method", method, "err", err)
			return []common.Address{}, err
		}

		msg := core.NewMessage(header.Coinbase, &systemcontract.AddressListContractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, nil, false)

		// Note: It's safe to use minimalChainContext for executing AddressListContract
		result, err := vmcaller.ExecuteMsg(msg, parentState, header, newMinimalChainContext(inihash), inihash.chainConfig)
		if err != nil {
			return []common.Address{}, err
		}

		// unpack data
		ret, err := abi.Unpack(method, result)
		if err != nil {
			return []common.Address{}, err
		}
		if len(ret) != 1 {
			return []common.Address{}, errors.New("invalid params length")
		}
		blacks, ok := ret[0].([]common.Address)
		if !ok {
			return []common.Address{}, errors.New("invalid blacklist format")
		}
		return blacks, nil
	}
	froms, err := get("getBlacksFrom")
	if err != nil {
		return nil, err
	}
	tos, err := get("getBlacksTo")

	if err != nil {
		return nil, err
	}

	m := make(map[common.Address]blacklistDirection)
	for _, from := range froms {
		m[from] = DirectionFrom
	}
	for _, to := range tos {
		if _, exist := m[to]; exist {
			m[to] = DirectionBoth
		} else {
			m[to] = DirectionTo
		}
	}
	inihash.blacklists.Add(header.ParentHash, m)
	return m, nil
}

func (inihash *Inihash) Delay(_ consensus.ChainReader, _ *types.Header, leftOver *time.Duration) *time.Duration {
	return nil
}

// SealHash returns the hash of a block prior to it being sealed.
func (inihash *Inihash) SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()

	rlp.Encode(hasher, []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra,
		header.Provider,
		header.TeamAddress,
		header.ValidatorRate,
		header.TeamRate,
	})

	hasher.Sum(hash[:0])
	return hash
}

// Some weird constants to avoid constant memory allocs for them.
var (
	big8  = big.NewInt(8)
	big32 = big.NewInt(32)
)

// AccumulateRewards credits the coinbase of the given block with the mining
// reward. The total reward consists of the static block reward and rewards for
// included uncles. The coinbase of each uncle block is also rewarded.
func accumulateRewards(config *params.ChainConfig, state vm.StateDB, header *types.Header, uncles []*types.Header) {
	// Skip block reward in catalyst mode
	//if config.IsCatalyst(header.Number) {
	//	return
	//}
	// Select the correct block reward based on chain progression
	var blockReward *big.Int
	if config.ChainID.Int64() == 7233 {
		//mainnet
		blockReward = CalBlockReward(header.Number.Uint64(), 50, params.MainnetForkBlockNumber, 3)
	} else if config.ChainID.Int64() == 7234 {
		blockReward = CalBlockReward(header.Number.Uint64(), 1, params.TestnetForkBlockNumber, 3)
	} else {
		blockReward = CalBlockReward(header.Number.Uint64(), 1, -1, 1)
	}

	// Accumulate the rewards for the miner and any included uncles
	reward := new(big.Int).Set(blockReward)
	r := new(big.Int)
	for _, uncle := range uncles {
		r.Add(uncle.Number, big8)
		r.Sub(r, header.Number)
		r.Mul(r, blockReward)
		r.Div(r, big8)
		state.AddBalance(uncle.Coinbase, uint256.MustFromBig(r), tracing.BalanceIncreaseRewardMineUncle)

		r.Div(blockReward, big32)
		reward.Add(reward, r)
	}
	PersonalReward := reward
	//TeamReward := big.NewInt(0).Div(big.NewInt(0).Mul(reward, big1), big10)

	state.AddBalance(header.Coinbase, uint256.MustFromBig(PersonalReward), tracing.BalanceIncreaseRewardMineBlock)
	//state.AddBalance(header.TeamAddress, TeamReward)
}
