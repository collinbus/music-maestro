package domain

type User struct {
	Id        string
	Name      string
	Urls      *Urls
	ImageUrl  string
	Followers int
}

type Urls struct {
	Internal string
	External string
	Uri      string
}

func NewUser(id string, name string, urls *Urls, imageUrl string, followers int) *User {
	return &User{Id: id, Name: name, Urls: urls, ImageUrl: imageUrl, Followers: followers}
}

func NewUrls(internal string, external string, uri string) *Urls {
	return &Urls{Internal: internal, External: external, Uri: uri}
}
