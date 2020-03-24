package network

import (
	"compress/gzip"
	"io/ioutil"
	"log"
	"net/http"
)

func DecompressResponse(response *http.Response) []byte {
	defer response.Body.Close()
	reader, err := gzip.NewReader(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	responseBytes, _ := ioutil.ReadAll(reader)
	return responseBytes
}
