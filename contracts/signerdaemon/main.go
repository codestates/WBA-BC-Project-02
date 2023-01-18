package main

import (
	"fmt"
	"log"
	"os"
	
	// "os/signal"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/codestates/WBA-BC-Project-02/common/enum"
	"github.com/codestates/WBA-BC-Project-02/common/ciper"
	"github.com/codestates/WBA-BC-Project-02/contracts/signerdaemon/listener"
	configContract "github.com/codestates/WBA-BC-Project-02/common/config/dev"
	"github.com/bwmarrin/discordgo"
	"github.com/naoina/toml"
	"github.com/joho/godotenv"
	"crypto/aes"
)

// var (
// 	Token string
// )

// func init() {
// 	flag.StringVar(&Token, "t", "", "MTA2NTE2OTY5NTM0MDI0NTA0NA.GILwNt.2hpQ2b7_w0J_VQldtTZGUnV5IENLIfYk3-80Rg")
// }

func NewConfig[T any](path string, t T)T {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal("start app... does not exists config file in ", path)
	}

	if err := toml.NewDecoder(file).Decode(t); err != nil {
		log.Fatal("start app... toml decode, fail ::", err.Error())
	}
	return t
}

func checkDevModAndLoadEnv() {
    if err := godotenv.Load(enum.DevModeENVFilePath); err != nil {
        log.Fatal("Error loading .env file ::", err.Error())
    }
}


func main() {
	var err error
	t := &configContract.Contract{}

	ContractConfig := NewConfig("common/config/dev/config.toml", t)

	checkDevModAndLoadEnv()
	key := os.Getenv("WEMIXON_DEV_KEY")

	block, err := aes.NewCipher([]byte(key))
    if err != nil {
        log.Fatal(err.Error())
    }

	fk, err := ciper.AESDecrypt(block, ContractConfig.ServerPrivateKey)
	if err != nil {
        log.Fatal(err.Error())
    }
    ContractConfig.ServerPrivateKey = fk

    ms, err := ciper.AESDecrypt(block, ContractConfig.MultiSigAddr)
    if err != nil {
        log.Fatal(err.Error())
    }
    ContractConfig.MultiSigAddr = ms

    da, err := ciper.AESDecrypt(block, ContractConfig.DexAddr)
    if err != nil {
        log.Fatal(err.Error())
    }
    ContractConfig.DexAddr = da

	ct, err := ciper.AESDecrypt(block, ContractConfig.ChannelToken)
	if err != nil {
		log.Fatal(err.Error())
	}
	ContractConfig.ChannelToken = ct

	spk, err := ciper.AESDecrypt(block, ContractConfig.SecondPrivateKey)
	if err != nil {
		log.Fatal(err.Error())
	}
	ContractConfig.SecondPrivateKey = spk

	btk, err := ciper.AESDecrypt(block, ContractConfig.BotToken)
	if err != nil {
		log.Fatal(err.Error())
	}
	ContractConfig.BotToken = btk
	
	discord, err := discordgo.New("Bot " + ContractConfig.BotToken)
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
	listener.MultisigListener(ContractConfig, client, signerCh, discord)

	for {
		select {
		case <- signerCh:
			go listener.MultisigListener(ContractConfig, client, signerCh, discord)
		}
	}
}
