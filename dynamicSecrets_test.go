package vault

import "testing"

func Test_singleSecretMaintainer_start_notRenewable_returnsAtOnce(t *testing.T) {
	// Arrange
	receivedSecrets := []UpdatedSecret{}
	s := &Secret{
		Renewable:     false,
		LeaseDuration: 0,
	}
	g := &mockSecretGetter{
		invocationCount: 0,
		returnSecrets:   []*Secret{s},
	}
	c := make(chan UpdatedSecret)
	p := "test/kv/secret"
	m := singleSecretMaintainer{
		path: p,
		callbackChan: c,
		v: g,
	}

	go func(){
		us := <- c
		receivedSecrets = append(receivedSecrets, us)
	}()

	// Act
	m.start()

	// Assert
	if len(receivedSecrets) != 1 {
		t.Errorf("Expected 1 received secret. Got %d", len(receivedSecrets))
	}
}

type mockSecretGetter struct {
	invocationCount int
	returnSecrets []*Secret
}

func (m *mockSecretGetter) GetSecret(path string) (*Secret, error) {
	if len(m.returnSecrets) == 1 {
		ss := m.returnSecrets[0]
		m.returnSecrets = []*Secret{}
		return ss, nil
	}

	s, rs := m.returnSecrets[0], m.returnSecrets[1:]
	m.returnSecrets = rs
	return s, nil
}