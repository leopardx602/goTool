package model

type GoogleToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"` // 3600
	Scope     string `json:"scope"`
	TokenType string `json:"token_type"` // Bearer
	IDToken   string `json:"id_token"`
}

type GoogleUserInfo struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	Locale     string `json:"locale"` // zh-TW
}
