package request

type MintingContract struct {
	Password            string `json:"password" binding:"required"`
	MintingName         string `json:"minting_name" binding:"required,eq=draco|eq=tig"`
	BlackIronBurnAmount int    `json:"burn_amount" binding:"required,min=1"`
}
