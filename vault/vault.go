package vault

//Vault contains all information needed to get and interact with Vault secrets, after initial configuration.
type Vault struct {
	Config Config
	Client Client
	Token  Token
}

//New initiaizes a new Vault. It reads configuration information from the provided file, and populates a new Vault struct
//with the information needed to interact with secrets.
func New(cfgFile string) (*Vault, error) {
	vault := new(Vault)

	if err := vault.NewConfig(cfgFile); err != nil {
		return vault, err
	}

	if err := vault.NewClient(); err != nil {
		return vault, err
	}

	if err := vault.Authenticate(); err != nil {
		return vault, err
	}

	return vault, nil
}
