package dom

type Transaction struct {
	TxHash          string `bson:"tx_hash,omitempty" json:"tx_hash"`
	ContractAddress string `bson:"contract_address,omitempty" json:"contract_address"`
	TxType          string `bson:"tx_type,omitempty" json:"tx_type"`
	From            string `bson:"from,omitempty" json:"from"`
	To              string `bson:"to,omitempty" json:"to"`
	Amount          string `bson:"amount,omitempty" json:"amount"`
	CreatedAt       string `bson:"created_at,omitempty" json:"created_at"`
}
