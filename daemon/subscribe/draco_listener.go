package subscribe

import (
	"context"
	"fmt"
	"log"
	"strings"

	draco "github.com/codestates/WBA-BC-Project-02/contracts/draco"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DracoListener(address string, client *ethclient.Client, ch chan<- bool) {
	contractAddr := common.HexToAddress(address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractABI, err := abi.JSON(strings.NewReader(string(draco.ContractsABI)))
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

			fmt.Println(event)

			if event.Name == "Mint" {
				fmt.Println(vLog.TxHash.Hex())
				fmt.Println(vLog.Address)

				result, err := contractABI.Unpack(event.Name, vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(result[0])
				// ...
			}
		}
	}
}
