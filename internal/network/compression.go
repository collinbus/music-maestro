package network

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
)

func DecompressResponse(response io.ReadCloser) []byte {
	defer response.Close()
	reader, err := gzip.NewReader(response)

	if err != nil {
		log.Fatal(err)
	}

	responseBytes := getBytesFrom(reader)
	return responseBytes
}

func getBytesFrom(reader *gzip.Reader) []byte {
	all, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Fatal(err)
	}

	return all
}
