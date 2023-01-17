package request

type ContractName struct {
	Name string `json:"name" binding:"required,eq=credit|eq=draco|eq=tig"`
}
