package domain

type User struct {
	Id        string
	Name      string
	Urls      *Urls
	Image     *Image
	Followers int
}

func NewUser(id string, name string, urls *Urls, image *Image, followers int) *User {
	return &User{Id: id, Name: name, Urls: urls, Image: image, Followers: followers}
}
