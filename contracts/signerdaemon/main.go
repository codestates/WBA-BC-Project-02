package main

import (
	"fmt"
	"log"
	// "os"
	// "os/signal"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/codestates/WBA-BC-Project-02/contracts/signerdaemon/listener"
	"github.com/codestates/WBA-BC-Project-02/contracts/config"
	"github.com/bwmarrin/discordgo"
)

// var (
// 	Token string
// )

// func init() {
// 	flag.StringVar(&Token, "t", "", "MTA2NTE2OTY5NTM0MDI0NTA0NA.GILwNt.2hpQ2b7_w0J_VQldtTZGUnV5IENLIfYk3-80Rg")
// }

func main() {
	conf := config.GetConfig("contracts/config/config.toml")

	channelId := conf.Info.ChannelID
	firstPk := conf.Info.FirstPk
	secondPk := conf.Info.SecondPk
	multisigAddr := conf.Info.MultisigAddr
	
	discord, err := discordgo.New("Bot " + conf.Info.BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	discord.Open()
	defer discord.Close()
	fmt.Println("Bot running...")

	client, err := ethclient.Dial("wss://ws.test.wemix.com")
	if err != nil {
		log.Fatal(err)
	}	

	signerCh := make(chan bool, 1)
	listener.MultisigListener(firstPk, secondPk, multisigAddr , client, signerCh, discord, channelId)

	for {
		select {
		case <- signerCh:
			go listener.MultisigListener(firstPk, secondPk, multisigAddr , client, signerCh, discord, channelId)
		}
	}
}
