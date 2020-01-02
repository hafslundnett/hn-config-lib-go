package service

// Service expl
type Service interface {
	Configure(*Config) error
	GetEnvars() map[string]string
	GetAccess() error
}

// Setup expl
func Setup(service Service) (Service, error) {
	Configure(service, service.GetEnvars())

	err := service.GetAccess()

	return service, err
}
