package types

var Allowed = []string{
	"text/csv",
}

func ValidMime(mime string) bool {
	for _, item := range Allowed {
		if item == mime {
			return true
		}
	}

	return false
}
