package hid

import "github.com/pkg/errors"

// Token exp
type Token struct {
	Raw   string `json:"access_token"`
	Scope string `json:"scope"`
	Type  string `json:"token_type"`
	Exp   int    `json:"expires_in"`
}

// GetToken exp
func (hid HID) GetToken(user, secret string) (token *Token, err error) {
	form := map[string][]string{
		"client_id":     []string{user},
		"client_secret": []string{secret},
		"grant_type":    []string{"client_credentials"},
	}

	err = hid.Client.PostForm(hid.TokenEP, form, &token)
	err = errors.Wrap(err, "while getting token from HID")

	return
}
