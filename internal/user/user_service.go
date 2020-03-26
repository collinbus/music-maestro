package user

import (
	"log"
	"musicMaestro/internal/domain"
	"musicMaestro/internal/network"
	"musicMaestro/internal/persistence"
	"musicMaestro/internal/token"
	"strings"
)

const url string = "https://api.spotify.com/v1/me"

type Service struct {
	tokenService *token.Service
}

func (service Service) UpdateCurrentUser() {
	authorizationToken := service.tokenService.GetAuthorizationToken()
	requestBody := strings.NewReader("")
	mapper := NewResponseMapper()
	response, err := network.Get(url, requestBody, mapper, authorizationToken)

	if err != nil {
		log.Fatal(err)
	}

	usr := response.(*domain.User)
	persistence.SaveUser(usr)
}

func NewService(tokenService *token.Service) *Service {
	return &Service{tokenService: tokenService}
}
