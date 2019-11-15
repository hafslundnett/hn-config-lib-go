package hid

import (
	"hafslundnett/x/hn-config-lib/hnhttp"

	"github.com/pkg/errors"
)

// Token exp
type Token struct {
	Access     string `json:"access_token"`
	Expiration int    `json:"expires_in"`
	Type       string `json:"token_type"`
	Scope      string `json:"scope"`
}

// GetToken exp
func (hid HIDclient) GetToken(userID string) (*Token, error) {
	client, err := hnhttp.NewClient()
	if err != nil {
		return nil, errors.Wrap(err, "while setting up http client")
	}

	target := hid.Host + "/" + hid.Path

	values := map[string][]string{
		"client_id":     []string{userID},
		"client_secret": []string{hid.Secret},
		"grant_type":    []string{"client_credentials"},
	}

	var token Token
	err = client.PostForm(target, values, &token)
	if err != nil {
		return nil, errors.Wrap(err, "while getting token from HID")
	}

	return &token, nil
}
