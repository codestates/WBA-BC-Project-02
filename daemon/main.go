package main

import (
	"github.com/codestates/WBA-BC-Project-02/daemon/subscribe"
	"github.com/codestates/WBA-BC-Project-02/daemon/utils"

	conf "github.com/codestates/WBA-BC-Project-02/daemon/config"
	"github.com/ethereum/go-ethereum/ethclient"
)

// test용
var (
	CreditAddr = "0x279dF2e55fA16334033881525c443Ef4A3f341B9"
	DracoAddr  = "0x7D2397657A02f2AB0371f3B695F9041Cbca253De"
	DexAddr    = ""
	TigAddr    = ""
)

func main() {
	var cf = conf.GetConfig("./config/config.toml")

	client, err := ethclient.Dial(cf.Network.URL)
	utils.ErrorHandler(err)

	creditCh := make(chan bool, 1)
	dracoCh := make(chan bool, 1)
	dexCh := make(chan bool, 1)
	tigCh := make(chan bool, 1)

	subscribe.CreditListener(CreditAddr, client, creditCh)
	subscribe.DracoListener(DracoAddr, client, dracoCh)
	subscribe.DexListener(DexAddr, client, dexCh)
	subscribe.TigListener(TigAddr, client, tigCh)

	for {
		select {
		case <-creditCh:
			go subscribe.CreditListener(CreditAddr, client, creditCh)
		case <-dracoCh:
			go subscribe.DracoListener(DracoAddr, client, dracoCh)
		case <-dexCh:
			go subscribe.DexListener(DexAddr, client, dexCh)
		case <-tigCh:
			go subscribe.TigListener(TigAddr, client, tigCh)
		}
	}
}
