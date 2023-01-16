package main

import (
	"log"

	conf "github.com/codestates/WBA-BC-Project-02/daemon/config"
	"github.com/codestates/WBA-BC-Project-02/daemon/subscribe"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	creditAddr = ""
	dracoAddr  = "0x3E46f05A6E8eFb386dE21249af792735AEBec19e"
	tigAddr    = ""
)

func main() {
	cf := conf.GetConfig("./config/config.toml")

	client, err := ethclient.Dial(cf.Network.URL)
	if err != nil {
		log.Fatal(err)
	}

	creditCh := make(chan bool, 1)
	dracoCh := make(chan bool, 1)

	subscribe.CreditListener(creditAddr, client, creditCh)
	subscribe.ERC20Listener(dracoAddr, client, dracoCh)

	for {
		select {
		case <-creditCh:
			go subscribe.CreditListener(creditAddr, client, creditCh)
		case <-dracoCh:
			go subscribe.ERC20Listener(dracoAddr, client, dracoCh)
		}
	}
}
