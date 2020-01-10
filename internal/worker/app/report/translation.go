package report

var dict = map[string]string{
	"date/time": "date/time",
}

func (r *Report) translateHeader(header string) string {
	if _, ok := dict[header]; ok {
		return dict[header]
	}

	return header
}
