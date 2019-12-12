package hid

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// PublicKey expl
type PublicKey struct {
	KeyID     string   `json:"kid"`
	Algorithm string   `json:"alg"`
	X5C       []string `json:"x5c"`
}

// PublicKeySet expl
type PublicKeySet struct {
	Keys []PublicKey `json:"keys"`
}

// NewPKS expl
func (hid *HID) NewPKS() error {
	err := hid.Client.Get(hid.JWKSuri, &hid.PKS)
	if err != nil {
		return errors.Wrap(err, "while renewing HID public key set")
	}

	return nil
}

// keyFunc converts a jwt token to a RSA public key
func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	cert, err := getPemCert(token)
	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
}

func getPemCert(token *jwt.Token) (cert string, err error) {
	for _, k := range pks.Keys {
		if kid, ok := token.Header["kid"].(string); ok {
			if kid == k.KeyID {
				cert = "-----BEGIN CERTIFICATE-----\n" + k.X5C[0] + "\n-----END CERTIFICATE-----"
				return
			}
		} else {
			return "", errors.New("expecting JWT header to have string kid")
		}
	}

	return "", errors.New("Unable to find corresponding kid")
}

// needs to be globally accessible because of how dgrijalva/jwt-go works. Not for caching; set before each use, nil after use.
var pks PublicKeySet

func provideKeys(keys []PublicKey) {
	pks = PublicKeySet{
		Keys: keys,
	}
}
