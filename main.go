//TODO: delete
package main

import (
	"fmt"
	"hafslundnett/hn-config-lib-go/vault"
	"log"
)

//temporary example implementation
func main() {
	vault, err := vault.New()
	if err != nil {
		log.Fatalf("while configuring vault: %v", err)
	}

	path := "documentor/kv/storage%2Fdocumentationdata"

	secret, err := vault.GetSecret(path)
	if err != nil {
		log.Fatalf("while getting secret: %v", err)
	}

	fmt.Println(secret)
}

//TODO: delete
