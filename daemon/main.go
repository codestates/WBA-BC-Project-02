package main

import (
	"log"

	"github.com/codestates/WBA-BC-Project-02/daemon/subscribe"

	conf "github.com/codestates/WBA-BC-Project-02/daemon/config"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	CreditAddr = ""
	DracoAddr  = "0x3E46f05A6E8eFb386dE21249af792735AEBec19e"
)

func main() {
	var cf = conf.GetConfig("./config/config.toml")

	client, err := ethclient.Dial(cf.Network.URL)
	if err != nil {
		log.Fatal(err)
	}

	creditCh := make(chan bool, 1)
	dracoCh := make(chan bool, 1)

	subscribe.CreditListener(CreditAddr, client, creditCh)
	subscribe.DracoListener(DracoAddr, client, dracoCh)

	for {
		select {
		case <-creditCh:
			go subscribe.CreditListener(CreditAddr, client, creditCh)
		case <-dracoCh:
			go subscribe.DracoListener(DracoAddr, client, dracoCh)
		}
	}
}
