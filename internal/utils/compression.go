package utils

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
)

func Decompress(reader io.ReadCloser) []byte {
	defer reader.Close()
	gzipReader, err := gzip.NewReader(reader)

	if err != nil {
		log.Fatal(err)
	}

	responseBytes := getBytesFrom(gzipReader)
	return responseBytes
}

func getBytesFrom(reader *gzip.Reader) []byte {
	all, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Fatal(err)
	}

	return all
}
