package hid

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// Claims expl
type Claims struct {
	ClientID string   `json:"client_id,omitempty"`
	Scope    []string `json:"scope,omitempty"`
	jwt.StandardClaims
}

// VerifyClaims expl
func VerifyClaims(token *jwt.Token, issuer, audience, scope string) error {
	if !token.Claims.(*Claims).VerifyAudience(audience, false) {
		return errors.New("Invalid audience")
	}

	if !token.Claims.(*Claims).VerifyIssuer(issuer, false) {
		return errors.New("Invalid issuer")
	}

	return verifyScope(token, scope)
}

func verifyScope(token *jwt.Token, scope string) error {
	if claims, ok := token.Claims.(*Claims); ok {
		for _, s := range claims.Scope {
			if s == scope {
				return nil // scope found, verified scope
			}
		}
		return errors.New("Did not find valid scope")
	}
	return errors.New("Invalid token")
}
