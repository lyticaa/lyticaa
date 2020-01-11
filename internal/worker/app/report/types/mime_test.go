package types

import (
	"testing"
)

var (
	valid = []string{
		"text/csv",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	}
	invalid = []string{
		"application/vnd.ms-excel",
		"application/x-msexcel",
		"application/x-ms-excel",
		"application/x-excel",
		"application/x-dos_ms_excel",
		"application/xls",
		"application/x-xls",
	}
)

func TestValidMime(t *testing.T) {
	for _, mimeType := range valid {
		ok := ValidMime(mimeType)
		if !ok {
			t.Error("mime type invalid")
		}
	}

	for _, mimeType := range invalid {
		ok := ValidMime(mimeType)
		if ok {
			t.Errorf("mime type %v should be invalid but is being returned as valid", mimeType)
		}
	}
}

func TestIsCsv(t *testing.T) {
	ok := IsCsv(valid[0])
	if !ok {
		t.Error("expected CSV")
	}
}

func TestIsXlsx(t *testing.T) {
	ok := IsXlsx(valid[1])
	if !ok {
		t.Error("expected XLSX")
	}
}
