package network

import (
	"io"
	"io/ioutil"
	"net/http"
)

func DownloadImage(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		println("Unable to fetch image")
		return []byte{}
	}

	defer response.Body.Close()
	return readBytes(response.Body)
}

func readBytes(reader io.ReadCloser) []byte {
	all, err := ioutil.ReadAll(reader)
	if err != nil {
		println("Unable to read image")
		return []byte{}
	}
	return all
}
