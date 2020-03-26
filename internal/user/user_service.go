package user

import (
	"musicMaestro/internal/persistence"
	"musicMaestro/internal/token"
)

type Service struct {
	appDataService *persistence.ApplicationDataService
	tokenService   *token.Service
}

func (service Service) UpdateCurrentUser() {
	service.checkForTokenUpdate()
}

func (service Service) checkForTokenUpdate() {
	appData := service.appDataService.RetrieveApplicationData()
	if appData.RefreshToken == "" {
		appData = service.tokenService.RequestApiToken(appData)
		service.appDataService.SaveApplicationData(appData)
	} else if service.tokenService.IsTokenExpired(appData.TokenExpiration) {
		appData = service.tokenService.RefreshApiToken(appData)
		service.appDataService.SaveApplicationData(appData)
	}
}

func NewService(appDataService *persistence.ApplicationDataService, tokenService *token.Service) *Service {
	return &Service{appDataService: appDataService, tokenService: tokenService}
}
