package response

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"-"`
}
