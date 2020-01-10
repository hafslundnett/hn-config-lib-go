package lib

import (
	"log"
	"net/http"

	"github.com/hafslundnett/hn-config-lib-go/hid"
	"github.com/hafslundnett/hn-config-lib-go/vault"
)

// Example executes examples of the three core usecases of this package.
func Example() {
	mySecret := vaultExample()

	myRequest := hidClientExample(mySecret)

	hidAPIexample(myRequest)
}

// vaultExample represents the simplest way to get a secret from Vault.
// Requires, at a minimum, that env vars VAULT_ADDR and either GITHUB_TOKEN or the K8 related ones are set. See readme for more information.
func vaultExample() *vault.Secret {
	// Make reusable vault item
	myVault, err := vault.New()
	if err != nil {
		log.Fatal(err)
	}

	// Get a secret from the vault
	mySecret, err := myVault.GetSecret("path/to/secret")
	if err != nil {
		log.Fatal(err)
	}

	// Do something with the secret
	return mySecret
}

// hidClientExample represents the client side of a request with HID authorization. User and secret for HID.GetToken may be from a wide variety of sources.
// Requires, at a minimum, that env var HID_ADDR is set.
func hidClientExample(mySecret *vault.Secret) *http.Request {
	// Make reusable HID item
	myHID, err := hid.New()
	if err != nil {
		log.Fatal(err)
	}

	// Get a bearer token from HID
	myToken, err := myHID.GetToken("username", mySecret.Data["key"])
	if err != nil {
		log.Fatal(err)
	}

	// Make http.Request as usual
	myRequest, err := http.NewRequest("POST", "api.url", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add bearer token to http request header
	myToken.AppendToRequest(myRequest)

	// Send token to API with requests
	return myRequest
}

// hidAPIexample represents the minimal way for an API to authorize an incoming request.
// Requires, at a minimum, that env var HID_ADDR is set.
func hidAPIexample(myRequest *http.Request) {
	// Make reusable HID item
	myHID, err := hid.New()
	if err != nil {
		log.Fatal(err)
	}

	// Verify if token is valid. Invalid token throws an error
	err = myHID.AuthorizeRequest(myRequest, "audience", "scope")
	if err != nil {
		log.Fatal("Token is invalid")
	}

	// Handle the request
	log.Println("The request was successfull")
}
