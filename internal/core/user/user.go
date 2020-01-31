package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

type User struct {
	Logger zerolog.Logger
	UserId string
	Email  string
}

func NewUser(userId, email string, log zerolog.Logger) *User {
	return &User{
		Logger: log,
		UserId: userId,
		Email:  email,
	}
}

func (u *User) ResetPassword() error {
	req, err := json.Marshal(map[string]string{
		"client_id":  os.Getenv("AUTH0_CLIENT_ID"),
		"email":      u.Email,
		"connection": "Username-Password-Authentication",
	})
	if err != nil {
		u.Logger.Error().Err(err)
		return err
	}

	res, err := http.Post(os.Getenv("AUTH0_PASSWORD_RESET_URL"),
		"application/json", bytes.NewBuffer(req))
	if err != nil {
		u.Logger.Error().Err(err)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = fmt.Errorf("unable to request a password reset for %v", u.Email)
		u.Logger.Error().Err(err)
		return err
	}

	return nil
}
