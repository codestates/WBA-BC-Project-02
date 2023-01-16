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
	"github.com/ethereum/go-ethereum/core/types"
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
	client, err := ethclient.Dial("https://api.test.wemix.com")
	if err != nil {
		fmt.Println("client connect err")
		log.Fatal(err)
	}

	// 첫번째 계정의 privateKey
	firstPk, err := crypto.HexToECDSA("ed60003d1d768f0c2690b9ae3c418a0a1515abd0569de3701ac273bee0257003")
	if err != nil {
		fmt.Println("first Pk err")
		log.Fatal(err)
	}
	// 두번째 계정의 privateKey
	secondPk, err := crypto.HexToECDSA("2e33f33e005a4ed8986cbf4770ee022b4d28b59d7d078bc5a82e7be2c97e8f44")
	if err != nil {
		fmt.Println("second Pk err")
		log.Fatal(err)
	}
	// multisig contract의 주소
	multisigAddr := common.HexToAddress("0x9aC473C093DEbFA01e23a89d274EE07Ff0482Ee4")
	// 실행하고자 하는 컨트랙트의 주소
	destinationContractAddr := common.HexToAddress("0x4DA3669a8a527461c7f85bf33fD6379C123CE3E3")

	// interaction 할 multisig contract instance 생성
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
	auth2, err := bind.NewKeyedTransactorWithChainID(secondPk, chainID)
	if err != nil {
		fmt.Println("bind auth2 Err")
		fmt.Println(err)
	}

	// multisig 컨트랙트의 submitTransaction 함수 실행 결과확인 및 txIdx 확인
	txIdx, err := SendSubmitTransaction(client, instance, auth1, destinationContractAddr, big.NewInt(0), GetDestinationfuncData(374))
	if err != nil {
		fmt.Println("get txIdx Err")
		log.Fatal(err)
	}

	fmt.Println(txIdx)
	
	// 해당 txIdx에 대해서 confirmTransaction 실행
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	// 먼저 1번 계정으로
	go SendConfirmTransaction(ch1, client, instance, auth1, big.NewInt(txIdx))
	// 다음 2번 계정으로
	go SendConfirmTransaction(ch2, client, instance, auth2, big.NewInt(txIdx))
	
	confirmCount := 0

	for {
		if confirmCount == 2 {
			break
		}
		ch1Value := <- ch1
		ch2Value := <- ch2
		if ch1Value {
			confirmCount++
		} else {
			break
		}
		if ch2Value {
			confirmCount++
		} else {
			break
		}
	}

	// 위에서 txIdx에 대한 confirm까지 진행되었으므로, 다음은 excuteTransaction을 실행한다.
	result := SendExcuteTransaction(client, instance, auth1, big.NewInt(txIdx))
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
		fmt.Println("SendSubmitTransaction Err")
		return 0, err
	}
	fmt.Println(txSubmitTransaction.Hash())
	
	var receipt *types.Receipt
	for {
		receipt, _ = client.TransactionReceipt(context.Background(), txSubmitTransaction.Hash())
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

func SendConfirmTransaction(
		ch chan bool,
		client *ethclient.Client,
		instance *multisig.Multisig,
		auth *bind.TransactOpts,
		txIdx *big.Int,
	) {
	txConfirmTransaction, err := instance.ConfirmTransaction(auth, txIdx)
	if err != nil {
		fmt.Println("confirmTransaction Err", err)
		ch <- false
		return
	}
	var receipt *types.Receipt
	for {
		receipt, _ = client.TransactionReceipt(context.Background(), txConfirmTransaction.Hash())
		if receipt != nil {
			break
		}
	}
	if err != nil {
		fmt.Println("receipt Err", err)
		ch <- false
		return
	// 현재 보낸 트랜잭션의 상태 확인. 1이면 정상
	} else if receipt.Status == 0 {
		// 에러처리
		fmt.Println("receipt status Err", err)
		ch <- false
		return
	}
	ch <- true
}

func SendExcuteTransaction(
		client *ethclient.Client,
		instance *multisig.Multisig,
		auth *bind.TransactOpts,
		txIdx *big.Int,
	) bool {
	txExcuteTransaction, err := instance.ExecuteTransaction(auth, txIdx)
	if err != nil {
		fmt.Println("Execute Err", err)
		return false
	}
	var receipt *types.Receipt
	for {
		receipt, _ = client.TransactionReceipt(context.Background(), txExcuteTransaction.Hash())
		if receipt != nil {
			break
		}
	}
	if err != nil {
		fmt.Println("Execute Receipt Err", err)
		return false
	// 현재 보낸 트랜잭션의 상태 확인. 1이면 정상
	} else if receipt.Status == 0 {
		// 에러처리
		fmt.Println("Execute Status Err", err)
		return false
	}
	return true
}