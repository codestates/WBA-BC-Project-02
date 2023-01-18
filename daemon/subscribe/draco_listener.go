package subscribe

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
	"github.com/codestates/WBA-BC-Project-02/common/model/entity/dom"
	draco "github.com/codestates/WBA-BC-Project-02/contracts/draco"
	"github.com/codestates/WBA-BC-Project-02/daemon/model"
	"github.com/codestates/WBA-BC-Project-02/daemon/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

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
	utils.ErrorHandler(err)

	contractABI, err := abi.JSON(strings.NewReader(draco.DracoABI))
	utils.ErrorHandler(err)

	for {
		select {
		case <-sub.Err():
			ch <- true
			return
		case vLog := <-logs:
			event, err := contractABI.EventByID(vLog.Topics[0])
			utils.ErrorHandler(err)

			if event.Name == "Mint" {
				fmt.Println(event.Name)
				result, err := contractABI.Unpack(event.Name, vLog.Data)
				utils.ErrorHandler(err)

				// event parameter 값
				to := fmt.Sprintf("%v", result[0])
				amount := fmt.Sprintf("%v", result[1])

				// transaction 값 생성
				transaction := dom.Transaction{
					TxHash:          vLog.TxHash.Hex(),
					ContractAddress: vLog.Address.String(),
					TxType:          event.Name,
					From:            vLog.Address.String(),
					To:              to,
					Amount:          amount,
					CreatedAt:       time.Now().Format("2006-01-02 15:04:05"),
				}

				r, err := model.NewModel()
				utils.ErrorHandler(err)

				// user 조회
				var user entity.User
				userFilter := bson.D{{Key: "address", Value: to}}
				r.ColUser.FindOne(context.TODO(), userFilter).Decode(&user)

				// amount 계산
				updateAmount := new(big.Int)
				biUserAmount := new(big.Int)
				userAmount, _ := biUserAmount.SetString(user.DracoAmount, 10)
				biAmount := new(big.Int)
				mintAmount, _ := biAmount.SetString(amount, 10)

				updateAmount.Add(userAmount, mintAmount)

				fmt.Printf("user amount: %s + mint amount: %s = total amount: %s\n", userAmount, mintAmount, updateAmount)

				// user update
				userUpdate := bson.M{
					"$set":  bson.M{"draco_amount": updateAmount.String()},
					"$push": bson.M{"transactions": transaction},
				}

				userUpdateResult, err := r.ColUser.UpdateOne(context.TODO(), userFilter, userUpdate)
				utils.ErrorHandler(err)

				// contract update
				contractFilter := bson.D{{Key: "contract_address", Value: transaction.ContractAddress}}
				contractUpdate := bson.M{
					"$push": bson.M{"transactions": transaction},
				}

				contractUpdateResult, err := r.ColContract.UpdateOne(context.TODO(), contractFilter, contractUpdate)

				fmt.Printf("user update: %v, contract update: %v\n", userUpdateResult.ModifiedCount, contractUpdateResult.ModifiedCount)
			}

			if event.Name == "CustomTransfer" {
				fmt.Println(event.Name)

				result, err := contractABI.Unpack(event.Name, vLog.Data)
				utils.ErrorHandler(err)

				// event prameter
				from := fmt.Sprintf("%v", result[0])
				to := fmt.Sprintf("%v", result[1])
				amount := fmt.Sprintf("%v", result[2])

				// transaction
				transaction := dom.Transaction{
					TxHash:          vLog.TxHash.Hex(),
					ContractAddress: vLog.Address.String(),
					TxType:          event.Name,
					From:            from,
					To:              to,
					Amount:          amount,
					CreatedAt:       "",
				}

				r, err := model.NewModel()
				utils.ErrorHandler(err)

				// from user 조회
				var fromUser entity.User
				fromUserFilter := bson.D{{Key: "address", Value: from}}
				r.ColUser.FindOne(context.TODO(), fromUserFilter).Decode(&fromUser)

				// to user 조회
				var toUser entity.User
				toUserFilter := bson.D{{Key: "address", Value: to}}
				r.ColUser.FindOne(context.TODO(), toUserFilter).Decode(&toUser)

				zeroObjectId, _ := primitive.ObjectIDFromHex("000000000000000000000000")

				// amount 계산
				updateAddAmount := new(big.Int)
				updateSubAmount := new(big.Int)

				biUserAmount := new(big.Int)
				userAmount, _ := biUserAmount.SetString(fromUser.DracoAmount, 10)
				biAmount := new(big.Int)
				trasnferAmount, _ := biAmount.SetString(amount, 10)

				updateAddAmount.Add(userAmount, trasnferAmount)
				updateSubAmount.Sub(userAmount, trasnferAmount)

				// from일 경우
				if zeroObjectId != fromUser.ID {
					fmt.Printf("user amount: %s - trasnfer amount: %s = total amount: %s\n", userAmount, trasnferAmount, updateSubAmount)

					// user update
					update := bson.M{
						"$set":  bson.M{"draco_amount": updateSubAmount.String()},
						"$push": bson.M{"transactions": transaction},
					}

					result, err := r.ColUser.UpdateOne(context.TODO(), fromUserFilter, update)
					utils.ErrorHandler(err)

					fmt.Printf("from user udpate: %v\n", result.ModifiedCount)
				}

				// to일 경우
				if zeroObjectId != toUser.ID {
					fmt.Printf("user amount: %s + trasnfer amount: %s = total amount: %s\n", userAmount, trasnferAmount, updateAddAmount)

					// user update
					update := bson.M{
						"$set":  bson.M{"draco_amount": updateAddAmount.String()},
						"$push": bson.M{"transactions": transaction},
					}

					result, err := r.ColUser.UpdateOne(context.TODO(), toUserFilter, update)
					utils.ErrorHandler(err)

					fmt.Printf("to user udpate: %v\n", result.ModifiedCount)
				}

				// contract update
				contractFilter := bson.D{{Key: "contract_address", Value: transaction.ContractAddress}}
				contractUpdate := bson.M{
					"$push": bson.M{"transactions": transaction},
				}

				contractUpdateResult, err := r.ColContract.UpdateOne(context.TODO(), contractFilter, contractUpdate)

				fmt.Printf("contract update: %v\n", contractUpdateResult.ModifiedCount)
			}
		}
	}
}
