package types

type MsgSign struct {
	RequestId    string `json:"request_id"`
	VoterAddress string `json:"voter_address"`
	IsProposer   bool   `json:"is_proposer"`
	CreateTime   int64  `json:"create_time"`
}

type MsgSignCreateWalletMessage struct {
	MsgSign
	Task CreateWalletTask `json:"task"`
}

type TssUpdateMessage struct {
	FromPartyId  string   `json:"from_party_id"`
	ToPartyIds   []string `json:"to_party_ids"`
	IsBroadcast  bool     `json:"is_broadcast"`
	MsgWireBytes []byte   `json:"msg_wire_bytes"`
}
