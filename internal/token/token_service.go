package token

import (
	"log"
	"musicMaestro/internal/network"
	"musicMaestro/internal/persistence"
	"strings"
)

const url = "https://accounts.spotify.com/api/token"

func RequestApiToken(applicationData *persistence.ApplicationData) *persistence.ApplicationData {
	requestBody := createApiRequestBody(applicationData)
	success, err := network.Post(url, requestBody, NewApiTokenResponseMapper())

	if err != nil {
		log.Fatal(err)
	}

	tokenResponse := success.(ApiTokenResponseBody)

	applicationData.AccessCode = tokenResponse.AccessToken
	applicationData.RefreshToken = tokenResponse.RefreshToken
	applicationData.TokenExpiration = network.CalculateExpirationDate(tokenResponse.ExpiresIn)
	return applicationData
}

func RefreshApiToken(applicationData *persistence.ApplicationData) *persistence.ApplicationData {
	requestBody := createRefreshRequestBody(applicationData)
	success, err := network.Post(url, requestBody, NewRefreshTokenResponseMapper())

	if err != nil {
		log.Fatal(err)
	}

	tokenResponse := success.(RefreshTokenResponseBody)

	applicationData.AccessCode = tokenResponse.AccessToken
	applicationData.TokenExpiration = network.CalculateExpirationDate(tokenResponse.ExpiresIn)
	return applicationData
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
