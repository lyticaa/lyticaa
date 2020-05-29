package types

type ProfitLossTable struct {
	Item   string  `json:"item"`
	Amount float64 `json:"amount"`
}

type ProfitLoss struct {
	Data            []ProfitLossTable `json:"data"`
	Draw            int64             `json:"draw"`
	RecordsTotal    int64             `json:"recordsTotal"`
	RecordsFiltered int64             `json:"recordsFiltered"`
}
