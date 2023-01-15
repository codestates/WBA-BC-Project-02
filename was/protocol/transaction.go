package protocol

import "github.com/codestates/WBA-BC-Project-02/common/model/entity/dom"

type Transaction struct {
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
	TxType          string `json:"tx_type"`
	From            string `json:"from"`
	To              string `json:"to"`
	Amount          string `json:"amount"`
	CreatedAt       string `json:"created_at"`
}

func FromTransactionDoms(txs []*dom.Transaction) []*Transaction {
	transactions := make([]*Transaction, len(txs))
	for _, tx := range txs {
		t := FromTransactionDom(tx)
		transactions = append(transactions, t)
	}
	return transactions
}

func FromTransactionDom(tx *dom.Transaction) *Transaction {
	return &Transaction{
		TxHash:          tx.TxHash,
		ContractAddress: tx.ContractAddress,
		TxType:          tx.TxType,
		From:            tx.From,
		To:              tx.To,
		Amount:          tx.Amount,
		CreatedAt:       tx.CreatedAt,
	}
}
