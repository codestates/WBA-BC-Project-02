package main

import (
	"log"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/codestates/WBA-BC-Project-02/contracts/signerdaemon/listener"
)

func main() {
	multisigAddr := "0x6f574c6325B3cB3F86E8bfA5f306310D63dD217d"

	client, err := ethclient.Dial("wss://ws.test.wemix.com")
	if err != nil {
		log.Fatal(err)
	}	

	signerCh := make(chan bool, 1)
	listener.MultisigListener(multisigAddr , client, signerCh)

	for {
		select {
		case <- signerCh:
			go listener.MultisigListener(multisigAddr , client, signerCh)
		}
	}
}