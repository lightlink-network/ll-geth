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

package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// The base fee portion of the transaction fee accumulates at this predeploy
	OptimismBaseFeeRecipient = common.HexToAddress("0x4200000000000000000000000000000000000019")
	// The L1 portion of the transaction fee accumulates at this predeploy
	OptimismL1FeeRecipient = common.HexToAddress("0x420000000000000000000000000000000000001A")
	// The L2 withdrawals contract predeploy address
	OptimismL2ToL1MessagePasser = common.HexToAddress("0x4200000000000000000000000000000000000016")
	// The operator fee portion of the transaction fee accumulates at this predeploy
	OptimismOperatorFeeRecipient = common.HexToAddress("0x420000000000000000000000000000000000001B")
)

const (
	GasLimitBoundDivisor uint64 = 1024               // The bound divisor of the gas limit, used in update calculations.
	MinGasLimit          uint64 = 5000               // Minimum the gas limit may ever be.
	MaxGasLimit          uint64 = 0x7fffffffffffffff // Maximum the gas limit (2^63-1).
	GenesisGasLimit      uint64 = 4712388            // Gas limit of the Genesis block.

	MaximumExtraDataSize  uint64 = 32    // Maximum size extra data may be after Genesis.
	ExpByteGas            uint64 = 10    // Times ceil(log256(exponent)) for the EXP instruction.
	SloadGas              uint64 = 50    // Multiplied by the number of 32-byte words that are copied (round up) for any *COPY operation and added.
	CallValueTransferGas  uint64 = 9000  // Paid for CALL when the value transfer is non-zero.
	CallNewAccountGas     uint64 = 25000 // Paid for CALL when the destination address didn't exist prior.
	TxGas                 uint64 = 21000 // Per transaction not creating a contract. NOTE: Not payable on data of calls between transactions.
	TxGasContractCreation uint64 = 53000 // Per transaction that creates a contract. NOTE: Not payable on data of calls between transactions.
	TxDataZeroGas         uint64 = 4     // Per byte of data attached to a transaction that equals zero. NOTE: Not payable on data of calls between transactions.
	QuadCoeffDiv          uint64 = 512   // Divisor for the quadratic particle of the memory cost equation.
	LogDataGas            uint64 = 8     // Per byte in a LOG* operation's data.
	CallStipend           uint64 = 2300  // Free gas given at beginning of call.

	Keccak256Gas     uint64 = 30 // Once per KECCAK256 operation.
	Keccak256WordGas uint64 = 6  // Once per word of the KECCAK256 operation's data.
	InitCodeWordGas  uint64 = 2  // Once per word of the init code when creating a contract.

	SstoreSetGas    uint64 = 20000 // Once per SSTORE operation.
	SstoreResetGas  uint64 = 5000  // Once per SSTORE operation if the zeroness changes from zero.
	SstoreClearGas  uint64 = 5000  // Once per SSTORE operation if the zeroness doesn't change.
	SstoreRefundGas uint64 = 15000 // Once per SSTORE operation if the zeroness changes to zero.

	NetSstoreNoopGas  uint64 = 200   // Once per SSTORE operation if the value doesn't change.
	NetSstoreInitGas  uint64 = 20000 // Once per SSTORE operation from clean zero.
	NetSstoreCleanGas uint64 = 5000  // Once per SSTORE operation from clean non-zero.
	NetSstoreDirtyGas uint64 = 200   // Once per SSTORE operation from dirty.

	NetSstoreClearRefund      uint64 = 15000 // Once per SSTORE operation for clearing an originally existing storage slot
	NetSstoreResetRefund      uint64 = 4800  // Once per SSTORE operation for resetting to the original non-zero value
	NetSstoreResetClearRefund uint64 = 19800 // Once per SSTORE operation for resetting to the original zero value

	SstoreSentryGasEIP2200            uint64 = 2300  // Minimum gas required to be present for an SSTORE call, not consumed
	SstoreSetGasEIP2200               uint64 = 20000 // Once per SSTORE operation from clean zero to non-zero
	SstoreResetGasEIP2200             uint64 = 5000  // Once per SSTORE operation from clean non-zero to something else
	SstoreClearsScheduleRefundEIP2200 uint64 = 15000 // Once per SSTORE operation for clearing an originally existing storage slot

	ColdAccountAccessCostEIP2929 = uint64(2600) // COLD_ACCOUNT_ACCESS_COST
	ColdSloadCostEIP2929         = uint64(2100) // COLD_SLOAD_COST
	WarmStorageReadCostEIP2929   = uint64(100)  // WARM_STORAGE_READ_COST

	// In EIP-2200: SstoreResetGas was 5000.
	// In EIP-2929: SstoreResetGas was changed to '5000 - COLD_SLOAD_COST'.
	// In EIP-3529: SSTORE_CLEARS_SCHEDULE is defined as SSTORE_RESET_GAS + ACCESS_LIST_STORAGE_KEY_COST
	// Which becomes: 5000 - 2100 + 1900 = 4800
	SstoreClearsScheduleRefundEIP3529 uint64 = SstoreResetGasEIP2200 - ColdSloadCostEIP2929 + TxAccessListStorageKeyGas

	JumpdestGas   uint64 = 1     // Once per JUMPDEST operation.
	EpochDuration uint64 = 30000 // Duration between proof-of-work epochs.

	CreateDataGas         uint64 = 200   //
	CallCreateDepth       uint64 = 1024  // Maximum depth of call/create stack.
	ExpGas                uint64 = 10    // Once per EXP instruction
	LogGas                uint64 = 375   // Per LOG* operation.
	CopyGas               uint64 = 3     //
	StackLimit            uint64 = 1024  // Maximum size of VM stack allowed.
	TierStepGas           uint64 = 0     // Once per operation, for a selection of them.
	LogTopicGas           uint64 = 375   // Multiplied by the * of the LOG*, per LOG transaction. e.g. LOG0 incurs 0 * c_txLogTopicGas, LOG4 incurs 4 * c_txLogTopicGas.
	CreateGas             uint64 = 32000 // Once per CREATE operation & contract-creation transaction.
	Create2Gas            uint64 = 32000 // Once per CREATE2 operation
	CreateNGasEip4762     uint64 = 1000  // Once per CREATEn operations post-verkle
	SelfdestructRefundGas uint64 = 24000 // Refunded following a selfdestruct operation.
	MemoryGas             uint64 = 3     // Times the address of the (highest referenced byte in memory + 1). NOTE: referencing happens on read, write and in instructions such as RETURN and CALL.

	TxDataNonZeroGasFrontier  uint64 = 68    // Per byte of data attached to a transaction that is not equal to zero. NOTE: Not payable on data of calls between transactions.
	TxDataNonZeroGasEIP2028   uint64 = 16    // Per byte of non zero data attached to a transaction after EIP 2028 (part in Istanbul)
	TxTokenPerNonZeroByte     uint64 = 4     // Token cost per non-zero byte as specified by EIP-7623.
	TxCostFloorPerToken       uint64 = 10    // Cost floor per byte of data as specified by EIP-7623.
	TxAccessListAddressGas    uint64 = 2400  // Per address specified in EIP 2930 access list
	TxAccessListStorageKeyGas uint64 = 1900  // Per storage key specified in EIP 2930 access list
	TxAuthTupleGas            uint64 = 12500 // Per auth tuple code specified in EIP-7702

	// These have been changed during the course of the chain
	CallGasFrontier              uint64 = 40  // Once per CALL operation & message call transaction.
	CallGasEIP150                uint64 = 700 // Static portion of gas for CALL-derivates after EIP 150 (Tangerine)
	BalanceGasFrontier           uint64 = 20  // The cost of a BALANCE operation
	BalanceGasEIP150             uint64 = 400 // The cost of a BALANCE operation after Tangerine
	BalanceGasEIP1884            uint64 = 700 // The cost of a BALANCE operation after EIP 1884 (part of Istanbul)
	ExtcodeSizeGasFrontier       uint64 = 20  // Cost of EXTCODESIZE before EIP 150 (Tangerine)
	ExtcodeSizeGasEIP150         uint64 = 700 // Cost of EXTCODESIZE after EIP 150 (Tangerine)
	SloadGasFrontier             uint64 = 50
	SloadGasEIP150               uint64 = 200
	SloadGasEIP1884              uint64 = 800  // Cost of SLOAD after EIP 1884 (part of Istanbul)
	SloadGasEIP2200              uint64 = 800  // Cost of SLOAD after EIP 2200 (part of Istanbul)
	ExtcodeHashGasConstantinople uint64 = 400  // Cost of EXTCODEHASH (introduced in Constantinople)
	ExtcodeHashGasEIP1884        uint64 = 700  // Cost of EXTCODEHASH after EIP 1884 (part in Istanbul)
	SelfdestructGasEIP150        uint64 = 5000 // Cost of SELFDESTRUCT post EIP 150 (Tangerine)

	// EXP has a dynamic portion depending on the size of the exponent
	ExpByteFrontier uint64 = 10 // was set to 10 in Frontier
	ExpByteEIP158   uint64 = 50 // was raised to 50 during Eip158 (Spurious Dragon)

	// Extcodecopy has a dynamic AND a static cost. This represents only the
	// static portion of the gas. It was changed during EIP 150 (Tangerine)
	ExtcodeCopyBaseFrontier uint64 = 20
	ExtcodeCopyBaseEIP150   uint64 = 700

	// CreateBySelfdestructGas is used when the refunded account is one that does
	// not exist. This logic is similar to call.
	// Introduced in Tangerine Whistle (Eip 150)
	CreateBySelfdestructGas uint64 = 25000

	DefaultBaseFeeChangeDenominator = 8          // Bounds the amount the base fee can change between blocks.
	DefaultElasticityMultiplier     = 2          // Bounds the maximum gas limit an EIP-1559 block may have.
	InitialBaseFee                  = 1000000000 // Initial base fee for EIP-1559 blocks.

	MaxCodeSize     = 24576           // Maximum bytecode to permit for a contract
	MaxInitCodeSize = 2 * MaxCodeSize // Maximum initcode to permit in a creation transaction and create instructions

	// Precompiled contract gas prices

	EcrecoverGas        uint64 = 3000 // Elliptic curve sender recovery gas price
	Sha256BaseGas       uint64 = 60   // Base price for a SHA256 operation
	Sha256PerWordGas    uint64 = 12   // Per-word price for a SHA256 operation
	Ripemd160BaseGas    uint64 = 600  // Base price for a RIPEMD160 operation
	Ripemd160PerWordGas uint64 = 120  // Per-word price for a RIPEMD160 operation
	IdentityBaseGas     uint64 = 15   // Base price for a data copy operation
	IdentityPerWordGas  uint64 = 3    // Per-work price for a data copy operation

	Bn256AddGasByzantium             uint64 = 500    // Byzantium gas needed for an elliptic curve addition
	Bn256AddGasIstanbul              uint64 = 150    // Gas needed for an elliptic curve addition
	Bn256ScalarMulGasByzantium       uint64 = 40000  // Byzantium gas needed for an elliptic curve scalar multiplication
	Bn256ScalarMulGasIstanbul        uint64 = 6000   // Gas needed for an elliptic curve scalar multiplication
	Bn256PairingBaseGasByzantium     uint64 = 100000 // Byzantium base price for an elliptic curve pairing check
	Bn256PairingBaseGasIstanbul      uint64 = 45000  // Base price for an elliptic curve pairing check
	Bn256PairingPerPointGasByzantium uint64 = 80000  // Byzantium per-point price for an elliptic curve pairing check
	Bn256PairingPerPointGasIstanbul  uint64 = 34000  // Per-point price for an elliptic curve pairing check

	Bn256PairingMaxInputSizeGranite uint64 = 112687 // Maximum input size for an elliptic curve pairing check

	Bls12381G1AddGas          uint64 = 375   // Price for BLS12-381 elliptic curve G1 point addition
	Bls12381G1MulGas          uint64 = 12000 // Price for BLS12-381 elliptic curve G1 point scalar multiplication
	Bls12381G2AddGas          uint64 = 600   // Price for BLS12-381 elliptic curve G2 point addition
	Bls12381G2MulGas          uint64 = 22500 // Price for BLS12-381 elliptic curve G2 point scalar multiplication
	Bls12381PairingBaseGas    uint64 = 37700 // Base gas price for BLS12-381 elliptic curve pairing check
	Bls12381PairingPerPairGas uint64 = 32600 // Per-point pair gas price for BLS12-381 elliptic curve pairing check
	Bls12381MapG1Gas          uint64 = 5500  // Gas price for BLS12-381 mapping field element to G1 operation
	Bls12381MapG2Gas          uint64 = 23800 // Gas price for BLS12-381 mapping field element to G2 operation

	Bls12381G1MulMaxInputSizeIsthmus   uint64 = 513760 // Maximum input size for BLS12-381 G1 multiple-scalar-multiply operation
	Bls12381G2MulMaxInputSizeIsthmus   uint64 = 488448 // Maximum input size for BLS12-381 G2 multiple-scalar-multiply operation
	Bls12381PairingMaxInputSizeIsthmus uint64 = 235008 // Maximum input size for BLS12-381 pairing check

	P256VerifyGas uint64 = 3450 // secp256r1 elliptic curve signature verifier gas price

	// The Refund Quotient is the cap on how much of the used gas can be refunded. Before EIP-3529,
	// up to half the consumed gas could be refunded. Redefined as 1/5th in EIP-3529
	RefundQuotient        uint64 = 2
	RefundQuotientEIP3529 uint64 = 5

	BlobTxBytesPerFieldElement         = 32      // Size in bytes of a field element
	BlobTxFieldElementsPerBlob         = 4096    // Number of field elements stored in a single data blob
	BlobTxBlobGasPerBlob               = 1 << 17 // Gas consumption of a single data blob (== blob byte size)
	BlobTxMinBlobGasprice              = 1       // Minimum gas price for data blobs
	BlobTxPointEvaluationPrecompileGas = 50000   // Gas price for the point evaluation precompile.

	HistoryServeWindow = 8192 // Number of blocks to serve historical block hashes for, EIP-2935.
)

// Bls12381G1MultiExpDiscountTable is the gas discount table for BLS12-381 G1 multi exponentiation operation
var Bls12381G1MultiExpDiscountTable = [128]uint64{1000, 949, 848, 797, 764, 750, 738, 728, 719, 712, 705, 698, 692, 687, 682, 677, 673, 669, 665, 661, 658, 654, 651, 648, 645, 642, 640, 637, 635, 632, 630, 627, 625, 623, 621, 619, 617, 615, 613, 611, 609, 608, 606, 604, 603, 601, 599, 598, 596, 595, 593, 592, 591, 589, 588, 586, 585, 584, 582, 581, 580, 579, 577, 576, 575, 574, 573, 572, 570, 569, 568, 567, 566, 565, 564, 563, 562, 561, 560, 559, 558, 557, 556, 555, 554, 553, 552, 551, 550, 549, 548, 547, 547, 546, 545, 544, 543, 542, 541, 540, 540, 539, 538, 537, 536, 536, 535, 534, 533, 532, 532, 531, 530, 529, 528, 528, 527, 526, 525, 525, 524, 523, 522, 522, 521, 520, 520, 519}

// Bls12381G2MultiExpDiscountTable is the gas discount table for BLS12-381 G2 multi exponentiation operation
var Bls12381G2MultiExpDiscountTable = [128]uint64{1000, 1000, 923, 884, 855, 832, 812, 796, 782, 770, 759, 749, 740, 732, 724, 717, 711, 704, 699, 693, 688, 683, 679, 674, 670, 666, 663, 659, 655, 652, 649, 646, 643, 640, 637, 634, 632, 629, 627, 624, 622, 620, 618, 615, 613, 611, 609, 607, 606, 604, 602, 600, 598, 597, 595, 593, 592, 590, 589, 587, 586, 584, 583, 582, 580, 579, 578, 576, 575, 574, 573, 571, 570, 569, 568, 567, 566, 565, 563, 562, 561, 560, 559, 558, 557, 556, 555, 554, 553, 552, 552, 551, 550, 549, 548, 547, 546, 545, 545, 544, 543, 542, 541, 541, 540, 539, 538, 537, 537, 536, 535, 535, 534, 533, 532, 532, 531, 530, 530, 529, 528, 528, 527, 526, 526, 525, 524, 524}

// Difficulty parameters.
var (
	DifficultyBoundDivisor = big.NewInt(2048)   // The bound divisor of the difficulty, used in the update calculations.
	GenesisDifficulty      = big.NewInt(131072) // Difficulty of the Genesis block.
	MinimumDifficulty      = big.NewInt(131072) // The minimum that the difficulty may ever be.
	DurationLimit          = big.NewInt(13)     // The decision boundary on the blocktime duration used to determine whether difficulty should go up or not.
)

// System contracts.
var (
	// SystemAddress is where the system-transaction is sent from as per EIP-4788
	SystemAddress = common.HexToAddress("0xfffffffffffffffffffffffffffffffffffffffe")

	// EIP-4788 - Beacon block root in the EVM
	BeaconRootsAddress = common.HexToAddress("0x000F3df6D732807Ef1319fB7B8bB8522d0Beac02")
	BeaconRootsCode    = common.FromHex("3373fffffffffffffffffffffffffffffffffffffffe14604d57602036146024575f5ffd5b5f35801560495762001fff810690815414603c575f5ffd5b62001fff01545f5260205ff35b5f5ffd5b62001fff42064281555f359062001fff015500")

	// EIP-2935 - Serve historical block hashes from state
	HistoryStorageAddress = common.HexToAddress("0x0000F90827F1C53a10cb7A02335B175320002935")
	HistoryStorageCode    = common.FromHex("3373fffffffffffffffffffffffffffffffffffffffe14604657602036036042575f35600143038111604257611fff81430311604257611fff9006545f5260205ff35b5f5ffd5b5f35611fff60014303065500")

	// EIP-7002 - Execution layer triggerable withdrawals
	WithdrawalQueueAddress = common.HexToAddress("0x00000961Ef480Eb55e80D19ad83579A64c007002")
	WithdrawalQueueCode    = common.FromHex("3373fffffffffffffffffffffffffffffffffffffffe1460cb5760115f54807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff146101f457600182026001905f5b5f82111560685781019083028483029004916001019190604d565b909390049250505036603814608857366101f457346101f4575f5260205ff35b34106101f457600154600101600155600354806003026004013381556001015f35815560010160203590553360601b5f5260385f601437604c5fa0600101600355005b6003546002548082038060101160df575060105b5f5b8181146101835782810160030260040181604c02815460601b8152601401816001015481526020019060020154807fffffffffffffffffffffffffffffffff00000000000000000000000000000000168252906010019060401c908160381c81600701538160301c81600601538160281c81600501538160201c81600401538160181c81600301538160101c81600201538160081c81600101535360010160e1565b910180921461019557906002556101a0565b90505f6002555f6003555b5f54807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14156101cd57505f5b6001546002828201116101e25750505f6101e8565b01600290035b5f555f600155604c025ff35b5f5ffd")

	// EIP-7251 - Increase the MAX_EFFECTIVE_BALANCE
	ConsolidationQueueAddress = common.HexToAddress("0x0000BBdDc7CE488642fb579F8B00f3a590007251")
	ConsolidationQueueCode    = common.FromHex("3373fffffffffffffffffffffffffffffffffffffffe1460d35760115f54807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1461019a57600182026001905f5b5f82111560685781019083028483029004916001019190604d565b9093900492505050366060146088573661019a573461019a575f5260205ff35b341061019a57600154600101600155600354806004026004013381556001015f358155600101602035815560010160403590553360601b5f5260605f60143760745fa0600101600355005b6003546002548082038060021160e7575060025b5f5b8181146101295782810160040260040181607402815460601b815260140181600101548152602001816002015481526020019060030154905260010160e9565b910180921461013b5790600255610146565b90505f6002555f6003555b5f54807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff141561017357505f5b6001546001828201116101885750505f61018e565b01600190035b5f555f6001556074025ff35b5f5ffd")

	// LLIP-2 - Gasless transactions
	GaslessRegistryAddress = common.HexToAddress("0x0000000000000000000000000000000000000099")
	GaslessRegistryCode    = common.FromHex("6080604052600436106100c1575f3560e01c8063b6b352721161007e578063d7e5fbf311610058578063d7e5fbf3146102a9578063e6f35204146102c5578063f6e4b62b14610301578063f8d3277d14610329576100c1565b8063b6b3527214610209578063c375c2ef14610245578063c3c5a5471461026d576100c1565b80633b66e9f6146100c55780635751824314610101578063679ce6811461012957806369dc9ff314610165578063871ff405146101a55780639f8a13d7146101cd575b5f80fd5b3480156100d0575f80fd5b506100eb60048036038101906100e69190610cb8565b610351565b6040516100f89190610cfb565b60405180910390f35b34801561010c575f80fd5b5061012760048036038101906101229190610d14565b610399565b005b348015610134575f80fd5b5061014f600480360381019061014a9190610d14565b61042f565b60405161015c9190610d6c565b60405180910390f35b348015610170575f80fd5b5061018b60048036038101906101869190610cb8565b6105d3565b60405161019c959493929190610d94565b60405180910390f35b3480156101b0575f80fd5b506101cb60048036038101906101c69190610e0f565b610647565b005b3480156101d8575f80fd5b506101f360048036038101906101ee9190610cb8565b6106a0565b6040516102009190610d6c565b60405180910390f35b348015610214575f80fd5b5061022f600480360381019061022a9190610d14565b6106f4565b60405161023c9190610d6c565b60405180910390f35b348015610250575f80fd5b5061026b60048036038101906102669190610cb8565b6107d8565b005b348015610278575f80fd5b50610293600480360381019061028e9190610cb8565b61087f565b6040516102a09190610d6c565b60405180910390f35b6102c360048036038101906102be9190610d14565b6108d2565b005b3480156102d0575f80fd5b506102eb60048036038101906102e69190610e0f565b610aa0565b6040516102f89190610d6c565b60405180910390f35b34801561030c575f80fd5b5061032760048036038101906103229190610e77565b610b6b565b005b348015610334575f80fd5b5061034f600480360381019061034a9190610d14565b610bc5565b005b5f805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600101549050919050565b60015f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505050565b5f805f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900460ff16610487575f90506105cd565b5f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0160019054906101000a900460ff166104df575f90506105cd565b5f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1680156105bb57505f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16155b156105c8575f90506105cd565b600190505b92915050565b5f602052805f5260405f205f91509050805f015f9054906101000a900460ff1690805f0160019054906101000a900460ff1690805f0160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002015f9054906101000a900460ff16905085565b805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206001015f8282546106959190610ee2565b925050819055505050565b5f805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0160019054906101000a900460ff169050919050565b5f805f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f9054906101000a900460ff1615806107d057505f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff165b905092915050565b5f808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8082015f6101000a81549060ff02191690555f820160016101000a81549060ff02191690555f820160026101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182015f9055600282015f6101000a81549060ff0219169055505050565b5f805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900460ff169050919050565b5f808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900460ff161561095d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095490610f6f565b60405180910390fd5b5f805f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090506001815f015f6101000a81548160ff0219169083151502179055506001815f0160016101000a81548160ff02191690831515021790555081815f0160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550670de0b6b3a764000081600101819055505f816002015f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f768fb430a0d4b201cb764ab221c316dd14d8babf2e4b2348e05964c6565318b660405160405180910390a3505050565b5f805f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090508281600101541015610af5575f915050610b65565b82816001015f828254610b089190610f8d565b925050819055508373ffffffffffffffffffffffffffffffffffffffff167f3c22f5db80c5276c536dd2c880cd698095e83b906c3ad64ffe1895eee8eb9f9f3285604051610b57929190610fc0565b60405180910390a260019150505b92915050565b805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002015f6101000a81548160ff0219169083151502179055505050565b5f805f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610c8782610c5e565b9050919050565b610c9781610c7d565b8114610ca1575f80fd5b50565b5f81359050610cb281610c8e565b92915050565b5f60208284031215610ccd57610ccc610c5a565b5b5f610cda84828501610ca4565b91505092915050565b5f819050919050565b610cf581610ce3565b82525050565b5f602082019050610d0e5f830184610cec565b92915050565b5f8060408385031215610d2a57610d29610c5a565b5b5f610d3785828601610ca4565b9250506020610d4885828601610ca4565b9150509250929050565b5f8115159050919050565b610d6681610d52565b82525050565b5f602082019050610d7f5f830184610d5d565b92915050565b610d8e81610c7d565b82525050565b5f60a082019050610da75f830188610d5d565b610db46020830187610d5d565b610dc16040830186610d85565b610dce6060830185610cec565b610ddb6080830184610d5d565b9695505050505050565b610dee81610ce3565b8114610df8575f80fd5b50565b5f81359050610e0981610de5565b92915050565b5f8060408385031215610e2557610e24610c5a565b5b5f610e3285828601610ca4565b9250506020610e4385828601610dfb565b9150509250929050565b610e5681610d52565b8114610e60575f80fd5b50565b5f81359050610e7181610e4d565b92915050565b5f8060408385031215610e8d57610e8c610c5a565b5b5f610e9a85828601610ca4565b9250506020610eab85828601610e63565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610eec82610ce3565b9150610ef783610ce3565b9250828201905080821115610f0f57610f0e610eb5565b5b92915050565b5f82825260208201905092915050565b7f616c7265616479207265676973746572656400000000000000000000000000005f82015250565b5f610f59601283610f15565b9150610f6482610f25565b602082019050919050565b5f6020820190508181035f830152610f8681610f4d565b9050919050565b5f610f9782610ce3565b9150610fa283610ce3565b9250828203905081811115610fba57610fb9610eb5565b5b92915050565b5f604082019050610fd35f830185610d85565b610fe06020830184610cec565b939250505056fea2646970667358221220a6934ae3a4fa48184868b0c2583329f45f74b948c5a9e2a570bb57963de1451564736f6c634300081a0033")
)
