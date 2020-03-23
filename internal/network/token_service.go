package network

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
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
	responseBody := NewApiTokenResponseBody()
	all, err := decompressResponse(response)
	println(string(all))
	err = json.NewDecoder(bytes.NewReader(all)).Decode(responseBody)

	if err != nil {
		println(fmt.Errorf(err.Error()))
		return nil
	}
	return responseBody
}

func decompressResponse(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	reader, err := gzip.NewReader(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return ioutil.ReadAll(reader)
}

func calculateExpirationDate(expiresIn int) string {
	now := time.Now()
	expirationDuration := time.Duration(expiresIn) * time.Second
	return now.Add(expirationDuration).String()
}
