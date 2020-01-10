package hid

import (
	"os"

	"github.com/hafslundnett/hn-config-lib-go/libhttp"
	"github.com/pkg/errors"
)

const (
	defDiscovery = "/.well-known/openid-configuration"
	defJWKS      = "/.well-known/openid-configuration/jwks"
	defTokenEP   = "/connect/token"
)

var envars = map[string]string{
	"addr":      "HID_ADDR",
	"cert":      "HID_CACERT",
	"discovery": "HID_DISCOVERY",
}

// Config contains the configuration information needed to do the initial setup and renewal of the HID service
type Config struct {
	Addr    string
	JWKSuri string `json:"jwks_uri"`
	TokenEP string `json:"token_endpoint"`

	client libhttp.Client
}

// Configure expl
func (hid *HID) Configure(client libhttp.Client) error {
	addr := os.Getenv(envars["addr"])
	if addr == "" {
		return errors.New("missing env var " + envars["addr"])
	}

	discovery := os.Getenv(envars["discovery"])
	if discovery == "" {
		discovery = defDiscovery
	}

	err := client.Get(addr+discovery, &hid)
	if err != nil {
		hid.JWKSuri = addr + defJWKS
		hid.TokenEP = addr + defTokenEP
	}

	hid.Addr = addr
	hid.client = client

	return nil
}
