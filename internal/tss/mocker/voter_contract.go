package mocker

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/types"
)

type VoterContractMocker struct{}

func (v VoterContractMocker) TssNonce() (*big.Int, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) Proposer() (common.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) NextSubmitter() (common.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) TssSigner() (common.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) LastSubmissionTime() (*big.Int, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) ForcedRotationWindow() (*big.Int, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) TaskCompletionThreshold() (*big.Int, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) EncodeVerifyAndCall(_target common.Address, _data []byte, _signature []byte) []byte {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) GenerateVerifyTaskUnSignMsg(contractAddress common.Address, calldata []byte, taskID *big.Int) (common.Hash, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) Participants() (types.Participants, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) IsParticipant(participant common.Address) (bool, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) GetRandomParticipant(participant common.Address) (common.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) GetLatestTask() (contracts.ITaskManagerTask, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) IsTaskCompleted(taskId *big.Int) (bool, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) GetUncompletedTasks() ([]contracts.ITaskManagerTask, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) EncodeSubmitTask(submitter common.Address, context []byte) []byte {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) NextTaskId() (*big.Int, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) Tasks(taskId *big.Int) (contracts.ITaskManagerTask, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) TaskSubmitter() (common.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) EncodeMarkTaskCompleted(taskId *big.Int, result []byte) []byte {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) EncodeRegisterNewAddress(_user common.Address, _account *big.Int, _chain uint8, _index *big.Int, _address string) []byte {
	// TODO implement me
	panic("implement me")
}

func (v VoterContractMocker) GetAddressRecord(opts *bind.CallOpts, _user common.Address, _account *big.Int, _chain uint8, _index *big.Int) (string, error) {
	// TODO implement me
	panic("implement me")
}
