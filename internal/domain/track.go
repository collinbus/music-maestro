package domain

import "time"

type Track struct {
	Id         string
	Name       string
	Popularity int
	Preview    string
	AddedAt    *time.Time
	Urls       *Urls
	Artists    []Artist
	Album      *Album
}

func NewTrack(id string, name string) *Track {
	return &Track{Id: id, Name: name}
}
