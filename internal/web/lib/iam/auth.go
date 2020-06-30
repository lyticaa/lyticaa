package iam

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type IAM struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

func NewIAM() (*IAM, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, os.Getenv("AUTH0_URL"))
	if err != nil {
		log.Printf("failed to get provider: %v", err)
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RedirectURL:  fmt.Sprintf("%v%v", os.Getenv("BASE_URL"), os.Getenv("AUTH0_CALLBACK_URI")),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &IAM{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil
}
