package lib

import (
	"log"
	"net/http"

	"github.com/hafslundnett/hn-config-lib-go/hid"
	"github.com/hafslundnett/hn-config-lib-go/vault"
)

func demo() {
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
	mySecret, err := myVault.GetSecret("path/to/secret")
	if err != nil {
		log.Fatal(err)
	}

	// Do something with the secret
	log.Println(mySecret)
}

func hidClientDemo() *http.Request {
	// Make reusable HID item
	myHID, err := hid.New()
	if err != nil {
		log.Fatal(err)
	}

	// Get a bearer token from HID
	myToken, err := myHID.GetToken("username", "secret")
	if err != nil {
		log.Fatal(err)
	}

	// Make request struct as usual
	myRequest, err := http.NewRequest("POST", "api.url", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add bearer token to http request header
	myRequest.Header.Add("Authorization", myToken.Raw)

	// Send token to API with requests
	return myRequest
}

func hidAPIdemo(r *http.Request) {
	// Make reusable HID item
	myHID, err := hid.New()
	if err != nil {
		log.Fatal(err)
	}

	// Verify if token is valid. Invalid token throws an error
	err = myHID.AuthorizeRequest(r, "audience", "scope")
	if err != nil {
		log.Println("Token is invalid")
		log.Fatal(err)
	}

	// Handle the request
	log.Println("The request was successfull")
}
