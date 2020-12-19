package types

type StatementTable struct {
	Item   string  `json:"item"`
	Amount float64 `json:"amount"`
}

type Statement struct {
	Data            []StatementTable `json:"data"`
	Draw            int64            `json:"draw"`
	RecordsTotal    int64            `json:"recordsTotal"`
	RecordsFiltered int64            `json:"recordsFiltered"`
}
