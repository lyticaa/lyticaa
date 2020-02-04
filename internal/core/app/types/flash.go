package types

type Flash struct {
	Success string `json:"success"`
	Error   string `json:"error"`
	Info    string `json:"info"`
}
