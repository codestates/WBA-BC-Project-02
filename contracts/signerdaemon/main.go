package main

import (
	"log"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/codestates/WBA-BC-Project-02/contracts/signerdaemon/listener"
	"github.com/codestates/WBA-BC-Project-02/contracts/config"
)

func main() {
	conf := config.GetConfig("config/config.toml")
	
	firstPk := conf.Info.FirstPk
	secondPk := conf.Info.SecondPk
	multisigAddr := conf.Info.MultisigAddr

	client, err := ethclient.Dial("wss://ws.test.wemix.com")
	if err != nil {
		log.Fatal(err)
	}	

	signerCh := make(chan bool, 1)
	listener.MultisigListener(firstPk, secondPk, multisigAddr , client, signerCh)

	for {
		select {
		case <- signerCh:
			go listener.MultisigListener(firstPk, secondPk, multisigAddr , client, signerCh)
		}
	}
}