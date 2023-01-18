package listener

import (
	"math/big"
	"github.com/codestates/WBA-BC-Project-02/contracts/multisig"
	"github.com/codestates/WBA-BC-Project-02/contracts/signerdaemon/txhandler"
    "context"
    "log"
	"strings"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"

    "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/bwmarrin/discordgo"
	configContract "github.com/codestates/WBA-BC-Project-02/common/config/dev"
)

func MultisigListener(
		configContract *configContract.Contract,
		client *ethclient.Client,
		ch chan<- bool,
		discord *discordgo.Session,
	) {
	firstPk := configContract.ServerPrivateKey
	secondPk := configContract.SecondPrivateKey
	address := configContract.MultiSigAddr
	channelId := configContract.ChannelToken
	

	contractAddress := common.HexToAddress(address)
	query := ethereum.FilterQuery {
		Addresses : []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractABI, err := abi.JSON(strings.NewReader(string(multisig.MultisigABI)))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-sub.Err():
			ch <- true
			return

		case vLog := <-logs:
			event, err := contractABI.EventByID(vLog.Topics[0])
			if err != nil {
				log.Fatal(err)
			}
			if event.Name == "SubmitTransaction" {
				result, err := contractABI.Unpack(event.Name, vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				txIdx := result[0]
				nonce := result[3] // 이후 논스로 gin서버에 체크하는 로직 필요
				url := "http://localhost:8080/contracts/nonce?value=" + nonce.(*big.Int).String()
				resp, err := http.Get(url)
				if err != nil {
					panic(err)
				}
				defer resp.Body.Close()
				// 결과 출력
				data, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					panic(err)
				}
				var unmarshaled map[string]interface{}
				json.Unmarshal(data, &unmarshaled)
				isValid := unmarshaled["data"].(map[string]interface{})["nonce"].(bool)
				fmt.Println("txIdx: ", txIdx)
				fmt.Println("nonce: ", nonce)
				// 서버에 보낸 Nonce값이 확인되었을 경우
				if isValid {
					fmt.Println("valid")	
					txhandler.RunTx(firstPk, secondPk, address, txIdx.(*big.Int))
					fmt.Println("success")
				// 서버에 보낸 Nonce값이 확인되지 않았을 경우
				} else if !isValid {
					fmt.Println("not Valid")
					msg := "Error Occured! txIdx: " + txIdx.(*big.Int).String() + " nonce: " + nonce.(*big.Int).String()
					discord.ChannelMessageSend(channelId, msg)
					txhandler.AbortTx(firstPk, secondPk, address, txIdx.(*big.Int))
				}
			}
		}
	}
}