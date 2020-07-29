package types

type Message struct {
	UserId int64  `json:"userId"`
	Body   string `json:"body"`
	Op     string `json:"op"`
}

type Payload struct {
	Message `json:"message"`
}
