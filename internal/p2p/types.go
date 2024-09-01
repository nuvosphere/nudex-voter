package p2p

type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type HeartbeatMessage struct {
	PeerID    string `json:"peer_id"`
	Message   string `json:"message"`
	Timestamp int64 `json:"ts"`
}
