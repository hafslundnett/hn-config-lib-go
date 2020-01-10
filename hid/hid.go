package hid

import (
	"net/http"
	"os"

	"github.com/hafslundnett/hn-config-lib-go/service"
)

// IDManager represents a service that is able to provide clients with authoriation tokens with the GetToken function,
// and is capable of authorizing these incoming tokens for the server with the AuthorizeRequest function.
type IDManager interface {
	GetToken(user, secret string) (token *Token, err error)
	AuthorizeRequest(r *http.Request, audience, scope string) error
}

// HID holds the configurations and keys neccessary to communicate with the HID service.
type HID struct {
	Config
	PKS
}

// New creates a new HID, performs neccessary setup, and returns it as an IDManager
func New() (IDManager, error) {
	hid := new(HID)
	cert := os.Getenv(envars["cert"])
	err := service.Setup(hid, cert)
	return hid, err
}

// ConnectToServer performs neccessary setup for connections to the external HID service
func (hid *HID) ConnectToServer() error {
	return hid.newPKS()
}
