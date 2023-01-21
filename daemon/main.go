package main

import (
	configContract "github.com/codestates/WBA-BC-Project-02/common/config/dev"
	"github.com/codestates/WBA-BC-Project-02/daemon/subscribe"
	"github.com/codestates/WBA-BC-Project-02/daemon/utils"
	wasConfig "github.com/codestates/WBA-BC-Project-02/was/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	cPath := "./common/config/dev/config.toml"
	contractConfig := wasConfig.NewConfig(cPath, &configContract.Contract{})

	client, err := ethclient.Dial("wss://ws.test.wemix.com")
	utils.ErrorHandler(err)

	creditCh := make(chan bool, 1)
	dracoCh := make(chan bool, 1)
	dexCh := make(chan bool, 1)
	tigCh := make(chan bool, 1)

	go subscribe.CreditListener(contractConfig.CreditAddr, client, creditCh)
	go subscribe.DracoListener(contractConfig.DracoAddr, client, dracoCh)
	go subscribe.DexListener(contractConfig.DexAddr, client, dexCh)
	go subscribe.TigListener(contractConfig.TigAddr, client, tigCh)

	for {
		select {
		case <-creditCh:
			go subscribe.CreditListener(contractConfig.CreditAddr, client, creditCh)
		case <-dracoCh:
			go subscribe.DracoListener(contractConfig.DracoAddr, client, dracoCh)
		case <-dexCh:
			go subscribe.DexListener(contractConfig.DexAddr, client, dexCh)
		case <-tigCh:
			go subscribe.TigListener(contractConfig.TigAddr, client, tigCh)
		}
	}
}
