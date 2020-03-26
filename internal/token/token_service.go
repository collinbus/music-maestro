package token

import (
	"log"
	"musicMaestro/internal/network"
	"musicMaestro/internal/persistence"
	"strings"
	"time"
)

const url = "https://accounts.spotify.com/api/token"

type Service struct{}

func (*Service) RequestApiToken(applicationData *persistence.ApplicationData) *persistence.ApplicationData {
	requestBody := createApiRequestBody(applicationData)
	success, err := network.Post(url, requestBody, NewApiTokenResponseMapper())

	if err != nil {
		log.Fatal(err)
	}

	tokenResponse := success.(*ApiTokenResponseBody)

	applicationData.AccessCode = tokenResponse.AccessToken
	applicationData.RefreshToken = tokenResponse.RefreshToken
	applicationData.TokenExpiration = network.CalculateExpirationDate(tokenResponse.ExpiresIn)
	return applicationData
}

func (*Service) RefreshApiToken(applicationData *persistence.ApplicationData) *persistence.ApplicationData {
	requestBody := createRefreshRequestBody(applicationData)
	success, err := network.Post(url, requestBody, NewRefreshTokenResponseMapper())

	if err != nil {
		log.Fatal(err)
	}

	tokenResponse := success.(*RefreshTokenResponseBody)

	applicationData.AccessCode = tokenResponse.AccessToken
	applicationData.TokenExpiration = network.CalculateExpirationDate(tokenResponse.ExpiresIn)
	return applicationData
}

func (*Service) IsTokenExpired(expiration string) bool {
	parsedExpirationTime, err := time.Parse("2006-01-02T15:04:05-0700", expiration)
	if err != nil {
		log.Fatal(err)
	}
	return parsedExpirationTime.Before(time.Now())
}

func createApiRequestBody(applicationData *persistence.ApplicationData) *strings.Reader {
	code := applicationData.AccessCode
	id := applicationData.ClientId
	secret := applicationData.ClientSecret

	return NewApiTokenRequestBody(code, id, secret)
}

func createRefreshRequestBody(applicationData *persistence.ApplicationData) *strings.Reader {
	refreshToken := applicationData.RefreshToken
	clientId := applicationData.ClientId
	clientSecret := applicationData.ClientSecret

	return NewRefreshTokenRequestBody(refreshToken, clientId, clientSecret)
}

func NewService() *Service {
	return &Service{}
}
