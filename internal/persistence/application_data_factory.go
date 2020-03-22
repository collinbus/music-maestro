package persistence

type ApplicationData struct {
	accessCode   string
	clientId     string
	clientSecret string
}

func NewApplicationData(accessCode string, clientId string, clientSecret string) *ApplicationData {
	return &ApplicationData{accessCode: accessCode, clientId: clientId, clientSecret: clientSecret}
}
