package request

type Nonce struct {
	Value int `json:"value" form:"value" binding:"required"`
}
