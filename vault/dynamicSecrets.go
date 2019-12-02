package vault

import (
	"log"
	"time"
)

const (
	tenSeconds = 10 * time.Second // GO-LINT: var tenSeconds is of type time.Duration; don't use unit-specific suffix "Seconds"
)

// SecretsSubscriber implementors have are dependant on secrets (connections strings,
// service account credentials and similar), and want the dynamic aspects to be
// handled automatically.
type SecretsSubscriber interface {
	GetSubscriptionSpec() SecretSubscriptionSpec

	ReceiveAtStartup(UpdatedSecret)

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
	Paths        []string
	CallbackChan chan<- UpdatedSecret
}

// UpdatedSecret a new version of a secret
type UpdatedSecret struct {
	Path    string
	Secrets map[string]*Secret
}

// GetAllData combines all data in all secrets to a single map
func (us UpdatedSecret) GetAllData() map[string]string {
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
	spec := dep.GetSubscriptionSpec()
	maintainers := []singleSecretMaintainer{}
	for _, path := range spec.Paths {
		maintainer := singleSecretMaintainer{
			path:         path,
			callbackChan: spec.CallbackChan,
			v:            vlt,
			doneChan:     dc,
		}
		s, renewable, ttl, err := maintainer.getSecret()
		if err != nil {
			log.Fatal(err)
		}
		maintainer.setInitialTtl(ttl)
		dep.ReceiveAtStartup(s)
		if renewable {
			maintainers = append(maintainers, maintainer)
		}
	}

	dep.StartSecretsListener()

	for _, mt := range maintainers {
		go func(m singleSecretMaintainer) {
			m.start()
		}(mt)
	}
}

type singleSecretMaintainer struct {
	path         string
	callbackChan chan<- UpdatedSecret
	v            secretGetter
	doneChan     chan<- bool
	initialTtl   time.Duration
}

func (m *singleSecretMaintainer) setInitialTtl(ttl time.Duration) {
	m.initialTtl = ttl
}

func (m singleSecretMaintainer) start() {
	d := m.initialTtl
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

func (m singleSecretMaintainer) doIteration() (time.Duration, bool) {
	us, renewable, ttl, _ := m.getSecret()
	m.callbackChan <- us
	return ttl, renewable
}

func (m singleSecretMaintainer) getSecret() (UpdatedSecret, bool, time.Duration, error) {
	ttl := time.Hour * 8760 // 1 year
	renewable := false
	secret, err := m.v.GetSecret(m.path)
	if secret.Renewable {
		renewable = true
		ttl = time.Duration(secret.LeaseDuration) * time.Millisecond
	}
	if err != nil {
		log.Printf("Error while getting secret %s :: %v", m.path, err)
		return UpdatedSecret{}, false, time.Second * 0, err
	}
	secrets := map[string]*Secret{m.path: secret}
	if sp, ok := secret.Data["secret-path"]; ok {
		innerSecret, err := m.v.GetSecret(sp)
		if err == nil {
			secrets[sp] = innerSecret
		}
		if innerSecret.Renewable {
			renewable = true
			ttl2 := time.Duration(innerSecret.LeaseDuration) * time.Millisecond
			if ttl2 < ttl {
				ttl = ttl
			}
		}
	}

	us := UpdatedSecret{
		Path:    m.path,
		Secrets: secrets,
	}
	return us, renewable, ttl, nil
}

func getWaitDuration(d time.Duration) time.Duration {
	if d <= tenSeconds {
		return d
	}

	return d - tenSeconds
}
