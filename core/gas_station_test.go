package core

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/holiman/uint256"
)

// TestGaslessTransactionGasUsage verifies that gasless transactions use the same amount of gas as standard transactions
func TestGaslessTransactionGasUsage(t *testing.T) {
	// Create test database and state
	db := rawdb.NewMemoryDatabase()
	statedb, err := state.New(types.EmptyRootHash, state.NewDatabase(triedb.NewDatabase(db, nil), nil))
	if err != nil {
		t.Fatalf("Failed to create state: %v", err)
	}

	// Create test accounts
	key, _ := crypto.GenerateKey()
	from := crypto.PubkeyToAddress(key.PublicKey)
	to := common.HexToAddress("0x1000000000000000000000000000000000000001")
	gasStationAddr := params.GasStationAddress

	// Setup account balances
	statedb.SetBalance(from, uint256.NewInt(1000000000000000000), tracing.BalanceChangeUnspecified) // 1 ETH
	statedb.SetBalance(to, uint256.NewInt(0), tracing.BalanceChangeUnspecified)

	// Setup gasless contract as registered and active with credits
	gasStationSlots := CalculateGasStationSlots(to)

	// Set registered (byte 31) and active (byte 30) to 0x01
	var statusSlot common.Hash
	statusSlot[31] = 0x01 // registered
	statusSlot[30] = 0x01 // active
	statedb.SetState(gasStationAddr, gasStationSlots.StructBaseSlotHash, statusSlot)

	// Set credits to 1M gas
	creditsAmount := big.NewInt(1000000)
	statedb.SetState(gasStationAddr, gasStationSlots.CreditSlotHash, common.BigToHash(creditsAmount))

	// Set whitelist disabled (default)
	statedb.SetState(gasStationAddr, gasStationSlots.WhitelistEnabledSlotHash, common.Hash{})

	// Set single-use disabled (default)
	statedb.SetState(gasStationAddr, gasStationSlots.SingleUseEnabledSlotHash, common.Hash{})

	// Test data: simple storage operation
	testData := common.Hex2Bytes("6060604052600560005500") // SSTORE operation

	// Create chainConfig
	chainConfig := params.MainnetChainConfig
	chainConfig.LondonBlock = big.NewInt(0) // Enable London fork

	// Create block context
	blockCtx := vm.BlockContext{
		CanTransfer: func(vm.StateDB, common.Address, *uint256.Int) bool { return true },
		Transfer:    func(vm.StateDB, common.Address, common.Address, *uint256.Int) {},
		GetHash:     func(uint64) common.Hash { return common.Hash{} },
		Coinbase:    common.Address{},
		BlockNumber: big.NewInt(1),
		Time:        1,
		Difficulty:  big.NewInt(1),
		BaseFee:     big.NewInt(1000000000), // 1 gwei
		GasLimit:    10000000,
	}

	// Test 1: Standard transaction
	standardMsg := &Message{
		To:                    &to,
		From:                  from,
		Nonce:                 0,
		Value:                 big.NewInt(0),
		GasLimit:              100000,
		GasPrice:              big.NewInt(1000000000),
		GasFeeCap:             big.NewInt(1000000000),
		GasTipCap:             big.NewInt(1000000000),
		Data:                  testData,
		AccessList:            nil,
		BlobGasFeeCap:         big.NewInt(0),
		BlobHashes:            nil,
		SetCodeAuthorizations: nil,
		SkipNonceChecks:       false,
		SkipFromEOACheck:      false,
		IsSystemTx:            false,
		IsDepositTx:           false,
		Mint:                  nil,
		RollupCostData:        types.RollupCostData{},
		IsGaslessTx:           false,
	}

	// Execute standard transaction
	statedb1 := statedb.Copy()
	evm1 := vm.NewEVM(blockCtx, statedb1, chainConfig, vm.Config{})
	evm1.SetTxContext(NewEVMTxContext(standardMsg))
	gp1 := new(GasPool).AddGas(10000000)

	result1, err := ApplyMessage(evm1, standardMsg, gp1)
	if err != nil {
		t.Fatalf("Standard transaction failed: %v", err)
	}

	// Test 2: Gasless transaction
	gaslessMsg := &Message{
		To:                    &to,
		From:                  from,
		Nonce:                 0,
		Value:                 big.NewInt(0),
		GasLimit:              100000,
		GasPrice:              big.NewInt(0), // No gas price for gasless
		GasFeeCap:             big.NewInt(0),
		GasTipCap:             big.NewInt(0),
		Data:                  testData,
		AccessList:            nil,
		BlobGasFeeCap:         big.NewInt(0),
		BlobHashes:            nil,
		SetCodeAuthorizations: nil,
		SkipNonceChecks:       false,
		SkipFromEOACheck:      false,
		IsSystemTx:            false,
		IsDepositTx:           false,
		Mint:                  nil,
		RollupCostData:        types.RollupCostData{},
		IsGaslessTx:           true,
	}

	// Execute gasless transaction
	statedb2 := statedb.Copy()
	evm2 := vm.NewEVM(blockCtx, statedb2, chainConfig, vm.Config{})
	evm2.SetTxContext(NewEVMTxContext(gaslessMsg))
	gp2 := new(GasPool).AddGas(10000000)

	result2, err := ApplyMessage(evm2, gaslessMsg, gp2)
	if err != nil {
		t.Fatalf("Gasless transaction failed: %v", err)
	}

	// Compare gas usage
	t.Logf("Standard transaction gas used: %d", result1.UsedGas)
	t.Logf("Gasless transaction gas used: %d", result2.UsedGas)

	if result1.UsedGas != result2.UsedGas {
		t.Errorf("Gas usage mismatch: standard=%d, gasless=%d, difference=%d",
			result1.UsedGas, result2.UsedGas, int64(result2.UsedGas)-int64(result1.UsedGas))
	}

	// Verify both transactions succeeded
	if result1.Failed() {
		t.Errorf("Standard transaction failed: %v", result1.Err)
	}
	if result2.Failed() {
		t.Errorf("Gasless transaction failed: %v", result2.Err)
	}

	// Verify that credits were deducted for gasless transaction
	remainingCredits := statedb2.GetState(gasStationAddr, gasStationSlots.CreditSlotHash)
	remainingCreditsBig := new(big.Int).SetBytes(remainingCredits.Bytes())
	expectedRemaining := new(big.Int).Sub(creditsAmount, big.NewInt(int64(result2.UsedGas)))

	if remainingCreditsBig.Cmp(expectedRemaining) != 0 {
		t.Errorf("Credits not properly deducted: expected=%s, actual=%s",
			expectedRemaining.String(), remainingCreditsBig.String())
	}

	t.Logf("✅ Gas usage is identical: %d gas", result1.UsedGas)
	t.Logf("✅ Credits properly deducted: %s -> %s", creditsAmount.String(), remainingCreditsBig.String())
}
