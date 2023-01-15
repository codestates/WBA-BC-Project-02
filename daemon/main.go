package main

import (
	"log"

	conf "github.com/codestates/WBA-BC-Project-02/daemon/config"
	"github.com/codestates/WBA-BC-Project-02/daemon/subscribe"

	"github.com/codestates/WBA-BC-Project-02/daemon/model"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	creditAddr = "1234"
	dracoAddr  = "0xEF04548b03453C0b6070afe55392132bDDeb93Dc"
)

func main() {
	cf := conf.GetConfig("./config/config.toml")

	_, err := model.NewModel(cf.DB.Host)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(cf.Network.URL)
	if err != nil {
		log.Fatal(err)
	}

	creditCh := make(chan bool, 1)
	dracoCh := make(chan bool, 1)

	subscribe.CreditListener(creditAddr, client, creditCh)
	subscribe.DracoListener(dracoAddr, client, dracoCh)

	for {
		select {
		case <-creditCh:
			go subscribe.CreditListener(creditAddr, client, creditCh)
		case <-dracoCh:
			go subscribe.DracoListener(dracoAddr, client, dracoCh)
		}
	}
}
