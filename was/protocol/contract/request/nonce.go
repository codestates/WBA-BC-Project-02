package request

type Nonce struct {
	Value int `json:"value" binding:"required"`
}
