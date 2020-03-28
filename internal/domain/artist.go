package domain

type Artist struct {
	Id   string
	Name string
	Urls *Urls
}

func NewArtist(id string, name string, urls *Urls) *Artist {
	return &Artist{Id: id, Name: name, Urls: urls}
}
