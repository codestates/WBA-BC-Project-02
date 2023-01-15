package request

type Recovery struct {
	Mnemonic []string `json:"mnemonic" binding:"required"`
	Password string   `json:"new_password" binding:"required,min=13"`
}
