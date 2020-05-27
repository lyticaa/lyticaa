package helpers

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

const (
	dtDraw   = 1
	dtStart  = 0
	dtLength = 10
	dtSort   = 0
	dtDir    = "asc"
)

var (
	dataTableUrl = fmt.Sprintf("/?draw=%v&start=%v&length=%v&order[0][column]=%v&order[0][dir]=%v", dtDraw, dtStart, dtLength, dtSort, dtDir)
)

func SetupDataTableTests(t *testing.T) *http.Request {
	r, err := http.NewRequest(http.MethodGet, dataTableUrl, nil)
	if err != nil {
		t.Error(err)
	}

	return r
}

func TestDtDraw(t *testing.T) {
	r := SetupDataTableTests(t)
	draw := DtDraw(r)

	if draw != dtDraw {
		t.Error()
	}
}

func TestDtStart(t *testing.T) {
	r := SetupDataTableTests(t)
	start := DtStart(r)

	if start != dtStart {
		t.Error()
	}
}

func TestDtLength(t *testing.T) {
	r := SetupDataTableTests(t)
	length := DtLength(r)

	if length != dtLength {
		t.Error()
	}
}

func TestDtSort(t *testing.T) {
	r := SetupDataTableTests(t)
	sort := DtSort(r)

	if sort != dtSort {
		t.Error()
	}
}

func TestDtDir(t *testing.T) {
	r := SetupDataTableTests(t)
	dir := DtDir(r)

	if dir != strings.ToUpper(dtDir) {
		t.Error()
	}
}
