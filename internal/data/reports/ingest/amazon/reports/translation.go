package reports

var dict = map[string]string{
	"date/time": "date/time",
}

func (r *Reports) translateHeader(header string) string {
	if _, ok := dict[header]; ok {
		return dict[header]
	}

	return header
}
