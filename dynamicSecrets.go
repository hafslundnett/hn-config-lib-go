package vault

import (
	"log"
	"time"
)

// SecretsSubscriber implementors have are dependant on secrets (connections strings,
// service account credentials and similar), and want the dynamic aspects to be
// handled automatically.
type SecretsSubscriber interface{
	GetSubscriptionSpec() SecretSubscriptionSpec

	// Here we assume that the subscriber starts its own
	// go routine for receiving updated secrets on the channel
	Start()
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
	Secret *Secret
}

// RegisterDynamicSecretDependency by registering
func RegisterDynamicSecretDependency(dep SecretsSubscriber, vlt *Vault) {
	spec := dep.GetSubscriptionSpec()
	for _, path := range spec.Paths {
		go func(p string, c chan<- UpdatedSecret, v *Vault) {
			maintainer := singleSecretMaintainer{
				path:         p,
				callbackChan: c,
				v:            v,
			}
			maintainer.start()
		}(path, spec.CallbackChan, vlt)
	}
}

type singleSecretMaintainer struct {
	path string
	callbackChan chan<- UpdatedSecret
	v secretGetter
}

func (m singleSecretMaintainer) start() {
	d, renewable := m.doIteration()
	if renewable {
		for {
			time.Sleep(d)
			d, _ = m.doIteration()
			if d <= 0 {
				// Exit loop, mostly for testing purposes
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

	us := UpdatedSecret{
		Path: m.path,
		Secret: secret,
	}
	m.callbackChan <- us
	d := time.Millisecond * 0
	if secret.LeaseDuration > 0 {
		d = time.Millisecond * time.Duration(secret.LeaseDuration)
	}

	return d, secret.Renewable
}