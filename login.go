package vault

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

type gitToken struct {
	Token string `json:"token"`
}

type k8Token struct {
	JWT  string `json:"jwt"`
	Role string `json:"role"`
}

func githubLogin(login string) (*bytes.Buffer, error) {
	return loginBuffer(&gitToken{
		Token: login,
	})
}

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

func loginBuffer(lt interface{}) (*bytes.Buffer, error) {
	js, err := json.Marshal(lt)
	if err != nil {
		return nil, errors.Wrap(err, "while marshaling token")
	}

	return bytes.NewBuffer(js), nil
}

func getJWT(k8ServicePath string) (string, error) {
	b, err := ioutil.ReadFile(k8ServicePath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read jwt token from %s", k8ServicePath)
	}

	return string(bytes.TrimSpace(b)), nil
}
