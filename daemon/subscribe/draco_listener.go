package subscribe

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
	"github.com/codestates/WBA-BC-Project-02/common/model/entity/dom"
	draco "github.com/codestates/WBA-BC-Project-02/contracts/draco"
	"github.com/codestates/WBA-BC-Project-02/daemon/model"
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
	if err != nil {
		log.Fatal(err)
	}

	contractABI, err := abi.JSON(strings.NewReader(draco.ContractsABI))
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

			if event.Name == "Mint" {
				fmt.Println(event.Name)
				result, err := contractABI.Unpack(event.Name, vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				to := fmt.Sprintf("%v", result[0])
				amount := fmt.Sprintf("%v", result[1])

				transaction := dom.Transaction{
					TxHash:          vLog.TxHash.Hex(),
					ContractAddress: vLog.Address.String(),
					TxType:          event.Name,
					From:            vLog.Address.String(),
					To:              to,
					Amount:          amount,
					CreatedAt:       "",
				}

				r, err := model.NewModel()
				if err != nil {
					log.Fatal(err)
				}

				var user entity.User
				userFilter := bson.D{{Key: "address", Value: to}}
				r.ColUser.FindOne(context.TODO(), userFilter).Decode(&user)

				updateAmount := new(big.Int)
				biUserAmount := new(big.Int)
				userAmount, _ := biUserAmount.SetString(user.DracoAmount, 10)
				biAmount := new(big.Int)
				Amount, _ := biAmount.SetString(amount, 10)

				updateAmount.Add(userAmount, Amount)

				fmt.Println(updateAmount, userAmount, Amount)

				userUpdate := bson.M{
					"$set":  bson.M{"draco_amount": updateAmount.String()},
					"$push": bson.M{"transactions": transaction},
				}

				userUpdateResult, err := r.ColUser.UpdateOne(context.TODO(), userFilter, userUpdate)
				if err != nil {
					log.Fatal(err)
				}

				contractFilter := bson.D{{Key: "contract_address", Value: transaction.ContractAddress}}
				contractUpdate := bson.M{
					"$push": bson.M{"transactions": transaction},
				}

				contractUpdateResult, err := r.ColContract.UpdateOne(context.TODO(), contractFilter, contractUpdate)

				fmt.Println("user", userUpdateResult.ModifiedCount, "contract", contractUpdateResult.ModifiedCount)
			}

			if event.Name == "CustomTransfer" {
				fmt.Println(event.Name)

				result, err := contractABI.Unpack(event.Name, vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				from := fmt.Sprintf("%v", result[0])
				to := fmt.Sprintf("%v", result[1])
				amount := fmt.Sprintf("%v", result[2])

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
				if err != nil {
					log.Fatal(err)
				}

				var fromUser entity.User
				fromUserFilter := bson.D{{Key: "address", Value: from}}
				r.ColUser.FindOne(context.TODO(), fromUserFilter).Decode(&fromUser)

				var toUser entity.User
				toUserFilter := bson.D{{Key: "address", Value: to}}
				r.ColUser.FindOne(context.TODO(), toUserFilter).Decode(&toUser)

				zeroObjectId, _ := primitive.ObjectIDFromHex("000000000000000000000000")

				if zeroObjectId != fromUser.ID {
					updateAmount := new(big.Int)
					biUserAmount := new(big.Int)
					userAmount, _ := biUserAmount.SetString(fromUser.CreditAmount, 10)
					biAmount := new(big.Int)
					Amount, _ := biAmount.SetString(amount, 10)

					updateAmount.Sub(userAmount, Amount)

					update := bson.M{
						"$set":  bson.M{"draco_amount": updateAmount.String()},
						"$push": bson.M{"transactions": transaction},
					}

					r.ColUser.UpdateOne(context.TODO(), fromUserFilter, update)
				}

				if zeroObjectId != toUser.ID {
					updateAmount := new(big.Int)
					biUserAmount := new(big.Int)
					userAmount, _ := biUserAmount.SetString(toUser.CreditAmount, 10)
					biAmount := new(big.Int)
					Amount, _ := biAmount.SetString(amount, 10)

					updateAmount.Add(userAmount, Amount)

					update := bson.M{
						"$set":  bson.M{"draco_amount": updateAmount.String()},
						"$push": bson.M{"transactions": transaction},
					}

					r.ColUser.UpdateOne(context.TODO(), toUserFilter, update)
				}
			}
		}
	}
}
