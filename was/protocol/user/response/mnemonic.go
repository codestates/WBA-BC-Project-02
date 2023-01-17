package response

type Mnemonic struct {
	Mnemonic []string `json:"mnemonic"`
	Token    Token    `json:"token"`
}
