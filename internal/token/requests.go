package token

import "strings"

func NewApiTokenRequestBody(code string, clientId string, clientSecret string) *strings.Reader {
	request := "grant_type=authorization_code&code=" + code +
		"&client_id=" + clientId +
		"&client_secret=" + clientSecret +
		"&redirect_uri=https://www.google.com/"
	return strings.NewReader(request)
}

func NewRefreshTokenRequestBody(refreshToken string, clientId string, clientSecret string) *strings.Reader {
	request := "grant_type=refresh_token&refresh_token=" + refreshToken +
		"&client_id=" + clientId +
		"&client_secret=" + clientSecret
	return strings.NewReader(request)
}
