package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gitlab.com/getlytica/lytica-app/internal/core/iam"
	"gitlab.com/getlytica/lytica-app/internal/models"

	"github.com/coreos/go-oidc"
)

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		a.logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	state := base64.StdEncoding.EncodeToString(b)

	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		a.logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		a.logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	authenticator, err := iam.NewIAM()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

func (a *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		a.logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logoutUrl, err := url.Parse(os.Getenv("AUTH0_URL"))
	if err != nil {
		a.logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logoutUrl.Path += "v2/logout"
	parameters := url.Values{}

	returnTo, err := url.Parse("https://" + r.Host)
	if err != nil {
		a.logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	session.Options.MaxAge = -1
	if err = session.Save(r, w); err != nil {
		a.logger.Error().Err(err).Msg("error removing session")
	}

	http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)
}

func (a *Auth) Callback(w http.ResponseWriter, r *http.Request) {
	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		a.logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.URL.Query().Get("state") != session.Values["state"] {
		a.logger.Error().Err(err)
		http.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	authenticator, err := iam.NewIAM()
	if err != nil {
		a.logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := authenticator.Config.Exchange(context.TODO(), r.URL.Query().Get("code"))
	if err != nil {
		a.logger.Error().Err(err).Msg("Unauthorized.")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}

	oidcConfig := &oidc.Config{
		ClientID: os.Getenv("AUTH0_CLIENT_ID"),
	}

	idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(context.TODO(), rawIDToken)
	if err != nil {
		a.logger.Error().Err(err)
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["id_token"] = rawIDToken
	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile

	parts := strings.Split(profile["sub"].(string), "|")
	userId := parts[1]

	user, err := models.CreateUser(userId, profile["name"].(string), a.db)
	if err != nil {
		a.logger.Error().Err(err).Msg("unable to create user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Nickname = profile["nickname"].(string)
	user.Picture = profile["picture"].(string)
	session.Values["User"] = user

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
