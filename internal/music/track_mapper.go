package music

import (
	"errors"
	"musicMaestro/internal/network"
	"musicMaestro/internal/token"
)

type TrackResponseMapper struct{}

func catchError(data []byte) error {
	errorResponseBody := &token.ErrorResponseBody{}
	network.DecodeJson(data, errorResponseBody)
	errorMessage := errorResponseBody.Error + ": " + errorResponseBody.Description
	return errors.New(errorMessage)
}

func (TrackResponseMapper) MapSuccess(data []byte) interface{} {
	responseBody := &GetUserTracksResponse{}
	network.DecodeJson(data, responseBody)
	return responseBody
}

func (TrackResponseMapper) MapError(data []byte) error {
	return catchError(data)
}

func NewTrackResponseMapper() *TrackResponseMapper {
	return &TrackResponseMapper{}
}
