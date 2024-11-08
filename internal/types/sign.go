package types

type BaseSignMsg struct {
	RequestId    string `json:"request_id"`
	VoterAddress string `json:"voter_address"`
	Nonce        uint64 `json:"nonce"`
	IsProposer   bool   `json:"is_proposer"`
	CreateTime   int64  `json:"create_time"`
}

type SignMessage struct {
	BaseSignMsg
	Task SignTask
}

type SignTask struct {
	TaskId uint32 `json:"task_id"`
	Data   []byte `json:"data"`
}

type TssMessage struct {
	FromPartyId  string   `json:"from_party_id"`
	ToPartyIds   []string `json:"to_party_ids"`
	IsBroadcast  bool     `json:"is_broadcast"`
	MsgWireBytes []byte   `json:"msg_wire_bytes"`
}
