package network

import "strings"

type RequestMapper interface {
	CreateRequestBody() *strings.Reader
}

type ResponseMapper interface {
	MapSuccess([]byte) interface{}
	MapError([]byte) error
}
