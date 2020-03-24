package network

import (
	"bytes"
	"encoding/json"
	"log"
)

func DecodeJson(data []byte, responseBody interface{}) {
	err := json.NewDecoder(bytes.NewReader(data)).Decode(responseBody)
	if err != nil {
		log.Fatal(err)
	}
}
