package core

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gitlab.com/getlytica/dashboard/internal/auth"
	"gitlab.com/getlytica/dashboard/internal/core/types"
	"gitlab.com/getlytica/dashboard/internal/models"

	"github.com/coreos/go-oidc"
)

func (c *Core) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := types.Health{Status: "OK"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		c.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		c.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Core) Home(w http.ResponseWriter, r *http.Request) {
	c.RenderTemplate(w, "home", nil)
}

func (c *Core) Login(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		c.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	state := base64.StdEncoding.EncodeToString(b)

	session, err := c.SessionStore.Get(r, "auth-session")
	if err != nil {
		c.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		c.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

func (c *Core) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := c.SessionStore.Get(r, "auth-session")
	if err != nil {
		c.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logoutUrl, err := url.Parse(os.Getenv("AUTH0_URL"))
	if err != nil {
		c.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logoutUrl.Path += "v2/logout"
	parameters := url.Values{}

	returnTo, err := url.Parse("https://" + r.Host)
	if err != nil {
		c.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	session.Options.MaxAge = -1
	if err = session.Save(r, w); err != nil {
		c.Logger.Error().Err(err).Msg("error removing session")
	}

	http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)
}

func (c *Core) Callback(w http.ResponseWriter, r *http.Request) {
	session, err := c.SessionStore.Get(r, "auth-session")
	if err != nil {
		c.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.URL.Query().Get("state") != session.Values["state"] {
		c.Logger.Error().Err(err)
		http.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		c.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := authenticator.Config.Exchange(context.TODO(), r.URL.Query().Get("code"))
	if err != nil {
		c.Logger.Error().Err(err)
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
		c.Logger.Error().Err(err)
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
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	parts := strings.Split(profile["sub"].(string), "|")
	userId := parts[1]
	models.CreateUser(userId, profile["name"].(string), c.Db)

	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

func (c *Core) User(w http.ResponseWriter, r *http.Request) {
	session, err := c.SessionStore.Get(r, "auth-session")
	if err != nil {
		c.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.RenderTemplate(w, "user", session.Values["profile"])
}

func (c *Core) AccountSubscribe(w http.ResponseWriter, r *http.Request) {

}
