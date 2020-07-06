package iam

import (
	"context"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type IAM struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

