package network

import (
	"io"
	"log"
	"net/http"
)

const contentType = "application/x-www-form-urlencoded"

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
	all := DecompressResponse(response.Body)
	if response.StatusCode == 200 {
		return mapper.MapSuccess(all), nil
	} else {
		return nil, mapper.MapError(all)
	}
}

func addRequestHeaders(request *http.Request) {
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept-Encoding", "gzip, deflate, br")
}
