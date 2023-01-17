package main

import (
	"log"

	"github.com/codestates/WBA-BC-Project-02/daemon/subscribe"

	conf "github.com/codestates/WBA-BC-Project-02/daemon/config"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	CreditAddr = "0x279dF2e55fA16334033881525c443Ef4A3f341B9"
	DracoAddr  = "0xF0Cd034Ce5e613a80b9f7930CD076b2De8A47241"
	DexAddr    = ""
	TigAddr    = ""
)

func main() {
	var cf = conf.GetConfig("./config/config.toml")

	client, err := ethclient.Dial(cf.Network.URL)
	if err != nil {
		log.Fatal(err)
	}

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
