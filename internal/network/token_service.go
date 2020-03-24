package network

import (
	"log"
	"musicMaestro/internal/persistence"
	"net/http"
	"strings"
	"time"
)

const url = "https://accounts.spotify.com/api/token"
const contentType = "application/x-www-form-urlencoded"

func RequestApiToken(applicationData *persistence.ApplicationData) *persistence.ApplicationData {
	requestBody := createApiRequestBody(applicationData)
	request, _ := http.NewRequest(http.MethodPost, url, requestBody)
	addRequestHeaders(request)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	tokenResponse := parseApiTokenResponse(response)

	applicationData.AccessCode = tokenResponse.AccessToken
	applicationData.RefreshToken = tokenResponse.RefreshToken
	applicationData.TokenExpiration = CalculateExpirationDate(tokenResponse.ExpiresIn)
	return applicationData
}

func RefreshApiToken(applicationData *persistence.ApplicationData) *persistence.ApplicationData {
	requestBody := createRefreshRequestBody(applicationData)
	request, _ := http.NewRequest(http.MethodPost, url, requestBody)

	addRequestHeaders(request)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	tokenResponse := parseRefreshTokenResponse(response)

	applicationData.AccessCode = tokenResponse.AccessToken
	applicationData.TokenExpiration = CalculateExpirationDate(tokenResponse.ExpiresIn)
	return applicationData
}

func createApiRequestBody(applicationData *persistence.ApplicationData) *strings.Reader {
	code := applicationData.AccessCode
	id := applicationData.ClientId
	secret := applicationData.ClientSecret

	return NewApiTokenRequestBody(code, id, secret)
}

func parseApiTokenResponse(response *http.Response) *ApiTokenResponseBody {
	all := DecompressResponse(response)
	responseBody := parseApiTokenJsonResponse(all, response.StatusCode)
	return responseBody
}

func parseApiTokenJsonResponse(data []byte, statusCode int) *ApiTokenResponseBody {
	if statusCode == 200 {
		responseBody := NewApiTokenResponseBody()
		DecodeJson(data, responseBody)
		return responseBody
	} else {
		errorResponseBody := NewErrorResponseBody()
		DecodeJson(data, errorResponseBody)
		log.Fatalf("%s: %s", errorResponseBody.Error, errorResponseBody.Description)
		return nil
	}
}

func createRefreshRequestBody(applicationData *persistence.ApplicationData) *strings.Reader {
	refreshToken := applicationData.RefreshToken
	clientId := applicationData.ClientId
	clientSecret := applicationData.ClientSecret

	return NewRefreshTokenRequestBody(refreshToken, clientId, clientSecret)
}

func parseRefreshTokenResponse(response *http.Response) *RefreshTokenResponseBody {
	all := DecompressResponse(response)
	responseBody := parseRefreshTokenJsonResponse(all, response.StatusCode)
	return responseBody
}

func parseRefreshTokenJsonResponse(data []byte, statusCode int) *RefreshTokenResponseBody {
	if statusCode == 200 {
		responseBody := NewRefreshTokenResponseBody()
		DecodeJson(data, responseBody)
		return responseBody
	} else {
		errorResponseBody := NewErrorResponseBody()
		DecodeJson(data, errorResponseBody)
		log.Fatalf("%s: %s", errorResponseBody.Error, errorResponseBody.Description)
		return nil
	}
}

func addRequestHeaders(request *http.Request) {
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept-Encoding", "gzip, deflate, br")
}
