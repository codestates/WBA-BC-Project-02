package request

type ExchangeContract struct {
	Password string `json:"password" binding:"required"`
	From     string `json:"from" binding:"required,eq=draco|eq=tig|eq=credit"`
	To       string `json:"to" binding:"required,eq=draco|eq=tig|eq=credit|eq=wemix"`
	Amount   int    `json:"amount" binding:"required,min=1"`
}
