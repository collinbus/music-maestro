package network

type ApiTokenResponseBody struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func NewApiTokenResponseBody() *ApiTokenResponseBody {
	return &ApiTokenResponseBody{}
}

type RefreshTokenResponseBody struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func NewRefreshTokenResponseBody() *RefreshTokenResponseBody {
	return &RefreshTokenResponseBody{}
}

type ErrorResponseBody struct {
	Error       string `json:"error"`
	Description string `json:"error_description"`
}

func NewErrorResponseBody() *ErrorResponseBody {
	return &ErrorResponseBody{}
}
