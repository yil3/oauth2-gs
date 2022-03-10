package user

import (
	"oauth2-gs/oauth"
)

// Service struct keeps db object to avoid passing it around
type Service struct {
	oauthService oauth.ServiceInterface
}

// NewService returns a new Service instance
func NewService(oauthService oauth.ServiceInterface) *Service {
	return &Service{
		oauthService: oauthService,
	}
}

// GetOauthService returns oauth.Service instance
func (s *Service) GetOauthService() oauth.ServiceInterface {
	return s.oauthService
}

// Close stops any running services
func (s *Service) Close() {}
