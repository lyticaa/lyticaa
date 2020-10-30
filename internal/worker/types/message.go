package types

type Message struct {
	UserID int64  `json:"userID"`
	Body   string `json:"body"`
	Op     string `json:"op"`
}

type Payload struct {
	Message `json:"message"`
}
