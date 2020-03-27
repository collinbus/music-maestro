package domain

type User struct {
	Id        string
	Name      string
	Urls      *Urls
	Image     *Image
	Followers int
}

type Urls struct {
	Internal string
	External string
	Uri      string
}

func NewUser(id string, name string, urls *Urls, image *Image, followers int) *User {
	return &User{Id: id, Name: name, Urls: urls, Image: image, Followers: followers}
}

func NewUrls(internal string, external string, uri string) *Urls {
	return &Urls{Internal: internal, External: external, Uri: uri}
}
