package helpers

import (
	"testing"
)

func TestNavForSession(t *testing.T) {
	nav := NavForSession(true)
	if nav != mainNav {
		t.Error()
	}

	nav = NavForSession(false)
	if nav != setupNav {
		t.Error()
	}
}
