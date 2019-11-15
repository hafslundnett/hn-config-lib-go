package hid

import (
	"hafslundnett/x/hn-config-lib/testing/mock"
)

func mockHID() HIDclient {
	return HIDclient{
		Host:   mock.URL,
		Path:   mock.Path,
		Secret: mock.Secret,
	}
}
