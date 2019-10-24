package vault

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

type gitToken struct {
	Token string `json:"token"`
}

type k8Token struct {
	JWT  string `json:"jwt"`
	Role string `json:"role"`
}

func gitLogin(login string) (*bytes.Buffer, error) {
	return loginBuffer(&gitToken{
		Token: login,
	})
}

func k8Login(jwt string, role string) (*bytes.Buffer, error) {
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
