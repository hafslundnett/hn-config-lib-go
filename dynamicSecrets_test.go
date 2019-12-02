package vault

import (
	"testing"
	"time"
)

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
		path:         p,
		callbackChan: c,
		v:            g,
	}

	go func() {
		us := <-c
		receivedSecrets = append(receivedSecrets, us)
	}()

	// Act
	m.start()

	// Assert
	if len(receivedSecrets) != 1 {
		t.Errorf("Expected 1 received secret. Got %d", len(receivedSecrets))
	}
}

func Test_singleSecretMaintainer_start_renewable_iteratesAsExpected(t *testing.T) {
	// Arrange
	s1 := &Secret{
		Renewable:     true,
		LeaseDuration: 10100,
	}
	s2 := &Secret{
		Renewable:     true,
		LeaseDuration: -1,
	}
	g := &mockSecretGetter{
		invocationCount: 0,
		returnSecrets:   []*Secret{s1, s2},
	}
	c := make(chan UpdatedSecret)
	p := "test/kv/secret"
	doneChan := make(chan bool)
	m := singleSecretMaintainer{
		path:         p,
		callbackChan: c,
		v:            g,
		doneChan:     doneChan,
	}

	receivedSecrets := []UpdatedSecret{}
	go func() {
		for {
			select {
			case us := <-c:
				receivedSecrets = append(receivedSecrets, us)
			case <-doneChan:
				return
			}
		}
	}()

	// Act
	m.start()

	// Assert
	if len(receivedSecrets) != 2 {
		t.Errorf("Expected 2 received secret. Got %d", len(receivedSecrets))
	}
}

func Test_getWaitDuration(t *testing.T) {
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{name: "lessThan10", args: args{d: 9 * time.Second}, want: 9 * time.Second},
		{name: "moreThan10", args: args{d: 11 * time.Second}, want: 1 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWaitDuration(tt.args.d); got != tt.want {
				t.Errorf("getWaitDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockSecretGetter struct {
	invocationCount int
	returnSecrets   []*Secret
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
