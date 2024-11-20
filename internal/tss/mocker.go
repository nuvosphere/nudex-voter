package tss

import (
	"context"
	"encoding/json"
	"math/big"
	"slices"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/p2p"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper/testutil"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
)

type P2PMocker struct {
	typeBindEvent *sync.Map // MessageType:eventbus.Event
	bus           eventbus.Bus
}

func NewP2PMocker(bus eventbus.Bus) *P2PMocker {
	return &P2PMocker{
		bus:           bus,
		typeBindEvent: &sync.Map{},
	}
}

func (p *P2PMocker) Bind(msgType p2p.MessageType, event eventbus.Event) {
	p.typeBindEvent.Store(msgType, event)
}

func (p *P2PMocker) PublishMessage(ctx context.Context, msg any) error {
	data, err := json.Marshal(msg)
	utils.Assert(err)

	var receivedMsg p2p.Message[json.RawMessage]
	err = json.Unmarshal(data, &receivedMsg)
	utils.Assert(err)

	event, ok := p.typeBindEvent.Load(receivedMsg.MessageType)
	if ok {
		p.bus.Publish(event, receivedMsg)
	}

	return nil
}

func (p *P2PMocker) OnlinePeerCount() int {
	return testutil.TestPartyCount
}

func (p *P2PMocker) IsOnline(partyID string) bool {
	return true
}

type VoterContractMocker struct {
	nonce        big.Int
	proposer     common.Address
	participants types.Participants
}

func NewVoterContractMocker() *VoterContractMocker {
	return &VoterContractMocker{
		nonce: *big.NewInt(1),
	}
}

func (v *VoterContractMocker) TssNonce() (*big.Int, error) {
	return &v.nonce, nil
}

func (v *VoterContractMocker) SetProposer(p common.Address) {
	v.proposer = p
}

func (v *VoterContractMocker) Proposer() (common.Address, error) {
	return v.proposer, nil
}

func (v *VoterContractMocker) NextSubmitter() (common.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) TssSigner() (common.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) LastSubmissionTime() (*big.Int, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) ForcedRotationWindow() (*big.Int, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) TaskCompletionThreshold() (*big.Int, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) EncodeVerifyAndCall(_target common.Address, _data []byte, _signature []byte) []byte {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) GenerateVerifyTaskUnSignMsg(contractAddress common.Address, calldata []byte, taskID *big.Int) (common.Hash, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) SetParticipants(pp types.Participants) {
	v.participants = pp
}

func (v *VoterContractMocker) Participants() (types.Participants, error) {
	return v.participants, nil
}

func (v *VoterContractMocker) IsParticipant(participant common.Address) (bool, error) {
	return slices.Contains(v.participants, participant), nil
}

func (v *VoterContractMocker) GetRandomParticipant(participant common.Address) (common.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) GetLatestTask() (contracts.ITaskManagerTask, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) IsTaskCompleted(taskId *big.Int) (bool, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) GetUncompletedTasks() ([]contracts.ITaskManagerTask, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) EncodeSubmitTask(submitter common.Address, context []byte) []byte {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) NextTaskId() (*big.Int, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) Tasks(taskId *big.Int) (contracts.ITaskManagerTask, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) TaskSubmitter() (common.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) EncodeMarkTaskCompleted(taskId *big.Int, result []byte) []byte {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) EncodeRegisterNewAddress(_user common.Address, _account *big.Int, _chain uint8, _index *big.Int, _address string) []byte {
	// TODO implement me
	panic("implement me")
}

func (v *VoterContractMocker) GetAddressRecord(opts *bind.CallOpts, _user common.Address, _account *big.Int, _chain uint8, _index *big.Int) (string, error) {
	// TODO implement me
	panic("implement me")
}

//
//=== RUN   TestScheduler
//scheduler_test.go:110: submitters [0x3a818294ca1F3C27d7588b123Ec43F2546fa07f4 0x5091FC3cb4E4BB014141Aa41375d8Dd73b34AfA2 0x04d9389Cf937b1e6F2258d842e7237E955d6ab04 0xf6D37CE75dB465DcDb4c7097bEB9c1D46b171037 0x1D2cd50A3cF3c55a7982AD54F9f364C1e953Bc57]
//scheduler_test.go:117: index: 0 submitter: 0x3a818294ca1F3C27d7588b123Ec43F2546fa07f4
//time="2024-11-20T11:40:48+08:00" level=info msg="p2p loop started"
//time="2024-11-20T11:40:48+08:00" level=info msg="proposal loop started"
//scheduler_test.go:117: index: 1 submitter: 0x5091FC3cb4E4BB014141Aa41375d8Dd73b34AfA2
//time="2024-11-20T11:40:48+08:00" level=info msg="TSS keygen process started Candidate:0x3a818294ca1F3C27d7588b123Ec43F2546fa07f4proposer: 0x04d9389Cf937b1e6F2258d842e7237E955d6ab04"
//time="2024-11-20T11:40:48+08:00" level=info msg="p2p loop started"
//time="2024-11-20T11:40:48+08:00" level=info msg="proposal loop started"
//time="2024-11-20T11:40:48+08:00" level=info msg="TSS keygen process started Candidate:0x5091FC3cb4E4BB014141Aa41375d8Dd73b34AfA2proposer: 0x04d9389Cf937b1e6F2258d842e7237E955d6ab04"
//scheduler_test.go:117: index: 2 submitter: 0x04d9389Cf937b1e6F2258d842e7237E955d6ab04
//time="2024-11-20T11:40:48+08:00" level=info msg="p2p loop started"
//time="2024-11-20T11:40:48+08:00" level=info msg="proposal loop started"
//time="2024-11-20T11:40:48+08:00" level=info msg="TSS keygen process started leader:0x04d9389Cf937b1e6F2258d842e7237E955d6ab04proposer: 0x04d9389Cf937b1e6F2258d842e7237E955d6ab04"
//time="2024-11-20T11:40:48+08:00" level=debug msg="creating new local party"
//time="2024-11-20T11:40:48+08:00" level=debug msg="local party createdpartyID{0,4d9389cf937b1e6f2258d842e7237e955d6ab04-secp256k1}"
//time="2024-11-20T11:40:48+08:00" level=debug msg="Starting out/in message loop"
//time="2024-11-20T11:40:48+08:00" level=debug msg="waiting for next message...partyID: {0,4d9389cf937b1e6f2258d842e7237e955d6ab04-secp256k1}"
//time="2024-11-20T11:40:48+08:00" level=debug msg="party {0,4d9389cf937b1e6f2258d842e7237e955d6ab04-secp256k1} waiting for []"
//time="2024-11-20T11:40:48+08:00" level=debug msg="Starting party"
//scheduler_test.go:117: index: 3 submitter: 0x5091FC3cb4E4BB014141Aa41375d8Dd73b34AfA2
//time="2024-11-20T11:40:48+08:00" level=info msg="p2p loop started"
//time="2024-11-20T11:40:48+08:00" level=info msg="proposal loop started"
//time="2024-11-20T11:40:48+08:00" level=info msg="TSS keygen process started Candidate:0x5091FC3cb4E4BB014141Aa41375d8Dd73b34AfA2proposer: 0x04d9389Cf937b1e6F2258d842e7237E955d6ab04"
//scheduler_test.go:117: index: 4 submitter: 0xf6D37CE75dB465DcDb4c7097bEB9c1D46b171037
//time="2024-11-20T11:40:48+08:00" level=info msg="p2p loop started"
