package user

import (
	"musicMaestro/internal/token"
)

type Service struct {
	tokenService *token.Service
}

func (service Service) UpdateCurrentUser() {
	service.tokenService.GetAuthorizationToken()
}

func NewService(tokenService *token.Service) *Service {
	return &Service{tokenService: tokenService}
}
