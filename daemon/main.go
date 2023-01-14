package main

import (
	conf "lecture/WBA-BC-Project-02/daemon/config"
	"lecture/WBA-BC-Project-02/daemon/model"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	cf := conf.GetConfig("./config/config.toml")

	_, err := model.NewModel(cf.DB.Host)
	if err != nil {
		log.Fatal(err)
	}

	_, err = ethclient.Dial(cf.Network.URL)
	if err != nil {
		log.Fatal(err)
	}
}
