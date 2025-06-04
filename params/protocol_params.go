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
	"os"

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
	// Temp get from env var for testing
	GasStationAddress = common.HexToAddress(os.Getenv("GAS_STATION"))
	GasStationCode    = common.FromHex("6080604052600436106101c8575f3560e01c806399c6066c116100f2578063c55b6bb711610092578063e559afd911610062578063e559afd9146107c3578063e73a914c146107e2578063f6e4b62b14610801578063fac2c62114610820575f80fd5b8063c55b6bb7146106ee578063d124d1bc1461070d578063d7e5fbf31461073e578063d9ba32fc1461075d575f80fd5b8063ad3e080a116100cd578063ad3e080a1461062e578063b6b352721461064d578063c375c2ef1461066c578063c3c5a5471461068b575f80fd5b806399c6066c146105865780639e4f8ab8146105a55780639f8a13d7146105c6575f80fd5b80633b66e9f61161016857806364efb22b1161013857806364efb22b1461040e57806369dc9ff3146104775780637901868e14610548578063871ff40514610567575f80fd5b80633b66e9f6146103035780634162169f146103665780634782f779146103d05780635e35359e146103ef575f80fd5b806315ea16ad116101a357806315ea16ad146102695780631c5d647c146102a65780632ce962cf146102c5578063368da168146102e4575f80fd5b8063108f5c69146101d3578063139e0aa7146101f457806314695ea414610207575f80fd5b366101cf57005b5f80fd5b3480156101de575f80fd5b506101f26101ed366004612e30565b61083f565b005b6101f2610202366004612ea8565b610a4f565b348015610212575f80fd5b50610254610221366004612ed2565b5f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db02602052604090205460ff1690565b60405190151581526020015b60405180910390f35b348015610274575f80fd5b507fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db03545b604051908152602001610260565b3480156102b1575f80fd5b506101f26102c0366004612ef6565b610cb6565b3480156102d0575f80fd5b506101f26102df366004612f24565b610e02565b3480156102ef575f80fd5b506101f26102fe366004612ea8565b611000565b34801561030e575f80fd5b5061029861031d366004612f50565b73ffffffffffffffffffffffffffffffffffffffff165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604090206001015490565b348015610371575f80fd5b507fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610260565b3480156103db575f80fd5b506101f26103ea366004612ea8565b611200565b3480156103fa575f80fd5b506101f2610409366004612f72565b611352565b348015610419575f80fd5b506103ab610428366004612f50565b73ffffffffffffffffffffffffffffffffffffffff9081165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db0160205260409020546201000090041690565b348015610482575f80fd5b50610501610491366004612f50565b73ffffffffffffffffffffffffffffffffffffffff9081165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604090208054600182015460029092015460ff808316956101008404821695620100009094049093169392911690565b604080519515158652931515602086015273ffffffffffffffffffffffffffffffffffffffff9092169284019290925260608301919091521515608082015260a001610260565b348015610553575f80fd5b506101f2610562366004612fb0565b611672565b348015610572575f80fd5b506101f2610581366004612ea8565b6118bc565b348015610591575f80fd5b506101f26105a0366004612ea8565b6119d7565b3480156105b0575f80fd5b506105b9611ac1565b604051610260919061301e565b3480156105d1575f80fd5b506102546105e0366004612f50565b73ffffffffffffffffffffffffffffffffffffffff165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db016020526040902054610100900460ff1690565b348015610639575f80fd5b506101f2610648366004613061565b611bcf565b348015610658575f80fd5b506102546106673660046130e2565b611d6e565b348015610677575f80fd5b506101f2610686366004612f50565b611e21565b348015610696575f80fd5b506102546106a5366004612f50565b73ffffffffffffffffffffffffffffffffffffffff165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604090205460ff1690565b3480156106f9575f80fd5b506101f26107083660046130e2565b611f53565b348015610718575f80fd5b5061072c610727366004612ed2565b6121a0565b6040516102609695949392919061310e565b348015610749575f80fd5b506101f26107583660046130e2565b6122b6565b348015610768575f80fd5b50610254610777366004612f50565b73ffffffffffffffffffffffffffffffffffffffff165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604090206002015460ff1690565b3480156107ce575f80fd5b506101f26107dd366004613061565b6124ee565b3480156107ed575f80fd5b506101f26107fc366004612f50565b612692565b34801561080c575f80fd5b506101f261081b366004612f24565b6127e6565b34801561082b575f80fd5b506101f261083a366004612f50565b612957565b7fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db005473ffffffffffffffffffffffffffffffffffffffff1633146108af576040517f4ead1a5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8590036108e9576040517f8dad8de600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612710811115610925576040517fe05f723400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8781527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db02602052604090206001810180546109609061319c565b90505f0361099a576040517fbdc474c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600181016109a9878983613265565b5060028101859055600381018490556004810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85161790556005810182905560405188907f73d628d7a9f63d75ab3f23c4bf349bfec022e61cc2ad8dc72f7ca093b45723e890610a3d908a908a908a908a908a908a9061337b565b60405180910390a25050505050505050565b73ffffffffffffffffffffffffffffffffffffffff82165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db016020526040902054829060ff16610ace576040517faba4733900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8281527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db0260205260409020600181018054610b099061319c565b90505f03610b43576040517fbdc474c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805460ff16610b7e576040517fd1d5af5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600481015473ffffffffffffffffffffffffffffffffffffffff16610bab57610ba681612a89565b610bec565b3415610be3576040517ffbccebae00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610bec81612b12565b600381015473ffffffffffffffffffffffffffffffffffffffff85165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604081206001018054909190610c47908490613428565b92505081905550828473ffffffffffffffffffffffffffffffffffffffff167f7852f393fd6a99c61648e39af92ae0e784b77281fc2af871edce1b51304ecd7c83600301548460020154604051610ca8929190918252602082015260400190565b60405180910390a350505050565b7fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db005473ffffffffffffffffffffffffffffffffffffffff163314610d26576040517f4ead1a5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8281527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db0260205260409020600181018054610d619061319c565b90505f03610d9b576040517fbdc474c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016821515908117825560405190815283907f10ae08733732b5e10d63d501510950b2a5967607149b3608881ecde96515780c906020015b60405180910390a2505050565b73ffffffffffffffffffffffffffffffffffffffff8083165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604090205483917fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db0091620100009004163314801590610e985750805473ffffffffffffffffffffffffffffffffffffffff163314155b15610ecf576040517fea8e4eb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff84165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db016020526040902054849060ff16610f4e576040517faba4733900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff85165f8181527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101008915159081029190911790915591519182527fa5ab8b72c18a722b7e92b557d227ba48dc2985b22fce6d0f95804be26703b595910160405180910390a25050505050565b7fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db005473ffffffffffffffffffffffffffffffffffffffff163314611070576040517f4ead1a5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db016020526040902054829060ff166110ef576040517faba4733900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604090206001015482811015611170576040517f43fb945300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61117a838261343b565b73ffffffffffffffffffffffffffffffffffffffff85165f8181527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db016020908152604091829020600101939093555185815290917f30a9d8d098632f590e4953b6171a6c999d2b1c4170ebde38136c9e27e6976b8191015b60405180910390a250505050565b7fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db005473ffffffffffffffffffffffffffffffffffffffff163314611270576040517f4ead1a5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff81166112be576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b475f83156112cc57836112ce565b815b90508181111561130a576040517f43fb945300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405173ffffffffffffffffffffffffffffffffffffffff86169082156108fc029083905f818181858888f1935050505015801561134a573d5f803e3d5ffd5b505050505050565b7fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db005473ffffffffffffffffffffffffffffffffffffffff1633146113c2576040517f4ead1a5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff8116611410576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff84166114bf57475f8315611439578361143b565b815b905081811115611477576040517f43fb945300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405173ffffffffffffffffffffffffffffffffffffffff86169082156108fc029083905f818181858888f193505050501580156114b7573d5f803e3d5ffd5b50505061166c565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015284905f9073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa15801561152b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061154f919061344e565b90505f841561155e5784611560565b815b90508181111561159c576040517f43fb945300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526024820183905284169063a9059cbb906044016020604051808303815f875af115801561160e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116329190613465565b611668576040517f045c4b0200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505b50505050565b7fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db005473ffffffffffffffffffffffffffffffffffffffff1633146116e2576040517f4ead1a5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f85900361171c576040517f8dad8de600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612710811115611758576040517fe05f723400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db03545f8181527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811782557fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db009291908101611802898b83613265565b506002810187905560038082018790556004820180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff88161790556005820185905583018054905f61186a83613480565b9190505550817f85855a4353e16703440df33dd6903f8689955fe665d2ed5b918f7a272286c8b98a8a8a8a8a8a6040516118a99695949392919061337b565b60405180910390a2505050505050505050565b7fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db005473ffffffffffffffffffffffffffffffffffffffff16331461192c576040517f4ead1a5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604081206001018054839290611982908490613428565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316907fed46984c46e11f42ec323727ba7d99dc16be2d248a8aaa8982d492688497f09d906020015b60405180910390a25050565b7fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db005473ffffffffffffffffffffffffffffffffffffffff163314611a47576040517f4ead1a5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f8181527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602090815260409182902060010184905590518381527fc2748283b871105da37ea4bdc2cc08eff4b3b0f472f66f2728cdf4a1b845ef7791016119cb565b60607fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db005f60015b8260030154811015611b22575f81815260028401602052604090205460ff1615611b1a5781611b1681613480565b9250505b600101611ae8565b505f8167ffffffffffffffff811115611b3d57611b3d6131ed565b604051908082528060200260200182016040528015611b66578160200160208202803683370190505b5090505f60015b8460030154811015611bc5575f81815260028601602052604090205460ff1615611bbd5780838381518110611ba457611ba46134b7565b602090810291909101015281611bb981613480565b9250505b600101611b6d565b5090949350505050565b73ffffffffffffffffffffffffffffffffffffffff8084165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604090205484917fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db0091620100009004163314801590611c655750805473ffffffffffffffffffffffffffffffffffffffff163314155b15611c9c576040517fea8e4eb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8381101561134a5773ffffffffffffffffffffffffffffffffffffffff86165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db016020526040812060030181878785818110611cff57611cff6134b7565b9050602002016020810190611d149190612f50565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040015f2080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055600101611c9e565b73ffffffffffffffffffffffffffffffffffffffff82165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604081206002015460ff161580611e18575073ffffffffffffffffffffffffffffffffffffffff8381165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db0160209081526040808320938616835260039093019052205460ff165b90505b92915050565b7fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db005473ffffffffffffffffffffffffffffffffffffffff163314611e91576040517f4ead1a5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f8181527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604080822080547fffffffffffffffffffff000000000000000000000000000000000000000000001681556001810183905560020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517f8d30d41865a0b811b9545d879520d2dde9f4cc49e4241f486ad9752bc904b5659190a250565b73ffffffffffffffffffffffffffffffffffffffff8083165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604090205483917fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db0091620100009004163314801590611fe95750805473ffffffffffffffffffffffffffffffffffffffff163314155b15612020576040517fea8e4eb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff84165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db016020526040902054849060ff1661209f576040517faba4733900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff81166120ed576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8681165f8181527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db0160205260408082208054620100008b87168181027fffffffffffffffffffff0000000000000000000000000000000000000000ffff84161790935592519290049094169392849290917f4eb572e99196bed0270fbd5b17a948e19c3f50a97838cb0d2a75a823ff8e6c509190a450505050505050565b5f606081808080807fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db005f89815260029182016020526040902080549181015460038201546004830154600584015460018501805495975060ff909616959473ffffffffffffffffffffffffffffffffffffffff9092169185906122229061319c565b80601f016020809104026020016040519081016040528092919081815260200182805461224e9061319c565b80156122995780601f1061227057610100808354040283529160200191612299565b820191905f5260205f20905b81548152906001019060200180831161227c57829003601f168201915b505050505094509650965096509650965096505091939550919395565b8173ffffffffffffffffffffffffffffffffffffffff8116612304576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff8116612352576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff84165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604090205460ff16156123d0576040517f3a81d6fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b833b5f81900361240c576040517f6eefed2000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8581165f8181527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604080822080546101017fffffffffffffffffffff0000000000000000000000000000000000000000000090911662010000968b169687021717815560018082018490556002820180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690911790559051909392917f768fb430a0d4b201cb764ab221c316dd14d8babf2e4b2348e05964c6565318b691a3505050505050565b73ffffffffffffffffffffffffffffffffffffffff8084165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604090205484917fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db00916201000090041633148015906125845750805473ffffffffffffffffffffffffffffffffffffffff163314155b156125bb576040517fea8e4eb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8381101561134a5773ffffffffffffffffffffffffffffffffffffffff86165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db0160205260408120600191600390910190878785818110612623576126236134b7565b90506020020160208101906126389190612f50565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040015f2080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790556001016125bd565b7fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db005473ffffffffffffffffffffffffffffffffffffffff163314612702576040517f4ead1a5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff8116612750576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db00805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f0429168a83556e356cd18563753346b9c9567cbf0fbea148d40aeb84a76cc5b9905f90a3505050565b73ffffffffffffffffffffffffffffffffffffffff8083165f9081527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604090205483917fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db009162010000900416331480159061287c5750805473ffffffffffffffffffffffffffffffffffffffff163314155b156128b3576040517fea8e4eb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff84165f8181527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602090815260409182902060020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001687151590811790915591519182527f8daaf060c3306c38e068a75c054bf96ecd85a3db1252712c4d93632744c42e0d91016111f2565b7fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db005473ffffffffffffffffffffffffffffffffffffffff1633146129c7576040517f4ead1a5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f8181527fc2eaf2cedf9e23687c6eb7c4717aa3eacbd015cc86eaad3f51aae2d3c955db01602052604080822080547fffffffffffffffffffff000000000000000000000000000000000000000000001681556001810183905560020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517f3475b9891ecf29e996feed01eeb42a860ec225283a439d214ffaeac5e006be7d9190a250565b8060020154341015612ac7576040517fcd1c886700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060020154341115612b0f57600281015433906108fc90612ae8903461343b565b6040518115909202915f818181858888f19350505050158015612b0d573d5f803e3d5ffd5b505b50565b60048181015460028301546040517fdd62ed3e000000000000000000000000000000000000000000000000000000008152339381019390935230602484015273ffffffffffffffffffffffffffffffffffffffff90911691829063dd62ed3e90604401602060405180830381865afa158015612b90573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612bb4919061344e565b1015612bec576040517fcd1c886700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028201546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481019190915273ffffffffffffffffffffffffffffffffffffffff8216906323b872dd906064016020604051808303815f875af1158015612c68573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c8c9190613465565b612cc2576040517f045c4b0200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600582015415612b0d575f61271083600501548460020154612ce491906134e4565b612cee91906134fb565b6004808501546040517f42966c6800000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff16916342966c6891612d4a9185910190815260200190565b5f604051808303815f87803b158015612d61575f80fd5b505af1925050508015612d72575060015b15612dc557600483015460405182815273ffffffffffffffffffffffffffffffffffffffff909116907ffd38818f5291bf0bb3a2a48aadc06ba8757865d1dabd804585338aab3009dcb690602001610df5565b505050565b5f8083601f840112612dda575f80fd5b50813567ffffffffffffffff811115612df1575f80fd5b602083019150836020828501011115612e08575f80fd5b9250929050565b73ffffffffffffffffffffffffffffffffffffffff81168114612b0f575f80fd5b5f805f805f805f60c0888a031215612e46575f80fd5b87359650602088013567ffffffffffffffff811115612e63575f80fd5b612e6f8a828b01612dca565b90975095505060408801359350606088013592506080880135612e9181612e0f565b8092505060a0880135905092959891949750929550565b5f8060408385031215612eb9575f80fd5b8235612ec481612e0f565b946020939093013593505050565b5f60208284031215612ee2575f80fd5b5035919050565b8015158114612b0f575f80fd5b5f8060408385031215612f07575f80fd5b823591506020830135612f1981612ee9565b809150509250929050565b5f8060408385031215612f35575f80fd5b8235612f4081612e0f565b91506020830135612f1981612ee9565b5f60208284031215612f60575f80fd5b8135612f6b81612e0f565b9392505050565b5f805f60608486031215612f84575f80fd5b8335612f8f81612e0f565b92506020840135612f9f81612e0f565b929592945050506040919091013590565b5f805f805f8060a08789031215612fc5575f80fd5b863567ffffffffffffffff811115612fdb575f80fd5b612fe789828a01612dca565b9097509550506020870135935060408701359250606087013561300981612e0f565b80925050608087013590509295509295509295565b602080825282518282018190525f9190848201906040850190845b8181101561305557835183529284019291840191600101613039565b50909695505050505050565b5f805f60408486031215613073575f80fd5b833561307e81612e0f565b9250602084013567ffffffffffffffff8082111561309a575f80fd5b818601915086601f8301126130ad575f80fd5b8135818111156130bb575f80fd5b8760208260051b85010111156130cf575f80fd5b6020830194508093505050509250925092565b5f80604083850312156130f3575f80fd5b82356130fe81612e0f565b91506020830135612f1981612e0f565b861515815260c060208201525f86518060c0840152806020890160e085015e5f60e0828501015260e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505085604083015284606083015273ffffffffffffffffffffffffffffffffffffffff841660808301528260a0830152979650505050505050565b600181811c908216806131b057607f821691505b6020821081036131e7577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b601f821115612dc557805f5260205f20601f840160051c8101602085101561323f5750805b601f840160051c820191505b8181101561325e575f815560010161324b565b5050505050565b67ffffffffffffffff83111561327d5761327d6131ed565b6132918361328b835461319c565b8361321a565b5f601f8411600181146132e1575f85156132ab5750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b17835561325e565b5f838152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08716915b8281101561332e578685013582556020948501946001909201910161330e565b5086821015613369577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b60a081528560a0820152858760c08301375f60c087830101525f60c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f890116830101905085602083015284604083015273ffffffffffffffffffffffffffffffffffffffff84166060830152826080830152979650505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820180821115611e1b57611e1b6133fb565b81810381811115611e1b57611e1b6133fb565b5f6020828403121561345e575f80fd5b5051919050565b5f60208284031215613475575f80fd5b8151612f6b81612ee9565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036134b0576134b06133fb565b5060010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b8082028115828204841417611e1b57611e1b6133fb565b5f8261352e577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b50049056fea164736f6c6343000819000a")
)
