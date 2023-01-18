package main

import (
	"log"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/codestates/WBA-BC-Project-02/contracts/signerdaemon"
)

func main() {
	multisigAddr := "0xFf4a73f42E69E70450E3eC97951a7b4639A866EB"

	client, err := ethclient.Dial("wss://ws.test.wemix.com")
	if err != nil {
		log.Fatal(err)
	}	

	signerCh := make(chan bool, 1)
	// dracoCh := make(chan bool, 1)
	// dexCh := make(chan bool, 1)
	signer.MultisigListener(multisigAddr , client, signerCh)

	for {
		select {
		case <- signerCh:
			go signer.MultisigListener(multisigAddr , client, signerCh)
		}
	}
}