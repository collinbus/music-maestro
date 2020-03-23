package network

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
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
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept-Encoding", "gzip, deflate, br")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	tokenResponse := parseApiTokenResponse(response)

	applicationData.RefreshToken = tokenResponse.RefreshToken
	applicationData.TokenExpiration = calculateExpirationDate(tokenResponse.ExpiresIn)
	return applicationData
}

func createApiRequestBody(applicationData *persistence.ApplicationData) *strings.Reader {
	code := applicationData.AccessCode
	id := applicationData.ClientId
	secret := applicationData.ClientSecret

	return NewApiTokenRequestBody(code, id, secret)
}

func parseApiTokenResponse(response *http.Response) *ApiTokenResponseBody {
	all := decompressResponse(response)
	responseBody := parseApiTokenJsonResponse(all, response.StatusCode)
	return responseBody
}

func parseApiTokenJsonResponse(data []byte, statusCode int) *ApiTokenResponseBody {
	if statusCode == 200 {
		responseBody := NewApiTokenResponseBody()
		decodeJson(data, responseBody)
		return responseBody
	} else {
		errorResponseBody := NewErrorResponseBody()
		decodeJson(data, errorResponseBody)
		log.Fatalf("%s: %s", errorResponseBody.Error, errorResponseBody.Description)
		return nil
	}
}

func decodeJson(data []byte, responseBody interface{}) {
	err := json.NewDecoder(bytes.NewReader(data)).Decode(responseBody)
	if err != nil {
		log.Fatal(err)
	}
}

func decompressResponse(response *http.Response) []byte {
	defer response.Body.Close()
	reader, err := gzip.NewReader(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	responseBytes, _ := ioutil.ReadAll(reader)
	return responseBytes
}

func calculateExpirationDate(expiresIn int) string {
	now := time.Now()
	expirationDuration := time.Duration(expiresIn) * time.Second
	return now.Add(expirationDuration).String()
}
