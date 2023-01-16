package main

import (
	"context"
	// "encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"log"
	"strconv"

	"github.com/codestates/WBA-BC-Project-02/contracts/multisig"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func GetDestinationfuncData(num int) []byte {
	// 실행할 함수의 signature 생성, methodID 생성
	transferFnSignature := []byte("callMe(uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	// fmt.Println(hexutil.Encode(methodID))

	// 아래는 인자로 넣을 data를 추가하는 과정
	// Destination address에 함수 인자값으로 넣을 uint값
	bigNum := big.NewInt(int64(num))
	paddedNum := common.LeftPadBytes(bigNum.Bytes(), 32)

	var pdata []byte
	pdata = append(pdata, methodID...)
	pdata = append(pdata, paddedNum...)
	// fmt.Println("0x"+hex.EncodeToString(pdata))
	return pdata
}


func main() {
	// server url과 연동
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	// 첫번째 계정의 privateKey
	firstPk, err := crypto.HexToECDSA("9734a6fab2886cdcc017c9441790681fdf99dbdbe16da22506d070a66f28755f")
	if err != nil {
		log.Fatal(err)
	}
	// 두번째 계정의 privateKey
	secondPk, err := crypto.HexToECDSA("30e9d002076fb2a3a10322738a62260a278073d6c42996b9f374032259cc8801")
	if err != nil {
		log.Fatal(err)
	}
	// multisig contract의 주소
	multisigAddr := common.HexToAddress("0xb80cE90781B7ed255695016fBcdd8389453ddA89")
	// 실행하고자 하는 컨트랙트의 주소
	destinationContractAddr := common.HexToAddress("0x7Ad6c6EF72e31FD8D7f5d2F584de1dcD3DC4769E")

	// interaction 할 multisig contract instance 생성
	instance, err := multisig.NewMultisig(multisigAddr, client)
	if err != nil {
		fmt.Println(err)
	}
	chainID, err := client.ChainID(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	auth1, err := bind.NewKeyedTransactorWithChainID(firstPk, chainID)
	if err != nil {
		fmt.Println(err)
	}
	auth2, err := bind.NewKeyedTransactorWithChainID(secondPk, chainID)
	if err != nil {
		fmt.Println(err)
	}

	// multisig 컨트랙트의 submitTransaction 함수 실행 결과확인 및 txIdx 확인
	txIdx, err := SendSubmitTransaction(client, instance, auth1, destinationContractAddr, big.NewInt(0), GetDestinationfuncData(374))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(txIdx)
	
	// 해당 txIdx에 대해서 confirmTransaction 실행
	// 먼저 1번 계정으로
	result := SendConfirmTransaction(client, instance, auth1, big.NewInt(txIdx))
	if !result {
		fmt.Println("auth1 confirmTransaction fail")
		return
	}

	// 다음 2번 계정으로
	result = SendConfirmTransaction(client, instance, auth2, big.NewInt(txIdx))
	if !result {
		fmt.Println("auth2 confirmTransaction fail")
		return
	}


	// 위에서 txIdx에 대한 confirm까지 진행되었으므로, 다음은 excuteTransaction을 실행한다.
	// 먼저 1번 계정으로
	result = SendExcuteTransaction(client, instance, auth1, big.NewInt(txIdx))
	if !result {
		fmt.Println("auth1 excuteTransaction fail")
		return
	}

	fmt.Println("success")
}

func SendSubmitTransaction(
		client *ethclient.Client,
		instance *multisig.Multisig,
		auth *bind.TransactOpts,
		destinationContractAddr common.Address,
		value *big.Int,
		data []byte,
	) (int64, error){
	txSubmitTransaction, err := instance.SubmitTransaction(auth, destinationContractAddr, value, data)
	if err != nil {
		return 0, err
	}
	receipt, err := client.TransactionReceipt(context.Background(), txSubmitTransaction.Hash())
	if err != nil {
		return 0, err
	// 현재 보낸 트랜잭션의 상태 확인. 1이면 정상
	} else if receipt.Status == 0 {
		// 에러처리
		return 0, err
	}
	
	// txIdx를 찾아서 return
	txData, err := receipt.MarshalJSON()
	if err != nil {
		return 0, err
	}
	var datas map[string]interface{}
	json.Unmarshal(txData, &datas)
	txIdxHex := datas["logs"].([]interface{})[0].(map[string]interface{})["topics"].([]interface{})[2]
	txIdxHex = txIdxHex.(string)[2:]
	i, _ := strconv.ParseInt(txIdxHex.(string), 16, 64)

	return i, nil
}

func SendConfirmTransaction(
		client *ethclient.Client,
		instance *multisig.Multisig,
		auth *bind.TransactOpts,
		txIdx *big.Int,
	) bool {
	txConfirmTransaction, err := instance.ConfirmTransaction(auth, txIdx)
	if err != nil {
		fmt.Println(err)
		return false
	}
	receipt, err := client.TransactionReceipt(context.Background(), txConfirmTransaction.Hash())
	if err != nil {
		fmt.Println(err)
		return false
	// 현재 보낸 트랜잭션의 상태 확인. 1이면 정상
	} else if receipt.Status == 0 {
		// 에러처리
		fmt.Println(err)
		return false
	}
	return true
}

func SendExcuteTransaction(
		client *ethclient.Client,
		instance *multisig.Multisig,
		auth *bind.TransactOpts,
		txIdx *big.Int,
	) bool {
	txConfirmTransaction, err := instance.ExecuteTransaction(auth, txIdx)
	if err != nil {
		fmt.Println(err)
		return false
	}
	receipt, err := client.TransactionReceipt(context.Background(), txConfirmTransaction.Hash())
	if err != nil {
		fmt.Println(err)
		return false
	// 현재 보낸 트랜잭션의 상태 확인. 1이면 정상
	} else if receipt.Status == 0 {
		// 에러처리
		fmt.Println(err)
		return false
	}
	return true
}