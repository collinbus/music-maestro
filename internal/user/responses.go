package user

type GetUserInfoResponse struct {
	Id           string `json:"id"`
	Name         string `json:"display_name"`
	ExternalUrls struct {
		SpotifyUserUrl string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Total int `json:"total"`
	} `json:"followers"`
	Link   string `json:"href"`
	Uri    string `json:"uri"`
	Images []struct {
		Url string `json:"url"`
	} `json:"images"`
}
