package core

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

// StateReader defines the minimal interface needed for gasless transaction validation
// This allows ValidateGaslessTx to work with any state implementation that can read storage
type StateReader interface {
	GetState(addr common.Address, hash common.Hash) common.Hash
}

// GasStation struct storage slots structs
type GasStationStorageSlots struct {
	StructBaseSlotHash             common.Hash
	CreditSlotHash                 common.Hash
	WhitelistEnabledSlotHash       common.Hash
	NestedWhitelistMapBaseSlotHash common.Hash
}

// calculateGasStationSlots computes the storage slot hashes for a specific
// registered contract within the GasStation's `contracts` mapping.
// It returns the base slot for the struct (holding packed fields), the slot for credits,
// the slot for whitelistEnabled, and the base slot for the nested whitelist mapping.
func CalculateGasStationSlots(registeredContractAddress common.Address) GasStationStorageSlots {
	gasStationStorageSlots := GasStationStorageSlots{}

	// ERC-7201 storage location for GasStationStorage
	// bytes32 private constant GasStationStorageLocation = 0xc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db00;
	gasStationStorageLocation := common.HexToHash("0xc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db00")

	// The 'contracts' mapping is at offset 1 from the storage location
	// (dao is at offset 0, contracts is at offset 1)
	contractsMapSlot := new(big.Int).Add(gasStationStorageLocation.Big(), big.NewInt(1))

	// Calculate the base slot for the struct entry in the mapping
	keyPadded := common.LeftPadBytes(registeredContractAddress.Bytes(), 32)
	mapSlotPadded := common.LeftPadBytes(contractsMapSlot.Bytes(), 32)
	combined := append(keyPadded, mapSlotPadded...)
	gasStationStorageSlots.StructBaseSlotHash = crypto.Keccak256Hash(combined)

	// Calculate subsequent slots by adding offsets to the base slot hash
	// New struct layout: bool registered, bool active, address admin (all packed in slot 0)
	// uint256 credits (slot 1), bool whitelistEnabled (slot 2), mapping whitelist (slot 3)
	structBaseSlotBig := gasStationStorageSlots.StructBaseSlotHash.Big()

	// Slot for 'credits' (offset 1 from base - after the packed bools and address)
	creditsSlotBig := new(big.Int).Add(structBaseSlotBig, big.NewInt(1))
	gasStationStorageSlots.CreditSlotHash = common.BigToHash(creditsSlotBig)

	// Slot for 'whitelistEnabled' (offset 2 from base)
	whitelistEnabledSlotBig := new(big.Int).Add(structBaseSlotBig, big.NewInt(2))
	gasStationStorageSlots.WhitelistEnabledSlotHash = common.BigToHash(whitelistEnabledSlotBig)

	// Base slot for the nested 'whitelist' mapping (offset 3 from base)
	nestedWhitelistMapBaseSlotBig := new(big.Int).Add(structBaseSlotBig, big.NewInt(3))
	gasStationStorageSlots.NestedWhitelistMapBaseSlotHash = common.BigToHash(nestedWhitelistMapBaseSlotBig)

	return gasStationStorageSlots
}

func ValidateGaslessTx(to *common.Address, from common.Address, gasLimit uint64, sdb StateReader) (*big.Int, *big.Int, *GasStationStorageSlots, error) {
	if to == nil {
		return nil, nil, nil, fmt.Errorf("gasless txn must have a valid to address")
	}

	// Calculate GasStation storage slots
	gasStationStorageSlots := CalculateGasStationSlots(*to)

	// Get the storage for the GaslessContract struct for the given address
	storageBaseSlot := sdb.GetState(params.GasStationAddress, gasStationStorageSlots.StructBaseSlotHash)

	// Extract the registered and active bytes from the storage slot
	isRegistered := storageBaseSlot[31] == 0x01
	isActive := storageBaseSlot[30] == 0x01

	if !isRegistered {
		return nil, nil, nil, fmt.Errorf("gasless transaction to unregistered address")
	}

	if !isActive {
		return nil, nil, nil, fmt.Errorf("gasless transaction to inactive address")
	}

	// Get the available credits from the credits storage slot in the GasStation contract for the given address
	availableCredits := sdb.GetState(params.GasStationAddress, gasStationStorageSlots.CreditSlotHash)

	// Convert credits (Hash) and tx gas (uint64) to big.Int for comparison
	availableCreditsBig := new(big.Int).SetBytes(availableCredits.Bytes())
	txRequiredCreditsBig := new(big.Int).SetUint64(gasLimit)

	// Check if contract has enough available credits to cover the cost of the tx
	if availableCreditsBig.Cmp(txRequiredCreditsBig) < 0 {
		return nil, nil, nil, fmt.Errorf("gasless transaction has insufficient credits: have %v, need %v", availableCreditsBig, txRequiredCreditsBig)
	}

	// Get the whitelist enabled slot
	whitelistEnabled := sdb.GetState(params.GasStationAddress, gasStationStorageSlots.WhitelistEnabledSlotHash)

	// Get the whitelist enabled byte from the whitelist enabled slot
	isWhitelistEnabled := whitelistEnabled[31] == 0x01

	if isWhitelistEnabled {
		// Calculate slot for the specific user in the nested whitelist map
		userKeyPadded := common.LeftPadBytes(from.Bytes(), 32)
		mapBaseSlotPadded := common.LeftPadBytes(gasStationStorageSlots.NestedWhitelistMapBaseSlotHash.Bytes(), 32)
		userCombined := append(userKeyPadded, mapBaseSlotPadded...)
		userWhitelistSlotHash := crypto.Keccak256Hash(userCombined)

		// Get the whitelist status for the specific user
		userWhitelist := sdb.GetState(params.GasStationAddress, userWhitelistSlotHash)

		// Check if the user is whitelisted
		userWhitelistByte := userWhitelist[31]
		isUserWhitelistStorage := userWhitelistByte == 0x01

		if !isUserWhitelistStorage {
			return nil, nil, nil, fmt.Errorf("gasless transaction to non-whitelisted address")
		}
	}

	return availableCreditsBig, txRequiredCreditsBig, &gasStationStorageSlots, nil
}
