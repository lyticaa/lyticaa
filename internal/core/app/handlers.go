package app

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gitlab.com/getlytica/lytica/internal/core/app/types"
	"gitlab.com/getlytica/lytica/internal/core/auth0"
	"gitlab.com/getlytica/lytica/internal/models"

	"github.com/coreos/go-oidc"
	"github.com/gorilla/sessions"
)

func (a *App) healthCheck(w http.ResponseWriter, r *http.Request) {
	response := types.Health{Status: "OK"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		a.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonResponse)
	if err != nil {
		a.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) home(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"partials/nav/_main", "home", "partials/_filters"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) login(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		a.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	state := base64.StdEncoding.EncodeToString(b)

	session, err := a.SessionStore.Get(r, "auth-session")
	if err != nil {
		a.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		a.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	authenticator, err := auth0.NewAuthenticator()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

func (a *App) logout(w http.ResponseWriter, r *http.Request) {
	session, err := a.SessionStore.Get(r, "auth-session")
	if err != nil {
		a.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logoutUrl, err := url.Parse(os.Getenv("AUTH0_URL"))
	if err != nil {
		a.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logoutUrl.Path += "v2/logout"
	parameters := url.Values{}

	returnTo, err := url.Parse("https://" + r.Host)
	if err != nil {
		a.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	session.Options.MaxAge = -1
	if err = session.Save(r, w); err != nil {
		a.Logger.Error().Err(err).Msg("error removing session")
	}

	http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)
}

func (a *App) callback(w http.ResponseWriter, r *http.Request) {
	session, err := a.SessionStore.Get(r, "auth-session")
	if err != nil {
		a.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.URL.Query().Get("state") != session.Values["state"] {
		a.Logger.Error().Err(err)
		http.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	authenticator, err := auth0.NewAuthenticator()
	if err != nil {
		a.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := authenticator.Config.Exchange(context.TODO(), r.URL.Query().Get("code"))
	if err != nil {
		a.Logger.Error().Err(err)
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
		a.Logger.Error().Err(err)
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

	user, err := models.CreateUser(userId, profile["name"].(string), a.Db)
	if err != nil {
		a.Logger.Error().Err(err).Msg("unable to create user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["User"] = user
	session.Values["nickname"] = profile["nickname"].(string)
	session.Values["email"] = profile["name"].(string)
	session.Values["picture"] = profile["picture"].(string)

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *App) getSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := a.SessionStore.Get(r, "auth-session")
	if err != nil {
		a.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return a.resetFlash(session)
}

func (a *App) resetFlash(session *sessions.Session) *sessions.Session {
	session.Values["Flash"] = nil
	return session
}
