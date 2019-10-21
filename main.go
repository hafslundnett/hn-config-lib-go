package main

import (
	"fmt"
	"hafslundnett/hn-config-lib-go/vault"
	"log"
)

func main() {
	cfgFile := "C:/secret/config.yaml"

	vault, err := vault.NewVault(cfgFile)
	if err != nil {
		log.Fatalf("while configuring vault: %v", err)
	}

	fmt.Println(vault)
	fmt.Println()

	secret, err := vault.GetSecret("documentor/kv/storage%2Fdocumentationdata")
	if err != nil {
		log.Fatalf("while getting secret: %v", err)
	}

	fmt.Println(secret)
}
