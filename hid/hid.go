package hid

import (
	"github.com/hafslundnett/hn-config-lib-go/libhttp"
)

// HID expl
type HID struct {
	Config
	Client *libhttp.Client
	PKS    PublicKeySet
}

// New expl
func New() (*HID, error) {
	hid := new(HID)

	if err := hid.NewClient(); err != nil {
		return hid, err
	}

	if err := hid.NewConfig(); err != nil {
		return hid, err
	}

	if err := hid.NewPKS(); err != nil {
		return hid, err
	}

	return hid, nil
}


// MakeClient returns a http client configured according to the supplied Config, for use with HID
func (hid *HID) MakeClient() (err error) {
	vault.Client, err = libhttp.NewClient(hid.Config.PemCert)
	return
}
