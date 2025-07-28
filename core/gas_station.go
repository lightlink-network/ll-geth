package core

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

// Event signature
var CreditsUsedEventSignature = crypto.Keccak256Hash([]byte("CreditsUsed(address,address,uint256)"))

var (
	addressType, _ = abi.NewType("address", "", nil)
	uint256Type, _ = abi.NewType("uint256", "", nil)

	// Pre-built arguments for CreditsUsed event
	CreditsUsedEventArgs = abi.Arguments{
		{Type: addressType}, // caller (not indexed)
		{Type: uint256Type}, // gasUsed (not indexed)
	}
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
	SingleUseEnabledSlotHash       common.Hash
	UsedAddressesMapBaseSlotHash   common.Hash
}

// calculateGasStationSlots computes the storage slot hashes for a specific
// registered contract within the GasStation's `contracts` mapping.
// It returns the base slot for the struct (holding packed fields), the slot for credits,
// the slot for whitelistEnabled, and the base slot for the nested whitelist mapping.
func CalculateGasStationSlots(registeredContractAddress common.Address) GasStationStorageSlots {
	gasStationStorageSlots := GasStationStorageSlots{}

	// ERC-7201 storage location for GasStationStorage
	// keccak256(abi.encode(uint256(keccak256("gasstation.main")) - 1)) & ~bytes32(uint256(0xff));
	gasStationStorageLocation := common.HexToHash("0x64d1d9a8a451551a9514a2c08ad4e1552ed316d7dd2778a4b9494de741d8e000")

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
	// bool singleUseEnabled (slot 4), mapping usedAddresses (slot 5)
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

	// Slot for 'singleUseEnabled' (offset 4 from base)
	singleUseEnabledSlotBig := new(big.Int).Add(structBaseSlotBig, big.NewInt(4))
	gasStationStorageSlots.SingleUseEnabledSlotHash = common.BigToHash(singleUseEnabledSlotBig)

	// Base slot for the nested 'usedAddresses' mapping (offset 5 from base)
	usedAddressesMapBaseSlotBig := new(big.Int).Add(structBaseSlotBig, big.NewInt(5))
	gasStationStorageSlots.UsedAddressesMapBaseSlotHash = common.BigToHash(usedAddressesMapBaseSlotBig)

	return gasStationStorageSlots
}

func ValidateGaslessTx(to *common.Address, from common.Address, gasLimit uint64, sdb StateReader) (*big.Int, *big.Int, *GasStationStorageSlots, error) {
	if to == nil {
		return nil, nil, nil, fmt.Errorf("gasless txn must have a valid to address")
	}

	if gasLimit == 0 {
		return nil, nil, nil, fmt.Errorf("gasless txn must have a non-zero gas limit")
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
		userWhitelistSlotHash := calculateNestedMappingSlot(from, gasStationStorageSlots.NestedWhitelistMapBaseSlotHash)

		// Get the whitelist status for the specific user
		userWhitelist := sdb.GetState(params.GasStationAddress, userWhitelistSlotHash)

		// Check if the user is whitelisted
		userWhitelistByte := userWhitelist[31]
		isUserWhitelistStorage := userWhitelistByte == 0x01

		if !isUserWhitelistStorage {
			return nil, nil, nil, fmt.Errorf("gasless transaction to non-whitelisted address")
		}
	}

	// Check single-use mode if enabled
	singleUseEnabled := sdb.GetState(params.GasStationAddress, gasStationStorageSlots.SingleUseEnabledSlotHash)
	isSingleUseEnabled := singleUseEnabled[31] == 0x01

	if isSingleUseEnabled {
		// Calculate slot for the specific user in the nested usedAddresses map
		userUsedSlotHash := calculateNestedMappingSlot(from, gasStationStorageSlots.UsedAddressesMapBaseSlotHash)

		// Get the used status for the specific user
		userUsed := sdb.GetState(params.GasStationAddress, userUsedSlotHash)

		// Check if the user has already used gasless transactions
		isUserAlreadyUsed := userUsed[31] == 0x01

		if isUserAlreadyUsed {
			return nil, nil, nil, fmt.Errorf("gasless transaction from address that has already used single-use gasless functionality")
		}
	}

	return availableCreditsBig, txRequiredCreditsBig, &gasStationStorageSlots, nil
}

// calculateNestedMappingSlot computes the storage slot hash for a nested mapping
func calculateNestedMappingSlot(key common.Address, baseSlot common.Hash) common.Hash {
	keyPadded := common.LeftPadBytes(key.Bytes(), 32)
	mapBaseSlotPadded := common.LeftPadBytes(baseSlot.Bytes(), 32)
	combined := append(keyPadded, mapBaseSlotPadded...)
	return crypto.Keccak256Hash(combined)
}
