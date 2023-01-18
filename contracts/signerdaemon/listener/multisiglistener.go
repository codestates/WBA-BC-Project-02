package listener

import (
	"math/big"
	"github.com/codestates/WBA-BC-Project-02/contracts/multisig"
	"github.com/codestates/WBA-BC-Project-02/contracts/signerdaemon/runner"
    "context"
    "log"
	"strings"
	"fmt"

    "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
)

func MultisigListener(firstPk string, secondPk string, address string, client *ethclient.Client, ch chan<- bool) {
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
				
				// 서버에 체크해서 맞다면, txIdx를 가지고 confirm, execute하는 로직을 이어가자
				

				// 서버에 체크해서 맞다고 확인받은 이후 컨펌 - execute 로직
				fmt.Println("txIdx: ", txIdx)
				fmt.Println("nonce: ", nonce)
				runner.RunTx(firstPk, secondPk, address, txIdx.(*big.Int))

				fmt.Println("success")
			}
		}
	}
}