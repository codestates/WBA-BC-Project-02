package request

type ExchangeContract struct {
	Password string `json:"password" binding:"required"`
	From     string `json:"from" binding:"required"`
	To       string `json:"to" binding:"required"`
}
