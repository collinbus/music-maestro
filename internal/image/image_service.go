package image

import (
	"encoding/base64"
	"musicMaestro/internal/network"
)

type Service struct{}

func (Service) DownloadImage(url string) string {
	imgBytes := network.DownloadImage(url)
	base64String := convertToBase64(imgBytes)
	return base64String
}

func convertToBase64(data []byte) string {
	base64String := base64.RawStdEncoding.EncodeToString(data)
	return base64String
}

func NewService() *Service {
	return &Service{}
}
