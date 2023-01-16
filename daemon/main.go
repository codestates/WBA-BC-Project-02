package main

import (
	"log"

	conf "github.com/codestates/WBA-BC-Project-02/daemon/config"
	"github.com/codestates/WBA-BC-Project-02/daemon/subscribe"

	"github.com/ethereum/go-ethereum/ethclient"
)

var cf = conf.GetConfig("./config/config.toml")
var (
	URL        = cf.Network.URL
	DracoAddr  = cf.Addr.DracoAddr
	TigAddr    = cf.Addr.TigAddr
	CreditAddr = cf.Addr.CreditAddr
	DexAddr    = cf.Addr.DexAddr
)

func main() {
	client, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatal(err)
	}

	creditCh := make(chan bool, 1)
	dracoCh := make(chan bool, 1)

	subscribe.CreditListener(CreditAddr, client, creditCh)
	subscribe.ERC20Listener(DracoAddr, client, dracoCh)

	for {
		select {
		case <-creditCh:
			go subscribe.CreditListener(CreditAddr, client, creditCh)
		case <-dracoCh:
			go subscribe.ERC20Listener(DracoAddr, client, dracoCh)
		}
	}
}
