package vault

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

type loginToken struct {
	Token string `json:"token"`
}

func loginBuffer(login string) (*bytes.Buffer, error) {
	lt := &loginToken{
		Token: login,
	}

	js, err := json.Marshal(lt)
	if err != nil {
		return nil, errors.Wrapf(err, "while marshaling token")
	}

	return bytes.NewBuffer(js), nil
}
