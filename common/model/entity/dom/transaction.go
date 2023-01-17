package dom

type Transaction struct {
	TxHash          string `bson:"tx_hash,omitempty"`
	ContractAddress string `bson:"contract_address,omitempty"`
	TxType          string `bson:"tx_type,omitempty"`
	From            string `bson:"from,omitempty"`
	To              string `bson:"to,omitempty"`
	Amount          string `bson:"amount,omitempty"`
	CreatedAt       string `bson:"created_at,omitempty"`
}
