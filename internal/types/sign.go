package types

type MsgSign struct {
	RequestId    string `json:"request_id"`
	VoterAddress string `json:"voter_address"`
	IsProposer   bool   `json:"is_proposer"`
	SigData      []byte `json:"sig_data"`
	// SigPk        []byte `json:"sig_pk"`
	CreateTime int64 `json:"create_time"`
}

type MsgSignCreateWalletMessage struct {
	MsgSign
	Task CreateWalletTask `json:"task"`
}

type MsgSignKeyPrepareMessage struct {
	MsgSign
	PublicKeys []string `json:"public_keys"`
	Threshold  int      `json:"threshold"`
}

type KeygenReqMessage struct {
	RequestId    string   `json:"request_id"`
	VoterAddress string   `json:"voter_address"`
	PublicKeys   []string `json:"public_keys"`
	Threshold    int      `json:"threshold"`
	CreateTime   int64    `json:"create_time"`
}

type KeygenReceiveMessage struct {
	RequestId         string   `json:"request_id"`
	VoterAddress      string   `json:"voter_address"`
	PublicKeys        []string `json:"public_keys"`
	Threshold         int      `json:"threshold"`
	PublicKeysMatched bool     `json:"public_keys_matched"`
	ThresholdMatched  bool     `json:"threshold_matched"`
	CreateTime        int64    `json:"create_time"`
}

type TssUpdateMessage struct {
	FromPartyId  string   `json:"from_party_id"`
	ToPartyIds   []string `json:"to_party_ids"`
	IsBroadcast  bool     `json:"is_broadcast"`
	MsgWireBytes []byte   `json:"msg_wire_bytes"`
}
