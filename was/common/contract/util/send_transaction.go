package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/codestates/WBA-BC-Project-02/was/config"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/codestates/WBA-BC-Project-02/contracts/multisig"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SendTransaction(client *ethclient.Client, nonce int64, data []byte) error {
	firstPk, err := crypto.HexToECDSA(config.ContractConfig.ServerPrivateKey) // 서버의 privateKey -> config에서 가져올 것
	if err != nil {
		fmt.Println("first Pk err")
		log.Fatal(err)
	}
<<<<<<<< HEAD:was/common/contract/util/send_transaction.go
	multisigAddr := common.HexToAddress(config.ContractConfig.MultiSigAddr) // multisig contract의 주소 -> config에서 가져올 것
	dexContractAddr := common.HexToAddress(config.ContractConfig.DexAddr)   // dex 컨트랙트의 주소 -> config에서 가져올 것
========
	multisigAddr := common.HexToAddress("0x6f574c6325B3cB3F86E8bfA5f306310D63dD217d") // multisig contract의 주소 -> config에서 가져올 것
	dexContractAddr := common.HexToAddress("0x8835f3bcEB451Ea73f0Ef3891f47bc66c7D24306") // dex 컨트랙트의 주소 -> config에서 가져올 것
>>>>>>>> dev:temp/util/sendTransaction.go

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

	fmt.Println("txIdx: ", txIdx)
	fmt.Println("nonce: ", nonce)

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
	fmt.Println("Sending...")
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
	txIdxHex := datas["logs"].([]interface{})[0].(map[string]interface{})["data"].(string)[2:66]
	i, _ := strconv.ParseInt(txIdxHex, 16, 64)
	return i, nil
}
