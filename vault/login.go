package vault

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

// gitToken holds github authentication information to be formatted to a bytes buffer
type gitToken struct {
	Token string `json:"token"`
}

// k8Token holds kubernetes authentication information to be formatted to a bytes buffer
type k8Token struct {
	JWT  string `json:"jwt"`
	Role string `json:"role"`
}

// githubLogin handles converting the github token string to a bytes buffer
func githubLogin(login string) (*bytes.Buffer, error) {
	return loginBuffer(&gitToken{
		Token: login,
	})
}

// k8Login handles converting the service path and role to a bytes buffer
func k8Login(k8ServicePath string, role string) (*bytes.Buffer, error) {
	jwt, err := getJWT(k8ServicePath)
	if err != nil {
		return nil, err
	}

	return loginBuffer(&k8Token{
		JWT:  jwt,
		Role: role,
	})
}

// loginBuffer converts a login token to a bytes buffer
func loginBuffer(lt interface{}) (*bytes.Buffer, error) {
	js, err := json.Marshal(lt)
	if err != nil {
		return nil, errors.Wrap(err, "while marshaling token")
	}

	return bytes.NewBuffer(js), nil
}

// getJWT reads JSON web token from file at the service path
func getJWT(k8ServicePath string) (string, error) {
	b, err := ioutil.ReadFile(k8ServicePath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read jwt token from %s", k8ServicePath)
	}

	return string(bytes.TrimSpace(b)), nil
}
