package service

// Service expl
type Service interface {
	NewConfig() error
	NewClient() error
	Authenticate() error
}

// New expl
func New() (s Service, err error) {
	if err = s.NewConfig(); err != nil {
		return
	}

	if err = s.NewClient(); err != nil {
		return
	}

	if err = s.Authenticate(); err != nil {
		return
	}

	return
}
