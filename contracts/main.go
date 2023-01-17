package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"log"

	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	// "github.com/chenzhijie/go-web3"
	// "github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)


func SubmitTrasaction(dData []byte) {
	// 첫번째 계정의 privateKey를 통한 address 도출
	firstPk, err := crypto.HexToECDSA("9734a6fab2886cdcc017c9441790681fdf99dbdbe16da22506d070a66f28755f")
	if err != nil {
		log.Fatal(err)
	}
	firstPublicKey := firstPk.Public()
    firstPublicKeyECDSA, ok := firstPublicKey.(*ecdsa.PublicKey)
    if !ok {
        fmt.Println("fail convert, publickey")
    }
    firstAddress := crypto.PubkeyToAddress(*firstPublicKeyECDSA)
	// multisig contract의 주소
	multisigAddr := common.HexToAddress("0xb80cE90781B7ed255695016fBcdd8389453ddA89")
	// 실행하고자 하는 컨트랙트의 주소
	destinationContractAddr := common.HexToAddress("0x7Ad6c6EF72e31FD8D7f5d2F584de1dcD3DC4769E")

	// server url과 연동
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		fmt.Println("client error")
	}

	// firstAddress의 nonce값 추출
	nonce, err := client.PendingNonceAt(context.Background(), firstAddress)
	if err != nil {
		fmt.Println(err)
	}

	// gasPrice 설정
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	gasLimit := uint64(200000)

	// 이녀석은 Tx를 보낼 때 전송할 value
	zValue := big.NewInt(0)

	// 실행할 함수의 signature 생성, methodID 생성
	transferFnSignature := []byte("submitTransaction(address,uint256,bytes)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	// 아래는 인자로 넣을 data를 추가하는 과정
	// Destination Address와 보낼 value, Destination에 도달해야 할 Transaction Data가 포함된다.
	paddedAddr := common.LeftPadBytes(destinationContractAddr.Bytes(), 32)
	paddedValue := common.LeftPadBytes(zValue.Bytes(), 32)
	paddedData := common.LeftPadBytes(dData, 32)
	fmt.Println(zValue.Bytes())
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddr...)
	data = append(data, paddedValue...)
	data = append(data, paddedData...)
	
	// 인자값은 차례대로 계정논스, 멀티시그 컨트랙트 주소, 전송할 value, 가스제한, 가스비, 전송할 데이터
	tx := types.NewTransaction(nonce, multisigAddr, zValue, gasLimit, gasPrice, data)
	fmt.Println(nonce)
	fmt.Println(multisigAddr)
	fmt.Println(zValue)
	fmt.Println(gasLimit)
	fmt.Println(gasPrice)
	fmt.Println(data)

	chainID := big.NewInt(1337)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), firstPk)
	if err != nil {
		fmt.Println(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tx)
	fmt.Println(signedTx)
}

func GetDestinationfuncData() []byte {
	// 실행할 함수의 signature 생성, methodID 생성
	transferFnSignature := []byte("callMe(uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	// 아래는 인자로 넣을 data를 추가하는 과정
	// Destination address에 함수 인자값으로 넣을 uint값
	num := big.NewInt(339)
	paddedNum := common.LeftPadBytes(num.Bytes(), 32)

	var pdata []byte
	pdata = append(pdata, methodID...)
	pdata = append(pdata, paddedNum...)
	fmt.Println("0x"+hex.EncodeToString(pdata))
	return pdata
}


func main() {
	// 두번째 계정의 privateKey를 통한 address 도출
	// secondPk, err := crypto.HexToECDSA("30e9d002076fb2a3a10322738a62260a278073d6c42996b9f374032259cc8801")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// secondPublicKey := firstPk.Public()
    // secondPublicKeyECDSA, ok := secondPublicKey.(*ecdsa.PublicKey)
    // if !ok {
    //     fmt.Println("fail convert, publickey")
    // }
    // secondAddress := crypto.PubkeyToAddress(*secondPublicKeyECDSA)

	data := GetDestinationfuncData()
	SubmitTrasaction(data)
}