package network

import "strings"

func NewApiTokenRequestBody(code string, clientId string, clientSecret string) *strings.Reader {
	request := "grant_type=authorization_code&code=" + code +
		"&client_id=" + clientId +
		"&client_secret=" + clientSecret +
		"&redirect_uri=https://www.google.com/"
	return strings.NewReader(request)
}
