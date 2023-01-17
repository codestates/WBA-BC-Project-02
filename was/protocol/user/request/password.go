package request

type Password struct {
	Password string `json:"password,omitempty" binding:"required,min=13"`
}
