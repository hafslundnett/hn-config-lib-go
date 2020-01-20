package querier

import (
	"context"
	b64 "encoding/base64"

	"cloud.google.com/go/bigquery"
	"github.com/hafslundnett/hn-config-lib-go/vault"
	"golang.org/x/oauth2/google"
)

// Querier expl
type Querier interface {
	DoQuery(ctx context.Context, query string, dst ...interface{}) error
}

func getCreds(ctx context.Context, credPath string) (*google.Credentials, error) {
	v, err := vault.New()
	if err != nil {
		return nil, err
	}

	s, err := v.GetSecret(credPath)
	if err != nil {
		return nil, err
	}

	credstring, err := b64.StdEncoding.DecodeString(s.Data["service-account-key"])
	if err != nil {
		return nil, err
	}

	creds, err := google.CredentialsFromJSON(ctx, []byte(credstring), bigquery.Scope)
	if err != nil {
		return nil, err
	}

	return creds, nil
}
