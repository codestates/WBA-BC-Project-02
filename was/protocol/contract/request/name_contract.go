package request

type ContractName struct {
	Name string `json:"name" form:"name" binding:"required,eq=credit|eq=draco|eq=tig"`
}
