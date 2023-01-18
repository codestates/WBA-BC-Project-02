package subscribe

import (
	"context"
	"fmt"
	"strings"

	"github.com/codestates/WBA-BC-Project-02/daemon/model"
	"github.com/codestates/WBA-BC-Project-02/daemon/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	utils.ErrorHandler(err)

	contractABI, err := abi.JSON(strings.NewReader(string("dex.ContractABI")))
	utils.ErrorHandler(err)

	for {
		select {
		case <-sub.Err():
			ch <- true
			return
		case vLog := <-logs:
			event, err := contractABI.EventByID(vLog.Topics[0])
			utils.ErrorHandler(err)

			if event.Name == "Ratio" {
				fmt.Println(event.Name)
				result, err := contractABI.Unpack(event.Name, vLog.Data)
				utils.ErrorHandler(err)

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
				utils.ErrorHandler(err)

				filter := bson.D{{Key: "contract_address", Value: address}}
				var update primitive.M

				// Draco일 경우
				if tokenName == "Draco" {
					update = bson.M{
						"$set":  bson.M{"draco_pool_token": poolData.PoolToken, "draco_pool_credit": poolData.PoolCredit},
						"$push": bson.M{"transaction_hashs": transactionHash},
					}
					// Tig일 경우
				} else if tokenName == "Tig" {
					update = bson.M{
						"$set":  bson.M{"tig_pool_token": poolData.PoolToken, "tig_pool_credit": poolData.PoolCredit},
						"$push": bson.M{"transaction_hashs": transactionHash},
					}
				}

				updateResult, err := r.ColDexContract.UpdateOne(context.TODO(), filter, update)
				utils.ErrorHandler(err)

				fmt.Printf("dex contract update: %v\n", updateResult.ModifiedCount)
			}
		}
	}
}
