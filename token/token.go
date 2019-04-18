package token

import (
	"context"
	"strings"

	"log"

	"github.com/docker/docker/api/types"
)

const (
	ecrURL = "amazonaws.com"
	gcrURL = "grc.io"
)

type Registry interface {
	GetCredential(ctx context.Context) (string, string, error)
}

func GetToken(ctx context.Context, auth types.AuthConfig) types.AuthConfig {
	if auth.Username != "" || auth.Password != "" {
		return auth
	}
	var registry Registry
	switch {
	case strings.HasSuffix(auth.ServerAddress, ecrURL):
		registry = NewECR()
	case strings.HasSuffix(auth.ServerAddress, gcrURL):
		registry = NewECR()
	}
	var err error
	auth.Username, auth.Password, err = registry.GetCredential(ctx)
	if err != nil {
		log.Printf("failed to get token: %s", err)
	}
	return auth
}
