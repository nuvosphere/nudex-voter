package contracts

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// FilterAddressRegistered is a free log retrieval operation binding the contract event 0x0ab661710c67363885e0e51920050375aff9dcd587adf3e2e468e060ee8f0e1e.
//
// Solidity: event AddressRegistered(address userAddr, uint256 indexed account, uint8 indexed chain, uint256 indexed index, string newAddress)

// FilterDepositRecorded is a free log retrieval operation binding the contract event 0xc81b018d055616352576702d0318bf7fc5c5b37693d9d4555113e2490d87dd80.
//
// Solidity: event DepositRecorded(string depositAddress, bytes32 indexed ticker, bytes32 indexed chainId, uint256 amount)

// FilterWithdrawalRecorded is a free log retrieval operation binding the contract event 0x2afe20970cc53fcbca49f0fd13ca943d027d7636ad0d9ac543a995a0cd03c9ec.
//
// Solidity: event WithdrawalRecorded(string depositAddress, bytes32 indexed ticker, bytes32 indexed chainId, uint256 amount)

// FilterParticipantsReset is a free log retrieval operation binding the contract event 0x32e9d8d19fb1e71c8dc610e5f45fd7f1e2f81babf8ea90e267475a708e09c35e.
//
// Solidity: event ParticipantsReset(address[] indexed participants)

// FilterTaskSubmitted is a free log retrieval operation binding the contract event 0x7c6cba37f838a9f6cd45be5dbe20a2a6c0a373fcb738333fbc39ab558183576f.
//
// Solidity: event TaskSubmitted(uint64 indexed taskId, address indexed submitter, bytes data)

// FilterTaskUpdated is a free log retrieval operation binding the contract event 0x30a99b2ffff1813c032a6b15bb8a15c2c3d1e9bc6dcb5f5cd80238514e86f364.
//
// Solidity: event TaskUpdated(uint64 indexed taskId, address indexed submitter, uint8 indexed state, uint256 updateTime, bytes32 txHash, bytes result)
var (
	SubmitterRotationRequestedTopic = crypto.Keccak256Hash([]byte(`SubmitterRotationRequested(address,address)`))             // VotingManagerUpgradeable
	SubmitterChosenTopic            = crypto.Keccak256Hash([]byte(`SubmitterChosen(address)`))                                // VotingManagerUpgradeable
	TaskSubmittedTopic              = crypto.Keccak256Hash([]byte(`TaskSubmitted(uint64,address,bytes)`))                     // TaskManager
	TaskUpdatedTopic                = crypto.Keccak256Hash([]byte(`TaskUpdated(uint64,address,uint8,uint256,bytes32,bytes)`)) // TaskManager
	AddressRegisteredTopic          = crypto.Keccak256Hash([]byte(`AddressRegistered(address,uint256,uint8,uint256,string)`)) // AccountManagerContract
	ParticipantAddedTopic           = crypto.Keccak256Hash([]byte(`ParticipantAdded(address)`))                               // IParticipantManager
	ParticipantRemovedTopic         = crypto.Keccak256Hash([]byte(`ParticipantRemoved(address)`))                             // IParticipantManager
	ParticipantsResetTopic          = crypto.Keccak256Hash([]byte(`ParticipantsReset(address[])`))                            // IParticipantManager
	DepositRecordedTopic            = crypto.Keccak256Hash([]byte(`DepositRecorded(string,bytes32,bytes32,uint256)`))         // DepositManagerContract
	WithdrawalRecordedTopic         = crypto.Keccak256Hash([]byte(`WithdrawalRecorded(string,bytes32,bytes32,uint256)`))      // DepositManagerContract

	WalletCreationRequestTopic = crypto.Keccak256Hash([]byte(`WalletCreationRequest(uint8,uint8,uint32,uint8,uint8)`))                                                    // TopicPayloadContract
	DepositRequestTopic        = crypto.Keccak256Hash([]byte(`DepositRequest(uint8,uint8,string,uint64,uint8,uint32,uint64,string,string,string,uint8,uint8)`))           // TopicPayloadContract
	WithdrawalRequestTopic     = crypto.Keccak256Hash([]byte(`WithdrawalRequest(uint8,uint8,string,uint64,uint8,uint32,uint64,string,string,string,uint8,uint8,uint64)`)) // TopicPayloadContract
	WalletCreationResultTopic  = crypto.Keccak256Hash([]byte(`WalletCreationResult(uint8,bool,uint8,string)`))                                                            // TopicPayloadContract
	DepositResultTopic         = crypto.Keccak256Hash([]byte(`DepositResult(uint8,bool,uint8)`))                                                                          // TopicPayloadContract
	WithdrawalResultTopic      = crypto.Keccak256Hash([]byte(`WithdrawalResult(uint8,bool,uint8)`))                                                                       // TopicPayloadContract
)

var Topics = [][]common.Hash{
	{SubmitterRotationRequestedTopic},
	{SubmitterChosenTopic},
	{TaskSubmittedTopic},
	{TaskUpdatedTopic},
	{AddressRegisteredTopic},
	{ParticipantRemovedTopic},
	{ParticipantAddedTopic},
	{DepositRecordedTopic},
	{WithdrawalRecordedTopic},
}
