package convertor

import "github.com/codestates/WBA-BC-Project-02/common/enum"

func DefaultAmount(amount string) string {
	if amount != enum.BlankSTR {
		return amount
	}
	return "0"
}
