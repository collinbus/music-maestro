package network

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		println(fmt.Errorf(err.Error()))
		return applicationData
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
	reader, err := gzip.NewReader(response.Body)
	defer response.Body.Close()
	all, err := ioutil.ReadAll(reader)
	println(string(all))
	err = json.NewDecoder(bytes.NewReader(all)).Decode(responseBody)

	if err != nil {
		println(fmt.Errorf(err.Error()))
		return nil
	}
	return responseBody
}

func calculateExpirationDate(expiresIn int) string {
	now := time.Now()
	expirationDuration := time.Duration(expiresIn) * time.Second
	now.Add(expirationDuration)
	return now.String()
}
