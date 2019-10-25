# hn-config-lib-go

GO-implemented config library

UNDER DEVELOPMENT

To test or run: Set environment variables VAULT_ADDR as the address of vault and set GITHUB_TOKEN as a github login token.
If the Vault does not have a publicly signed Ca certificate, you may set VAULT_CACERT as the file location of the self-signed certificate for the vault server (.pem format).


Example usage:
(Requires environment variables to be set)

```
import (
	vault "hafslundnett/hn-config-lib-go"
	"log"
)

func main() {
	vault, err := vault.New()
	if err != nil {
		log.Fatal(err)
	}

	secret, err := vault.GetSecret("path/to/secret")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(secret)
}
```
