package layer2

import "github.com/ethereum/go-ethereum/crypto"

var (
	SubmitterChosenTopic    = crypto.Keccak256Hash([]byte(`SubmitterChosen(address)`))                                 //VotingManagerUpgradeable
	TaskSubmittedTopic      = crypto.Keccak256Hash([]byte(`TaskSubmitted(uint256,bytes,address)`))                     //INuDexOperations
	AddressRegisteredTopic  = crypto.Keccak256Hash([]byte(`AddressRegistered(address,uint256,uint8,uint256,address)`)) //AccountManagerContract
	ParticipantAddedTopic   = crypto.Keccak256Hash([]byte(`ParticipantAdded(address)`))                                //IParticipantManager
	ParticipantRemovedTopic = crypto.Keccak256Hash([]byte(`ParticipantRemoved(address)`))                              //IParticipantManager
)
