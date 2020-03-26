package domain

type ApplicationData struct {
	AccessCode      string
	ClientId        string
	ClientSecret    string
	RefreshToken    string
	TokenExpiration string
}

func NewApplicationData(accessCode string, clientId string, clientSecret string) *ApplicationData {
	return &ApplicationData{
		AccessCode:      accessCode,
		ClientId:        clientId,
		ClientSecret:    clientSecret,
		RefreshToken:    "",
		TokenExpiration: "",
	}
}
