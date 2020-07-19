package types

type AdminTable struct {
	RowId   string `json:"DT_RowId"`
	Email   string `json:"email"`
	Created string `json:"date"`
}

type Admin struct {
	Data            []AdminTable `json:"data"`
	Draw            int64        `json:"draw"`
	RecordsTotal    int64        `json:"recordsTotal"`
	RecordsFiltered int64        `json:"recordsFiltered"`
}
