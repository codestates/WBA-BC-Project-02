package main

import (
	"time"
	"context"
	// "encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/codestates/WBA-BC-Project-02/contracts/multisig"
	"github.com/codestates/WBA-BC-Project-02/contracts/util"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	// "golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum/core/types"
)


func main() {
	// server url과 연동
	client, err := ethclient.Dial("https://api.test.wemix.com")
	if err != nil {
		fmt.Println("client connect err")
		log.Fatal(err)
	}

	// 첫번째 계정의 privateKey -> 이 녀석은 toml에서(?)
	firstPk, err := crypto.HexToECDSA("68fe8eda422fbdcd2bc44ced268dfb3c42ec999c2e4f237dfc52af079cee2ddd")
	// 첫번째 계정의 address
	firstAddress := common.HexToAddress("0x406A1Ee747CEEA7a8819Bb999c3B24cc88000461")
	if err != nil {
		fmt.Println("first Pk err")
		log.Fatal(err)
	}
	// toml에서 가져온 multisig contract의 주소
	multisigAddr := common.HexToAddress("0x2603df2DCC9aC15E56346B752e81677CbeE4cF7D")
	// toml에서 가져온 dex 컨트랙트의 주소
	dexContractAddr := common.HexToAddress("0xEC3b843a78D2d9430Acb73279Aed35fC0DEE427B")

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

	firstAddressNonce, err := client.PendingNonceAt(context.Background(), firstAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 아래를 실행하기 전에 redis에 nonce를 저장해야 함!
	//#############################################################

	// db로부터 가져온 유저의 address
	userAddress := "0x406A1Ee747CEEA7a8819Bb999c3B24cc88000461"
	// user로부터 요청받은 mint하고싶은 양 -> 예를 들어 user가 흑철 10개를 바꾸고싶다고 하면 10 : 1 비율로 온체인 트랜잭션에 민트하는 amount는 1이 될 것이다.
	tokenAmount := 2

	data := util.MintDracoTx(userAddress, tokenAmount)

	// multisig 컨트랙트의 submitTransaction 함수 실행 결과확인 및 txIdx 확인
	// data는 util을 통해 만든 byte값
	txIdx, err := SendSubmitTransaction(client, instance, auth1, dexContractAddr, big.NewInt(0), data, big.NewInt(int64(firstAddressNonce)))
	if err != nil {
		fmt.Println("get txIdx Err")
		log.Fatal(err)
	}

	fmt.Println(txIdx)

	count := 0
	for {
		// if count == 20 {
		// 	fmt.Println("false")
		// 	return
		// }
		result, _ := instance.GetTransaction(&bind.CallOpts{}, big.NewInt(txIdx))
		if result.Executed {
			break
		}
		count++
		fmt.Println(result.Executed)
		time.Sleep(time.Millisecond * 500)
		fmt.Println(count)
	}

	// gin 서버에서는 여기까지 실행 후, 위 txIdx를 가지고 실행되었는지 지속해서 요청을 보냄!
	
	fmt.Println("success")
}

func SendSubmitTransaction(
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
