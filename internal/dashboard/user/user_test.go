package user

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/rs/zerolog/log"
)

var (
	userId = "5de89aea5a61280de1f1bf2b"
	email  = "test@getlytica.com"
)

func TestResetPassword(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`an email to reset your password has been sent to you`))
		if err != nil {
			t.Error(err)
		}
	})

	s := httptest.NewServer(h)
	defer s.Close()

	_ = os.Setenv("AUTH0_PASSWORD_RESET_URL", s.URL+"/dbconnections/change_password")

	u := NewUser(userId, email, log.Logger)
	err := u.ResetPassword()

	if err != nil {
		t.Error(err)
	}
}
