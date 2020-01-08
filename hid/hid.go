package hid

import (
	"net/http"
	"os"

	"github.com/hafslundnett/hn-config-lib-go/service"
)

// IdentityManager expl
type IdentityManager interface {
	AuthorizeRequest(r *http.Request, audience, scope string) error
	GetToken(user, secret string) (token *Token, err error)
}

// HID expl
type HID struct {
	Config
	PKS PublicKeySet
}

// New expl
func New() (*HID, error) {
	hid := new(HID)
	cert := os.Getenv(envars["cert"])
	err := service.Setup(hid, cert)
	return hid, err
}

// ConnectToServer expl
func (hid *HID) ConnectToServer() error {
	return hid.NewPKS()
}
