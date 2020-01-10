# HN Config Lib GO  
  
GO-implementation of Hashicorp Vault, HID and additional related helper libraries for HafslundNett  
  
UNDER DEVELOPMENT

## Examples  

example_test.go

Examples:  
Make sure environment variables are set before running.  
See demo.go for a more comprehensive example  

Vault example:  
VAULT_ADDR: The address of the vault. If not set (or empty) it will default to localhost.  
GITHUB_TOKEN: A github login token. If not set (or empty) it will default to use K8 to login.  
VAULT_CACERT: If the Vault does not have a publicly signed Ca certificate, you may set VAULT_CACERT as the file location of the self-signed certificate for the vault server (.pem format).

```GO
import (
    "hafslundnett/hn-config-lib-go/vault"
    "log"
)

func main() {
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
```
  
## License  
  
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details  
