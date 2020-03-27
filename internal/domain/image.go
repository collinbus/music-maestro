package domain

type Image struct {
	Url  string
	Data string
}

func NewImage(url string, data string) *Image {
	return &Image{Url: url, Data: data}
}
