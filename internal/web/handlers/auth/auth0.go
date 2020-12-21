package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
	"github.com/lyticaa/lyticaa-app/internal/web/pkg/auth/iam"
	"github.com/lyticaa/lyticaa-app/internal/web/pkg/users"

	"github.com/coreos/go-oidc"
)

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	state := base64.StdEncoding.EncodeToString(b)

	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	authenticator, err := iam.NewAuth0()
	if err != nil {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

func (a *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logoutUrl, err := url.Parse(os.Getenv("AUTH0_URL"))
	if err != nil {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logoutUrl.Path += "v2/logout"
	parameters := url.Values{}

	returnTo, err := url.Parse(os.Getenv("BASE_URL"))
	if err != nil {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)
}

func (a *Auth) Callback(w http.ResponseWriter, r *http.Request) {
	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if r.URL.Query().Get("state") != session.Values["state"] {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authenticator, err := iam.NewAuth0()
	if err != nil {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := authenticator.Config.Exchange(context.TODO(), r.URL.Query().Get("code"))
	if err != nil {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	oidcConfig := &oidc.Config{ClientID: os.Getenv("AUTH0_CLIENT_ID")}
	idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(context.TODO(), rawIDToken)
	if err != nil {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	session.Values["id_token"] = rawIDToken
	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile

	parts := strings.Split(profile["sub"].(string), "|")
	userID := parts[1]

	user := users.FetchUser(r.Context(), userID, a.db)
	if user.ID == 0 {
		if err := users.CreateUser(r.Context(), userID, profile["name"].(string), profile["nickname"].(string), profile["picture"].(string), a.db); err != nil {
			a.logger.Err(err).Msg(helpers.ErrorTag)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user = users.FetchUser(r.Context(), userID, a.db)
	}

	session.Values["User"] = user
	err = session.Save(r, w)
	if err != nil {
		a.logger.Err(err).Msg(helpers.ErrorTag)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, helpers.DashboardRoute(), http.StatusSeeOther)
}
