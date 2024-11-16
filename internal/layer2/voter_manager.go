package layer2

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/utils"
)

type ContractVotingManager interface {
	Proposer() common.Address
	GenerateVerifyTaskUnSignMsg(contractAddress common.Address, calldata []byte, taskID *big.Int) common.Hash
}

func (l *Layer2Listener) ContractVotingManager() *contracts.VotingManagerContract {
	return l.contractVotingManager
}

func (l *Layer2Listener) Proposer() common.Address {
	proposer, err := l.contractVotingManager.NextSubmitter(nil)
	utils.Assert(err)

	return proposer
}

func (l *Layer2Listener) GenerateVerifyTaskUnSignMsg(contractAddress common.Address, calldata []byte, taskID *big.Int) common.Hash {
	nonce, err := l.contractVotingManager.TssNonce(nil)
	utils.Assert(err)
	nonce.Add(nonce, big.NewInt(1))
	encodeData, err := utils.AbiEncodePacked(nonce, contractAddress, calldata, taskID)
	utils.Assert(err)

	return crypto.Keccak256Hash(encodeData)
}
