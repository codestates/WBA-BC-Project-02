package request

type ReissueToken struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
