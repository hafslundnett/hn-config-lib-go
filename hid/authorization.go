package hid

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// AuthorizeRequest token
func (hid *HID) AuthorizeRequest(r *http.Request, audience, scope string) error {
	rawToken := r.Header.Get("Authorization")

	token, err := hid.Authenticate(rawToken)
	if err != nil {
		return err
	}

	err = VerifyClaims(token, hid.Addr, audience, scope)
	if err != nil {
		return err
	}

	return nil
}

// Authenticate expl
func (hid *HID) Authenticate(rawToken string) (*jwt.Token, error) {
	provideKeys(hid.PKS.Keys)
	defer provideKeys(nil)

	token, err := parseToken(rawToken)
	if err != nil {

		// might fail if stored public keys are outdated. Renew keys and retry once
		err = hid.NewPKS()
		if err != nil {
			return nil, err
		}

		token, err = parseToken(rawToken)
		if err != nil {
			return nil, err
		}
	}

	if !token.Valid {
		return token, errors.New("Invalid token")
	}

	return token, nil
}

func parseToken(rawToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(rawToken, &Claims{}, keyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}
