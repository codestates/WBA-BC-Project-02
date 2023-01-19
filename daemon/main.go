package main

import (
	configContract "github.com/codestates/WBA-BC-Project-02/common/config/dev"
	conf "github.com/codestates/WBA-BC-Project-02/daemon/config"
	"github.com/codestates/WBA-BC-Project-02/daemon/subscribe"
	"github.com/codestates/WBA-BC-Project-02/daemon/utils"
	wasConfig "github.com/codestates/WBA-BC-Project-02/was/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	cf := conf.GetConfig("./daemon/config/config.toml")
	cPath := "./common/config/dev/config.toml"
	contractConfig := wasConfig.NewConfig(cPath, &configContract.Contract{})

	client, err := ethclient.Dial(cf.Network.URL)
	utils.ErrorHandler(err)

	creditCh := make(chan bool, 1)
	dracoCh := make(chan bool, 1)
	dexCh := make(chan bool, 1)
	tigCh := make(chan bool, 1)

	subscribe.CreditListener(contractConfig.CreditAddr, client, creditCh)
	subscribe.DracoListener(contractConfig.DracoAddr, client, dracoCh)
	subscribe.DexListener(contractConfig.DexAddr, client, dexCh)
	subscribe.TigListener(contractConfig.TigAddr, client, tigCh)

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
