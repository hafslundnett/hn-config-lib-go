# hn-config-lib-go

GO-implemented config library

UNDER DEVELOPMENT

To testrun: set environment variables VAULT_ADDR as the address of vault, set VAULT_CACERT as the file location of the certificate for the vault server (.pem), set GIT_TOKEN as a github login token


Certificates
Get a freh CA certificate pool for http clients, with the contents of zero or more certificate files added to it.


Vault
Make and configure a HashiCorp Vault client, and use it to get secrets from Vault