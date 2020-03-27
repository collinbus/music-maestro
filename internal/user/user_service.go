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

func (service Service) UpdateCurrentUserFromServer() {
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

func (Service) FetchUser() *domain.User {
	return persistence.GetUser()
}

func (Service) UpdateUser(user *domain.User) {
	persistence.SaveUser(user)
}

func NewService(tokenService *token.Service) *Service {
	return &Service{tokenService: tokenService}
}
