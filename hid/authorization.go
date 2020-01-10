package hid

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// AuthorizeRequest takes an incoming request on behalf of the service and extracts the token from the "Authorization" header.
// The token is then checked for authenticity, and then the claims of thet token is verified against the provided scope and audince.
func (hid *HID) AuthorizeRequest(r *http.Request, audience, scope string) error {
	rawToken := r.Header.Get("Authorization")

	token, err := hid.authenticate(rawToken)
	if err != nil {
		return errors.Wrap(err, "while authenticating")
	}

	err = verifyClaims(token, hid.Addr, audience, scope)
	if err != nil {
		return errors.Wrap(err, "while verifying claims")
	}

	return nil
}

// authenticate verifies the authenticity of a provided raw token
func (hid *HID) authenticate(rawToken string) (*jwt.Token, error) {
	provideKeys(hid.PKS)
	defer revokeKeys()

	token, err := parseToken(rawToken)
	if err != nil {

		// might fail if stored public keys are outdated. Renew keys and retry once
		err = hid.newPKS()
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
	token, err := jwt.ParseWithClaims(rawToken, &claims{}, keyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}
