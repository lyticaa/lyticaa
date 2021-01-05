package types

var (
	types   = make([]string, 2)
	csv     = "text/csv"
	xlsx    = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	allowed = append(types, csv, xlsx)
)

func ValidMime(contentType string) bool {
	for _, item := range allowed {
		if item == contentType {
			return true
		}
	}

	return false
}

func IsCSV(contentType string) bool {
	match := false
	if contentType == csv {
		match = true
	}

	return match
}

func IsXLSX(contentType string) bool {
	match := false
	if contentType == xlsx {
		match = true
	}

	return match
}
