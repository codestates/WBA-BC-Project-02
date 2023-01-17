package request

type MintingContract struct {
	Password    string `json:"password" binding:"required"`
	MintingName string `json:"minting_name" binding:"required"`
}
