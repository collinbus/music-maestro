package image

import (
	"io/ioutil"
	"log"
	"musicMaestro/internal/network"
	"musicMaestro/internal/user"
)

type Service struct {
	userService *user.Service
}

func (service *Service) DownloadUserImage() {
	fetchedUser := service.userService.FetchUser()
	imgBytes := network.DownloadImage(fetchedUser.ImageUrl)
	err := ioutil.WriteFile("image.jpg", imgBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func NewService(userService *user.Service) *Service {
	return &Service{userService: userService}
}
