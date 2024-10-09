package types

type MsgSign struct {
	RequestId    string `json:"request_id"`
	VoterAddress string `json:"voter_address"`
	IsProposer   bool   `json:"is_proposer"`
	SigData      []byte `json:"sig_data"`
	// SigPk        []byte `json:"sig_pk"`
	CreateTime int64 `json:"create_time"`
}

type MsgSignKeyPrepareMessage struct {
	MsgSign
	PublicKeys []string `json:"public_keys"`
	Threshold  int      `json:"threshold"`
}

type KeygenReqMessage struct {
	RequestId    string   `json:"request_id"`
	VoterAddress string   `json:"voter_address"`
	IsProposer   bool     `json:"is_proposer"`
	PublicKeys   []string `json:"public_keys"`
	Threshold    int      `json:"threshold"`
	CreateTime   int64    `json:"create_time"`
}
