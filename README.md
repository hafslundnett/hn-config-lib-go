# hn-config-lib-go  
  
GO-implementation of Vault for HafslundNett  
  
UNDER DEVELOPMENT  
  
To run: Make sure environment variables are set before running. 
  
VAULT_ADDR: The address of the vault. If not set (or empty) it will default to localhost.  
GITHUB_TOKEN: A github login token. If not set (or empty) it will default to use K8 to login.  
VAULT_CACERT: If the Vault does not have a publicly signed Ca certificate, you may set VAULT_CACERT as the file location of the self-signed certificate for the vault server (.pem format).

Example usage:  
(Requires environment variables to be set)  

```
import (
	vault "hafslundnett/hn-config-lib-go"
	"log"
)

func main() {
	myVault, err := vault.New()
	if err != nil {
		log.Fatal(err)
	}

	secret, err := myVault.GetSecret("path/to/secret")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(secret)
}
```
