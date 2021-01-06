package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/lyticaa/lyticaa/internal/app/helpers"
)

func (a *Account) ChangePassword(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(a.sessionStore, a.logger, w, r))

	if err := a.changePasswordRequest(user.Email); err != nil {
		a.logger.Error().Err(err).Msg("password reset request failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *Account) changePasswordRequest(email string) error {
	req, err := json.Marshal(map[string]string{
		"client_id":  os.Getenv("AUTH0_CLIENT_ID"),
		"email":      email,
		"connection": "Username-Password-Authentication",
	})
	if err != nil {
		return err
	}

	res, err := http.Post(os.Getenv("AUTH0_PASSWORD_RESET_URL"),
		"application/json", bytes.NewBuffer(req))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = fmt.Errorf("unable to request a password reset for %v", email)
		return err
	}

	return nil
}
