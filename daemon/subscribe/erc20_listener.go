package subscribe

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
	"github.com/codestates/WBA-BC-Project-02/common/model/entity/dom"
	draco "github.com/codestates/WBA-BC-Project-02/contracts/draco"
	"github.com/codestates/WBA-BC-Project-02/daemon/model"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ERC20Listener(address string, client *ethclient.Client, ch chan<- bool) {
	erc20AddressMap := map[string]string{
		"dracoAddr": "0x3E46f05A6E8eFb386dE21249af792735AEBec19e",
		"tigAddr":   "",
	}

	contractAddr := common.HexToAddress(address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	var contractsABI string
	if address == erc20AddressMap["dracoAddr"] {
		contractsABI = draco.ContractsABI
	}

	contractABI, err := abi.JSON(strings.NewReader(contractsABI))
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

				// int64의 최대 범위는 9223372036854775807로 9.223... * 10 ** 18이 max
				// TODO 해결 방법 모색
				updateAmount := new(big.Int)
				iUserAmount, _ := strconv.Atoi(user.DracoAmount)
				biUserAmount := big.NewInt(int64(iUserAmount))
				iAmount, _ := strconv.Atoi(amount)
				biAmount := big.NewInt(int64(iAmount))
				updateAmount.Add(biUserAmount, biAmount)

				fmt.Println(biUserAmount, biAmount, updateAmount)

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
		}
	}
}
