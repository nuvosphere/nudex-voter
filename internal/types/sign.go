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
	Task Task
}

type TssMessage struct {
	FromPartyId  string   `json:"from_party_id"`
	ToPartyIds   []string `json:"to_party_ids"`
	IsBroadcast  bool     `json:"is_broadcast"`
	MsgWireBytes []byte   `json:"msg_wire_bytes"`
}
