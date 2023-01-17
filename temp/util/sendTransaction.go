package util

import (
	"time"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"errors"

	"github.com/codestates/WBA-BC-Project-02/contracts/multisig"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)


func SendTransaction(client *ethclient.Client, nonce int64, data []byte) error {
	firstPk, err := crypto.HexToECDSA("68fe8eda422fbdcd2bc44ced268dfb3c42ec999c2e4f237dfc52af079cee2ddd") // 계정의 privateKey -> config에서 가져올 것
	if err != nil {
		fmt.Println("first Pk err")
		log.Fatal(err)
	}
	multisigAddr := common.HexToAddress("0x2603df2DCC9aC15E56346B752e81677CbeE4cF7D") // multisig contract의 주소 -> config에서 가져올 것
	dexContractAddr := common.HexToAddress("0xEC3b843a78D2d9430Acb73279Aed35fC0DEE427B") // dex 컨트랙트의 주소 -> config에서 가져올 것

	instance, err := multisig.NewMultisig(multisigAddr, client)
	if err != nil {
		fmt.Println("get instance Err")
		fmt.Println(err)
	}
	chainID, err := client.ChainID(context.TODO())
	if err != nil {
		fmt.Println("get chainId Err")
		fmt.Println(err)
	}
	auth1, err := bind.NewKeyedTransactorWithChainID(firstPk, chainID)
	if err != nil {
		fmt.Println("bind auth1 Err")
		fmt.Println(err)
	}

	txIdx, err := sendSubmitTransaction(client, instance, auth1, dexContractAddr, big.NewInt(0), data, big.NewInt(int64(nonce)))
	if err != nil {
		fmt.Println("get txIdx Err")
		log.Fatal(err)
	}

	count := 0
	for {
		if count == 20 {
			err := errors.New("count over")
			return err
		}
		result, _ := instance.GetTransaction(&bind.CallOpts{}, big.NewInt(txIdx))
		if result.Executed {
			break
		}
		count++
		fmt.Println(result.Executed)
		time.Sleep(time.Millisecond * 500)
		fmt.Println(count)
	}

	// 10초 안에 성공했을 시, user에게 성공 메시지를 반환하는 로직 필요
	fmt.Println("success")
	return nil
}

func sendSubmitTransaction(
	client *ethclient.Client,
	instance *multisig.Multisig,
	auth *bind.TransactOpts,
	dexContractAddr common.Address,
	value *big.Int,
	data []byte,
	nonce *big.Int,
) (int64, error) {
	fmt.Println("Sendging...")
	txSubmitTransaction, err := instance.SubmitTransaction(auth, dexContractAddr, value, data, nonce)
	if err != nil {
		fmt.Println("SendSubmitTransaction Err")
		return 0, err
	}

	var receipt *types.Receipt
	for {
		receipt, err = client.TransactionReceipt(context.Background(), txSubmitTransaction.Hash())
		if receipt != nil {
			break
		}
	}
	if err != nil {
		fmt.Println("get submitTrasaction Receipt Err")
		return 0, err
		// 현재 보낸 트랜잭션의 상태 확인. 1이면 정상
	} else if receipt.Status == 0 {
		// 에러처리
		fmt.Println("Status Err")
		return 0, err
	}

	// txIdx를 찾아서 return
	txData, err := receipt.MarshalJSON()
	if err != nil {
		fmt.Println("txData Err")
		return 0, err
	}
	var datas map[string]interface{}
	json.Unmarshal(txData, &datas)
	txIdxHex := datas["logs"].([]interface{})[0].(map[string]interface{})["topics"].([]interface{})[2]
	txIdxHex = txIdxHex.(string)[2:]
	i, _ := strconv.ParseInt(txIdxHex.(string), 16, 64)
	return i, nil
}
