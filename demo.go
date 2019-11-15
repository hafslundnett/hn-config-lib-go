package lib

import (
	"hafslundnett/x/hn-config-lib/hid"
	"hafslundnett/x/hn-config-lib/vault"
	"log"
)

func main() {
	vaultDemo()

	token := hidClientDemo()

	hidAPIdemo(*token)
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

func hidClientDemo() *hid.Token {
	// Make reusable HID item
	myHID := hid.New("HIDaddress.com", "path/to/somewhere", "somesecret")

	// Get a token from HID
	myToken, err := myHID.GetToken("someUser")
	if err != nil {
		log.Fatal(err)
	}

	// Send token to API with requests
	return myToken
}

func hidAPIdemo(token hid.Token) {
	// Make reusable HID item
	myHID := hid.New("HIDaddress.com", "path/to/somewhere/else", "someothersecret")

	// Verify if token is valid
	valid, err := myHID.Introspection(token)
	if err != nil {
		log.Fatal(err)
	}

	// Do something with this information
	log.Println(valid)
}
