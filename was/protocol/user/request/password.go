package request

type Password struct {
	Password string `json:"password" binding:"required,min=13"`
}
