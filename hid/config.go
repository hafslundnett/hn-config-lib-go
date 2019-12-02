package hid

import (
	"os"
)

// Config expl
type Config struct {
	Addr    string
	JWKSuri string `json:"jwks_uri"`
	TokenEP string `json:"token_endpoint"`
}

// NewConfig expl
func (hid *HID) NewConfig() error {
	hid.Addr = os.Getenv("HID_ADDR")
	if hid.Addr == "" {
		hid.Addr = "https://127.0.0.1"
	}

	discovery := os.Getenv("HID_DISCOVERY")
	if discovery == "" {
		discovery = "/.well-known/openid-configuration"
	}

	err := hid.Client.Get(hid.Addr+discovery, &hid)
	if err != nil {
		hid.JWKSuri = hid.Addr + "/.well-known/openid-configuration/jwks"
		hid.TokenEP = hid.Addr + "/connect/token"
	}

	return nil
}
