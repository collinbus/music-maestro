package image

import (
	"encoding/base64"
	"musicMaestro/internal/network"
	"musicMaestro/internal/user"
)

type Service struct {
	userService *user.Service
}

func (service *Service) DownloadUserImage() {
	fetchedUser := service.userService.FetchUser()
	imgBytes := network.DownloadImage(fetchedUser.Image.Url)
	base64String := convertToBase64(imgBytes)

	fetchedUser.Image.Data = base64String
	service.userService.UpdateUser(fetchedUser)
}

func convertToBase64(data []byte) string {
	base64String := base64.RawStdEncoding.EncodeToString(data)
	return base64String
}

func NewService(userService *user.Service) *Service {
	return &Service{userService: userService}
}
