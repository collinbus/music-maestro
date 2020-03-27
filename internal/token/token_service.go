package token

import (
	"log"
	"musicMaestro/internal/domain"
	"musicMaestro/internal/network"
	"musicMaestro/internal/persistence"
	"musicMaestro/internal/utils"
	"strings"
)

const url = "https://accounts.spotify.com/api/token"

type Service struct {
	appDataService persistence.ApplicationDataService
}

func (service *Service) GetAuthorizationToken() string {
	appData := service.appDataService.RetrieveApplicationData()
	if appData.RefreshToken == "" {
		appData = requestApiToken(appData)
		service.appDataService.SaveApplicationData(appData)
	} else if utils.IsAfter(appData.TokenExpiration) {
		appData = refreshApiToken(appData)
		service.appDataService.SaveApplicationData(appData)
	}
	return appData.AccessCode
}

func requestApiToken(applicationData *domain.ApplicationData) *domain.ApplicationData {
	requestBody := createApiRequestBody(applicationData)
	success, err := network.Post(url, requestBody, NewApiTokenResponseMapper())

	if err != nil {
		log.Fatal(err)
	}

	tokenResponse := success.(*ApiTokenResponseBody)

	applicationData.AccessCode = tokenResponse.AccessToken
	applicationData.RefreshToken = tokenResponse.RefreshToken
	applicationData.TokenExpiration = utils.CalculateExpirationDate(tokenResponse.ExpiresIn)
	return applicationData
}

func refreshApiToken(applicationData *domain.ApplicationData) *domain.ApplicationData {
	requestBody := createRefreshRequestBody(applicationData)
	success, err := network.Post(url, requestBody, NewRefreshTokenResponseMapper())

	if err != nil {
		log.Fatal(err)
	}

	tokenResponse := success.(*RefreshTokenResponseBody)

	applicationData.AccessCode = tokenResponse.AccessToken
	applicationData.TokenExpiration = utils.CalculateExpirationDate(tokenResponse.ExpiresIn)
	return applicationData
}

func createApiRequestBody(applicationData *domain.ApplicationData) *strings.Reader {
	code := applicationData.AccessCode
	id := applicationData.ClientId
	secret := applicationData.ClientSecret

	return NewApiTokenRequestBody(code, id, secret)
}

func createRefreshRequestBody(applicationData *domain.ApplicationData) *strings.Reader {
	refreshToken := applicationData.RefreshToken
	clientId := applicationData.ClientId
	clientSecret := applicationData.ClientSecret

	return NewRefreshTokenRequestBody(refreshToken, clientId, clientSecret)
}

func NewService() *Service {
	return &Service{}
}
