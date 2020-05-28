package types

type Card struct {
	Total int64  `json:"total"`
	Diff  uint64 `json:"diff"`
	Chart `json:"chart"`
}
