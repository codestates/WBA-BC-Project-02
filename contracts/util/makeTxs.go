package util

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


func GenerateMintDracoTx(toAddress string, amount int) []byte {
	// 실행할 함수의 signature 생성, methodID 생성
	transferFnSignature := []byte("mintDraco(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	// fmt.Println(hexutil.Encode(methodID))

	// 아래는 인자로 넣을 data를 추가하는 과정
	// Destination address에 함수 인자값으로 넣을 uint값
	address := common.HexToAddress(toAddress)
	paddedAddress := common.LeftPadBytes(address.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(big.NewInt(int64(amount)).Bytes(), 32)

	var pdata []byte
	pdata = append(pdata, methodID...)
	pdata = append(pdata, paddedAddress...)
	pdata = append(pdata, paddedAmount...)
	// fmt.Println("0x"+hex.EncodeToString(pdata))
	return pdata
}

