package lib

import (
	"log"
	"net/http"
	"os"

	"github.com/hafslundnett/hn-config-lib-go/hid"
	"github.com/hafslundnett/hn-config-lib-go/vault"
)

// demo is a runnable version of the example_test.go that requires appropriate env vars to be set before executing. See readme for documentation on these.
// func main() { // <- Uncomment to run
func demo() { //    <- Comment to run
	vaultDemo()

	myRequest := hidClientDemo()

	hidAPIdemo(myRequest)
}

func vaultDemo() {
	// Make reusable vault item
	myVault, err := vault.New()
	if err != nil {
		log.Fatal(err)
	}

	// Get a secret from the vault
	mySecret, err := myVault.GetSecret(os.Getenv("TEST_SECRET"))
	if err != nil {
		log.Fatal(err)
	}

	// Do something with the secret
	log.Println(mySecret.RequestID)
}

func hidClientDemo() *http.Request {
	// Make reusable HID item
	myHID, err := hid.New()
	if err != nil {
		log.Fatal(err)
	}

	// Get a bearer token from HID
	myToken, err := myHID.GetToken(os.Getenv("TEST_HID_ID"), os.Getenv("TEST_HID_SECRET"))
	if err != nil {
		log.Fatal(err)
	}

	// Make http.Request as usual
	myRequest, err := http.NewRequest("", "api.url", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add bearer token to http request header
	myToken.AppendToRequest(myRequest)

	// Send token to API with requests
	return myRequest
}

func hidAPIdemo(myRequest *http.Request) {
	// Make reusable HID item
	myHID, err := hid.New()
	if err != nil {
		log.Fatal(err)
	}

	// Verify if token is valid. Invalid token throws an error
	err = myHID.AuthorizeRequest(myRequest, os.Getenv("TEST_AUDIENCE"), os.Getenv("TEST_SCOPE"))
	if err != nil {
		log.Fatal("Token is invalid")
	}

	// Handle the request
	log.Println("The request was successfull")
}
