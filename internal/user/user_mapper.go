package user

import (
	"errors"
	"musicMaestro/internal/network"
	"musicMaestro/internal/token"
)

type ResponseMapper struct{}

func catchError(data []byte) error {
	errorResponseBody := &token.ErrorResponseBody{}
	network.DecodeJson(data, errorResponseBody)
	errorMessage := errorResponseBody.Error + ": " + errorResponseBody.Description
	return errors.New(errorMessage)
}

func (ResponseMapper) MapSuccess(data []byte) interface{} {
	responseBody := &GetUserInfoResponse{}
	network.DecodeJson(data, responseBody)
	return NewUser(responseBody.Id, responseBody.Name)
}

func (ResponseMapper) MapError(data []byte) error {
	return catchError(data)
}

func NewResponseMapper() *ResponseMapper {
	return &ResponseMapper{}
}
