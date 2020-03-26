package token

import (
	"errors"
	"musicMaestro/internal/network"
	"strings"
)

func catchError(data []byte) error {
	errorResponseBody := &ErrorResponseBody{}
	network.DecodeJson(data, errorResponseBody)
	errorMessage := errorResponseBody.Error + ": " + errorResponseBody.Description
	return errors.New(errorMessage)
}

type ApiTokenResponseMapper struct{}

func (ApiTokenResponseMapper) MapSuccess(data []byte) interface{} {
	responseBody := &ApiTokenResponseBody{}
	network.DecodeJson(data, responseBody)
	return responseBody
}

func (ApiTokenResponseMapper) MapError(data []byte) error {
	return catchError(data)
}

func NewApiTokenResponseMapper() *ApiTokenResponseMapper {
	return &ApiTokenResponseMapper{}
}

type RefreshTokenResponseMapper struct{}

func (RefreshTokenResponseMapper) MapSuccess(data []byte) interface{} {
	responseBody := &RefreshTokenResponseBody{}
	network.DecodeJson(data, responseBody)
	return responseBody
}

func (RefreshTokenResponseMapper) MapError(data []byte) error {
	return catchError(data)
}

func NewRefreshTokenResponseMapper() *RefreshTokenResponseMapper {
	return &RefreshTokenResponseMapper{}
}

type ApiTokenRequestMapper struct {
	code         string
	clientId     string
	clientSecret string
}

func (mapper *ApiTokenRequestMapper) CreateRequestBody() *strings.Reader {
	return NewApiTokenRequestBody(mapper.code, mapper.clientId, mapper.clientSecret)
}

type RefreshTokenRequestMapper struct {
	code         string
	clientId     string
	clientSecret string
}

func (mapper *RefreshTokenRequestMapper) CreateRequestBody() *strings.Reader {
	return NewRefreshTokenRequestBody(mapper.code, mapper.clientId, mapper.clientSecret)
}
