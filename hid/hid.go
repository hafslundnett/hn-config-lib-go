package hid

import (
	"hafslundnett/x/hn-config-lib/hnhttp"
)

// HID expl
type HID struct {
	Config
	Client *hnhttp.Client
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

// NewClient returns a http client for use with HID
func (hid *HID) NewClient() error {
	var err error
	hid.Client, err = hnhttp.NewClient()
	return err
}
