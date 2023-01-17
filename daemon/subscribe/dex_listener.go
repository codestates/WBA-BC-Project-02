package subscribe

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/codestates/WBA-BC-Project-02/daemon/model"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/bson"
)

type DexPoolData struct {
	PoolToken  string
	PoolCredit string
}

func DexListener(address string, client *ethclient.Client, ch chan<- bool) {
	contractAddr := common.HexToAddress(address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractABI, err := abi.JSON(strings.NewReader(string("dex.ContractABI")))
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

			if event.Name == "Ratio" {
				fmt.Println(event.Name)
				result, err := contractABI.Unpack(event.Name, vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				// event parameter
				tokenName := fmt.Sprintf("%v", result[0])
				tokenAmount := fmt.Sprintf("%v", result[1])
				creditAmount := fmt.Sprintf("%v", result[2])

				transactionHash := vLog.TxHash.Hex()
				poolData := DexPoolData{
					PoolToken:  tokenAmount,
					PoolCredit: creditAmount,
				}

				r, err := model.NewModel()
				if err != nil {
					log.Fatal(err)
				}

				// draco일 경우
				if tokenName == "Draco" {
					// dex contract update
					filter := bson.D{{Key: "contract_address", Value: address}}
					update := bson.M{
						"$set":  bson.M{"draco_pool_token": poolData.PoolToken, "draco_pool_credit": poolData.PoolCredit},
						"$push": bson.M{"transaction_hashs": transactionHash},
					}

					result, err := r.ColDexContract.UpdateOne(context.TODO(), filter, update)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Printf("dex contract update: %v\n", result.ModifiedCount)

					// tig일 경우
				} else if tokenName == "Tig" {
					// dex contract update
					filter := bson.D{{Key: "contract_address", Value: address}}
					update := bson.M{
						"$set":  bson.M{"tig_pool_token": poolData.PoolToken, "tig_pool_credit": poolData.PoolCredit},
						"$push": bson.M{"transaction_hashs": transactionHash},
					}

					result, err := r.ColDexContract.UpdateOne(context.TODO(), filter, update)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Printf("dex contract update: %v\n", result.ModifiedCount)
				}
			}
		}
	}
}
