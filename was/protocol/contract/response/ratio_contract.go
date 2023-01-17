package response

import "github.com/codestates/WBA-BC-Project-02/was/protocol/contract"

type RatioContract struct {
	Price  string                   `json:"price"`
	Ratio  *contract.TokenAndCredit `json:"ratio"`
	Detail *contract.TokenAndCredit `json:"detail"`
}
