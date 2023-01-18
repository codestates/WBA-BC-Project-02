package util

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

// user에게 draco를 민팅해주는 함수 (이녀석은 오프체인데이터 → 온체인 데이터로 변경할 때만 실행)
func MintDracoTx(user string, amount int) []byte {
	// 실행할 함수의 signature 생성, methodID 생성
	transferFnSignature := []byte("mintDraco(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	// fmt.Println(hexutil.Encode(methodID))

	// 아래는 인자로 넣을 data를 추가하는 과정
	// Destination address에 함수 인자값으로 넣을 uint값
	address := common.HexToAddress(user)
	paddedAddress := common.LeftPadBytes(address.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(big.NewInt(int64(amount)).Bytes(), 32)

	var pdata []byte
	pdata = append(pdata, methodID...)
	pdata = append(pdata, paddedAddress...)
	pdata = append(pdata, paddedAmount...)
	// fmt.Println("0x"+hex.EncodeToString(pdata))
	return pdata
}

// user에게 draco를 민팅해주는 함수 (이녀석은 오프체인데이터 → 온체인 데이터로 변경할 때만 실행)
func MintTigTx(user string, amount int) []byte {
	// 실행할 함수의 signature 생성, methodID 생성
	transferFnSignature := []byte("mintTig(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	// fmt.Println(hexutil.Encode(methodID))

	// 아래는 인자로 넣을 data를 추가하는 과정
	// Destination address에 함수 인자값으로 넣을 uint값
	address := common.HexToAddress(user)
	paddedAddress := common.LeftPadBytes(address.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(big.NewInt(int64(amount)).Bytes(), 32)

	var pdata []byte
	pdata = append(pdata, methodID...)
	pdata = append(pdata, paddedAddress...)
	pdata = append(pdata, paddedAmount...)
	// fmt.Println("0x"+hex.EncodeToString(pdata))
	return pdata
}

// user가 draco를 팔고 credit을 구매하는 함수
func BuyCreditbyDracoTx(user string, amount int) []byte {
	// 실행할 함수의 signature 생성, methodID 생성
	transferFnSignature := []byte("buyCreditByDraco(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	// fmt.Println(hexutil.Encode(methodID))

	// 아래는 인자로 넣을 data를 추가하는 과정
	// Destination address에 함수 인자값으로 넣을 uint값
	address := common.HexToAddress(user)
	paddedAddress := common.LeftPadBytes(address.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(big.NewInt(int64(amount)).Bytes(), 32)

	var pdata []byte
	pdata = append(pdata, methodID...)
	pdata = append(pdata, paddedAddress...)
	pdata = append(pdata, paddedAmount...)
	// fmt.Println("0x"+hex.EncodeToString(pdata))
	return pdata
}

// user가 credit를 팔고 draco를 구매하는 함수
func BuyDracoByCreditTx(user string, amount int) []byte {
	// 실행할 함수의 signature 생성, methodID 생성
	transferFnSignature := []byte("buyDracoByCredit(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	// fmt.Println(hexutil.Encode(methodID))

	// 아래는 인자로 넣을 data를 추가하는 과정
	// Destination address에 함수 인자값으로 넣을 uint값
	address := common.HexToAddress(user)
	paddedAddress := common.LeftPadBytes(address.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(big.NewInt(int64(amount)).Bytes(), 32)

	var pdata []byte
	pdata = append(pdata, methodID...)
	pdata = append(pdata, paddedAddress...)
	pdata = append(pdata, paddedAmount...)
	// fmt.Println("0x"+hex.EncodeToString(pdata))
	return pdata
}

// credit토큰을 wemix로 변경해주는 함수
func CreditToWemixTx(user string, amount int) []byte {
	// 실행할 함수의 signature 생성, methodID 생성
	transferFnSignature := []byte("creditToWemix(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	// fmt.Println(hexutil.Encode(methodID))

	// 아래는 인자로 넣을 data를 추가하는 과정
	// Destination address에 함수 인자값으로 넣을 uint값
	address := common.HexToAddress(user)
	paddedAddress := common.LeftPadBytes(address.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(big.NewInt(int64(amount)).Bytes(), 32)

	var pdata []byte
	pdata = append(pdata, methodID...)
	pdata = append(pdata, paddedAddress...)
	pdata = append(pdata, paddedAmount...)
	// fmt.Println("0x"+hex.EncodeToString(pdata))
	return pdata
}

// user가 credit을 팔고 tig를 구매하는 함수
func BuyTigByCreditTx(user string, amount int) []byte {
	// 실행할 함수의 signature 생성, methodID 생성
	transferFnSignature := []byte("buyTigByCredit(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	// fmt.Println(hexutil.Encode(methodID))

	// 아래는 인자로 넣을 data를 추가하는 과정
	// Destination address에 함수 인자값으로 넣을 uint값
	address := common.HexToAddress(user)
	paddedAddress := common.LeftPadBytes(address.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(big.NewInt(int64(amount)).Bytes(), 32)

	var pdata []byte
	pdata = append(pdata, methodID...)
	pdata = append(pdata, paddedAddress...)
	pdata = append(pdata, paddedAmount...)
	// fmt.Println("0x"+hex.EncodeToString(pdata))
	return pdata
}

// user가 Tig를 팔고 credit을 구매하는 함수
func BuyCreditByTigTx(user string, amount int) []byte {
	// 실행할 함수의 signature 생성, methodID 생성
	transferFnSignature := []byte("buyCreditByTig(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	// fmt.Println(hexutil.Encode(methodID))

	// 아래는 인자로 넣을 data를 추가하는 과정
	// Destination address에 함수 인자값으로 넣을 uint값
	address := common.HexToAddress(user)
	paddedAddress := common.LeftPadBytes(address.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(big.NewInt(int64(amount)).Bytes(), 32)

	var pdata []byte
	pdata = append(pdata, methodID...)
	pdata = append(pdata, paddedAddress...)
	pdata = append(pdata, paddedAmount...)
	// fmt.Println("0x"+hex.EncodeToString(pdata))
	return pdata
}
