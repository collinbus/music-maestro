package image

import (
	"encoding/base64"
	"fmt"
	"musicMaestro/internal/network"
)

func DownloadImage(url string) string {
	imgBytes := network.DownloadImage(url)

	if len(imgBytes) != 0 {
		fmt.Printf("Image at %s was successfully downloaded\n", url)
	}

	base64String := convertToBase64(imgBytes)
	return base64String
}

func convertToBase64(data []byte) string {
	base64String := base64.RawStdEncoding.EncodeToString(data)
	return base64String
}
