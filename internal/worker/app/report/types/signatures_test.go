package types

import (
	"testing"
)

func TestShouldIgnore(t *testing.T) {
	for _, row := range Ignore {
		ok := ShouldIgnore(row)
		if !ok {
			t.Errorf("%v should be ignored", row)
		}
	}
}
