package user

import (
	"errors"
	"musicMaestro/internal/domain"
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

	return createUser(responseBody)
}

func (ResponseMapper) MapError(data []byte) error {
	return catchError(data)
}

func createUser(responseBody *GetUserInfoResponse) interface{} {
	id := responseBody.Id
	name := responseBody.Name
	followers := responseBody.Followers.Total
	imageUrl := responseBody.Images[0].Url

	userUrl := responseBody.ExternalUrls.SpotifyUserUrl
	apiUrl := responseBody.Link
	uri := responseBody.Uri

	urls := domain.NewUrls(apiUrl, userUrl, uri)
	image := domain.NewImage(imageUrl, "")
	return domain.NewUser(id, name, urls, image, followers)
}

func NewResponseMapper() *ResponseMapper {
	return &ResponseMapper{}
}
