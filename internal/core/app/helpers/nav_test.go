package helpers

import (
	"testing"
)

func TestNavForSession(t *testing.T) {
	nav := PrimaryNavForSession(true)
	if nav != mainPrimaryNav {
		t.Error()
	}

	nav = PrimaryNavForSession(false)
	if nav != setupPrimaryNav {
		t.Error()
	}

	nav = AccountNavForSession(true)
	if nav != mainAccountNav {
		t.Error()
	}

	nav = AccountNavForSession(false)
	if nav != setupAccountNav {
		t.Error()
	}
}
