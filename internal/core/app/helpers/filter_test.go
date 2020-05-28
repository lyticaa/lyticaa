package helpers

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestBuildFilter(t *testing.T) {
	url := fmt.Sprintf("%v", dataTableUrl)

	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Error(err)
	}

	filter := BuildFilter(r)

	if filter.Start != dtStart {
		t.Error()
	}

	if filter.Length != dtLength {
		t.Error()
	}

	if filter.Sort != dtSort {
		t.Error()
	}

	if filter.Dir != strings.ToUpper(dtDir) {
		t.Error()
	}
}
