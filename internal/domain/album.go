package domain

import "time"

type Album struct {
	Id             string
	Name           string
	NumberOfTracks int
	ReleaseDate    *time.Time
	AlbumArt       []AlbumArt
	Urls           *Urls
}

func NewAlbum(id string, name string, numberOfTracks int) *Album {
	return &Album{Id: id, Name: name, NumberOfTracks: numberOfTracks}
}

type AlbumArt struct {
	Width     int
	Height    int
	ImageUrl  string
	ImageData string
}

func NewAlbumArt(width int, height int, imageUrl string) *AlbumArt {
	return &AlbumArt{Width: width, Height: height, ImageUrl: imageUrl}
}
