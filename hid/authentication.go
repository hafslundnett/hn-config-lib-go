package hid

import "hafslundnett/x/hn-config-lib/hnhttp"

// Login exp
type Login struct {
	URL    string
	ID     string
	Secret string
}

// Token exp
type Token struct {
	Access     string `json:"access_token"`
	Expiration int    `json:"expires_in"`
	Type       string `json:"token_type"`
	Scope      string `json:"scope"`
}

// GetToken exp
func (l Login) GetToken() (*Token, error) {
	client, err := hnhttp.NewClient()
	if err != nil {
		return nil, err
	}

	var token Token
	values := map[string][]string{
		"client_id":     []string{l.ID},
		"client_secret": []string{l.Secret},
		"grant_type":    []string{"client_credentials"},
	}

	err = client.PostForm(l.URL, values, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
