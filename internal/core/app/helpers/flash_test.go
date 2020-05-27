package helpers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
)

const (
	successMsg = "Success"
	errorMsg   = "Error"
	warningMsg = "Warning"
	infoMsg    = "Info"
)

func SetupFlashTests(t *testing.T) *httptest.Server {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		store := testSessionStore(t)
		session, err := store.Get(r, "auth-session")
		if err != nil {
			t.Error()
		}

		switch r.RequestURI {
		case "/success":
			err := SetFlash("success", successMsg, session, r, w)
			if err != nil {
				t.Error(err)
			}

			flash := session.Values["Flash"].(types.Flash)
			if flash.Success != successMsg {
				t.Error()
			}

			ClearFlash(session, r, w)
			flash = session.Values["Flash"].(types.Flash)
			if flash.Success != "" {
				t.Error()
			}
		case "/error":
			err := SetFlash("error", errorMsg, session, r, w)
			if err != nil {
				t.Error(err)
			}

			flash := session.Values["Flash"].(types.Flash)
			if flash.Error != errorMsg {
				t.Error()
			}

			ClearFlash(session, r, w)
			flash = session.Values["Flash"].(types.Flash)
			if flash.Error != "" {
				t.Error()
			}
		case "/warning":
			err := SetFlash("warning", warningMsg, session, r, w)
			if err != nil {
				t.Error(err)
			}

			flash := session.Values["Flash"].(types.Flash)
			if flash.Warning != warningMsg {
				t.Error()
			}

			ClearFlash(session, r, w)
			flash = session.Values["Flash"].(types.Flash)
			if flash.Warning != "" {
				t.Error()
			}
		case "/info":
			err := SetFlash("info", infoMsg, session, r, w)
			if err != nil {
				t.Error(err)
			}

			flash := session.Values["Flash"].(types.Flash)
			if flash.Info != infoMsg {
				t.Error()
			}

			ClearFlash(session, r, w)
			flash = session.Values["Flash"].(types.Flash)
			if flash.Info != "" {
				t.Error()
			}
		}
	})
	s := httptest.NewServer(h)

	return s
}

func TestSetFlashSuccess(t *testing.T) {
	s := SetupFlashTests(t)

	r, err := http.NewRequest("GET", s.URL+"/success", nil)
	if err != nil {
		t.Error()
	}

	client := http.Client{}
	_, err = client.Do(r)
	if err != nil {
		t.Error()
	}
}

func TestSetFlashError(t *testing.T) {
	s := SetupFlashTests(t)

	r, err := http.NewRequest("GET", s.URL+"/error", nil)
	if err != nil {
		t.Error()
	}

	client := http.Client{}
	_, err = client.Do(r)
	if err != nil {
		t.Error()
	}
}

func TestSetFlashWarning(t *testing.T) {
	s := SetupFlashTests(t)

	r, err := http.NewRequest("GET", s.URL+"/warning", nil)
	if err != nil {
		t.Error()
	}

	client := http.Client{}
	_, err = client.Do(r)
	if err != nil {
		t.Error()
	}
}

func TestSetFlashInto(t *testing.T) {
	s := SetupFlashTests(t)

	r, err := http.NewRequest("GET", s.URL+"/info", nil)
	if err != nil {
		t.Error()
	}

	client := http.Client{}
	_, err = client.Do(r)
	if err != nil {
		t.Error()
	}
}
