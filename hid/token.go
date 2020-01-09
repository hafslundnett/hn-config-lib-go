package hid

import (
	"net/http"

	"github.com/pkg/errors"
)

// Token exp
type Token struct {
	Raw   string `json:"access_token"`
	Scope string `json:"scope"`
	Type  string `json:"token_type"`
	Exp   int    `json:"expires_in"`
}

// GetToken provides the credentials of a user or service, and returns a token for sending with requests to a service.
func (hid HID) GetToken(user, secret string) (token *Token, err error) {
	form := map[string][]string{
		"client_id":     []string{user},
		"client_secret": []string{secret},
		"grant_type":    []string{"client_credentials"},
	}

	err = hid.client.PostForm(hid.TokenEP, form, &token)
	err = errors.Wrap(err, "while getting token from HID")

	return
}

// AppendTokenToRequest is a simple quality of life function that appends the raw token to the header of the provided request.
func (hid HID) AppendTokenToRequest(req *http.Request, token Token) {
	req.Header.Add("Authorization", token.Raw)
}
