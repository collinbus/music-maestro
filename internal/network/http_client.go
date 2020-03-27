package network

import (
	"io"
	"log"
	"musicMaestro/internal/utils"
	"net/http"
)

const contentType = "application/x-www-form-urlencoded"

func Get(url string, requestBody io.Reader, mapper ResponseMapper, authenticationToken string) (interface{}, error) {
	request, _ := http.NewRequest(http.MethodGet, url, requestBody)
	addAuthHeader(request, authenticationToken)
	addRequestHeaders(request)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	return parseResponse(response, mapper)
}

func Post(url string, requestBody io.Reader, mapper ResponseMapper) (interface{}, error) {
	request, _ := http.NewRequest(http.MethodPost, url, requestBody)
	addRequestHeaders(request)
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	return parseResponse(response, mapper)
}

func parseResponse(response *http.Response, mapper ResponseMapper) (interface{}, error) {
	all := utils.Decompress(response.Body)
	if response.StatusCode == 200 {
		return mapper.MapSuccess(all), nil
	} else {
		return nil, mapper.MapError(all)
	}
}

func addAuthHeader(request *http.Request, authenticationToken string) {
	request.Header.Add("Authorization", "Bearer "+authenticationToken)
}

func addRequestHeaders(request *http.Request) {
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept-Encoding", "gzip, deflate, br")
}
