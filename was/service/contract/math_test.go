package contract

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestMath(t *testing.T) {
	tokenF, err := strconv.ParseFloat("139.3850193", 64)
	if err != nil {
		fmt.Println(err)
	}
	creditF, err := strconv.ParseFloat("11.492103", 64)
	if err != nil {
		fmt.Println(err)
	}

	price := math.Round((tokenF/creditF)*10000) / 10000
	fmt.Println(price)
	priceSTR := fmt.Sprintf("%f", price)
	fmt.Println(priceSTR)

	var u uint64 = 31294781297498127398127983192
	uSTR := fmt.Sprintf("%v", u)
	fmt.Println(uSTR)
}
