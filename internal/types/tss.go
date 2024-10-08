package types

type KeyGenPrepareMessage struct {
	PublicKeys  []string `json:"public_keys"`
	FromAddress string   `json:"from_address"`
	Threshold   int      `json:"threshold"`
	Timestamp   int64    `json:"ts"`
}

type KeygenMessage struct {
	Content string
}

type SigningMessage struct {
	Content string
}
