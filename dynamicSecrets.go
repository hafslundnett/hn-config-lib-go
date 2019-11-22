package vault

import (
	"log"
	"time"
)

const (
	tenSeconds = 10 * time.Second
)

// SecretsSubscriber implementors have are dependant on secrets (connections strings,
// service account credentials and similar), and want the dynamic aspects to be
// handled automatically.
type SecretsSubscriber interface{
	GetSubscriptionSpec() SecretSubscriptionSpec

	// Here we assume that the subscriber starts its own
	// go routine for receiving updated secrets on the channel
	StartSecretsListener()
}

type secretGetter interface {
	GetSecret(path string) (*Secret, error)
}

// SecretSubscriptionSpec a specification of the paths to the secrets that a
// subscriber is interested in
type SecretSubscriptionSpec struct {
	Paths []string
	CallbackChan chan<- UpdatedSecret
}

// UpdatedSecret a new version of a secret
type UpdatedSecret struct {
	Path string
	Secrets map[string]*Secret
}

func (us UpdatedSecret) getAllData() map[string]string {
	res := map[string]string{}
	for _, m := range us.Secrets {
		for k, v := range m.Data {
			res[k] = v
		}
	}
	return res
}

// RegisterDynamicSecretDependency by registering
func RegisterDynamicSecretDependency(dep SecretsSubscriber, vlt *Vault, dc chan<- bool) {
	dep.StartSecretsListener()
	spec := dep.GetSubscriptionSpec()
	for _, path := range spec.Paths {
		go func(p string, c chan<- UpdatedSecret, v *Vault) {
			maintainer := singleSecretMaintainer{
				path:         p,
				callbackChan: c,
				v:            v,
				doneChan:     dc,
			}
			maintainer.start()
		}(path, spec.CallbackChan, vlt)
	}
}

type singleSecretMaintainer struct {
	path string
	callbackChan chan<- UpdatedSecret
	v secretGetter
	doneChan chan<- bool
}

func (m singleSecretMaintainer) start() {
	d, renewable := m.doIteration()
	if renewable {
		for {
			w := getWaitDuration(d)
			time.Sleep(w)
			d, _ = m.doIteration()
			if d <= 0 {
				// Exit loop, mostly for testing purposes
				if m.doneChan != nil {
					m.doneChan <- true
				}

				return
			}
		}
	}
}

func (m singleSecretMaintainer) doIteration() (time.Duration, bool) {
	secret, err := m.v.GetSecret(m.path)
	if err != nil {
		log.Printf("Error while getting secret %s :: %v", m.path, err)
		return time.Second * 10, true
	}
	secrets := map[string]*Secret{m.path: secret}
	if sp, ok := secret.Data["secret-path"]; ok {
		innerSecret, err := m.v.GetSecret(sp)
		if err != nil {
			secrets[sp] = innerSecret
		}
	}

	us := UpdatedSecret{
		Path: m.path,
		Secrets: secrets,
	}
	m.callbackChan <- us
	d := time.Millisecond * 0
	if secret.LeaseDuration > 0 {
		d = time.Millisecond * time.Duration(secret.LeaseDuration)
	}

	return d, secret.Renewable
}


func getWaitDuration(d time.Duration) time.Duration {
	if d <= tenSeconds {
		return d
	}

	return d - tenSeconds
}