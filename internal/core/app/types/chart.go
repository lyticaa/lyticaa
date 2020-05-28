package types

type Sparkline struct {
	Plots []float64 `json:"plots"`
}

type Flot struct {
	Plots [][]float64 `json:"plots"`
}

type Chart struct {
	Sparkline `json:"sparkline,omitempty"`
	Flot      `json:"flot,omitempty"`
}
