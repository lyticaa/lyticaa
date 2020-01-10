package report

import (
	"testing"
)

func TestTranslateHeader(t *testing.T) {
	r, _, complete := SetupTests(t)
	defer complete(r)

	expected := "date/time"
	actual := r.translateHeader(expected)

	if actual != expected {
		t.Errorf("header does not match")
	}
}
