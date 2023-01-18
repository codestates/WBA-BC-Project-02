package practice

//
//import (
//	"context"
//	// "encoding/hex"
//	"encoding/json"
//	"fmt"
//	"log"
//	"math/big"
//	"strconv"
//
//	"github.com/codestates/WBA-BC-Project-02/contracts/multisig"
//	"github.com/codestates/WBA-BC-Project-02/contracts/util"
//	"github.com/ethereum/go-ethereum/common"
//	// "github.com/ethereum/go-ethereum/common/hexutil"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/ethereum/go-ethereum/ethclient"
//	// "golang.org/x/crypto/sha3"
//	"github.com/ethereum/go-ethereum/core/types"
//)
//
//// func GetDestinationfuncData(num int) []byte {
//// 	// 실행할 함수의 signature 생성, methodID 생성
//// 	transferFnSignature := []byte("callMe(uint256)")
//// 	hash := sha3.NewLegacyKeccak256()
//// 	hash.Write(transferFnSignature)
//// 	methodID := hash.Sum(nil)[:4]
//// 	// fmt.Println(hexutil.Encode(methodID))
//
//// 	// 아래는 인자로 넣을 data를 추가하는 과정
//// 	// Destination address에 함수 인자값으로 넣을 uint값
//// 	bigNum := big.NewInt(int64(num))
//// 	paddedNum := common.LeftPadBytes(bigNum.Bytes(), 32)
//
//// 	var pdata []byte
//// 	pdata = append(pdata, methodID...)
//// 	pdata = append(pdata, paddedNum...)
//// 	// fmt.Println("0x"+hex.EncodeToString(pdata))
//// 	return pdata
//// }
//
//func main() {
//	// server url과 연동
//	client, err := ethclient.Dial("https://api.test.wemix.com")
//	if err != nil {
//		fmt.Println("client connect err")
//		log.Fatal(err)
//	}
//
//	// 첫번째 계정의 privateKey -> 이 녀석은 toml에서(?)
//	firstPk, err := crypto.HexToECDSA("68fe8eda422fbdcd2bc44ced268dfb3c42ec999c2e4f237dfc52af079cee2ddd")
//	if err != nil {
//		fmt.Println("first Pk err")
//		log.Fatal(err)
//	}
//	// 첫번째 계정의 address
//	firstAddress := common.HexToAddress("0x406A1Ee747CEEA7a8819Bb999c3B24cc88000461")
//	// 두번째 계정의 privateKey -> 이 녀석은 DB에서(?)
//	secondPk, err := crypto.HexToECDSA("2e33f33e005a4ed8986cbf4770ee022b4d28b59d7d078bc5a82e7be2c97e8f44")
//	if err != nil {
//		fmt.Println("second Pk err")
//		log.Fatal(err)
//	}
//	// toml에서 가져온 multisig contract의 주소
//	multisigAddr := common.HexToAddress("0x2603df2DCC9aC15E56346B752e81677CbeE4cF7D")
//	// toml에서 가져온 dex 컨트랙트의 주소
//	dexContractAddr := common.HexToAddress("0xEC3b843a78D2d9430Acb73279Aed35fC0DEE427B")
//
//	// interaction 할 multisig contract instance 생성
//	instance, err := multisig.NewMultisig(multisigAddr, client)
//	if err != nil {
//		fmt.Println("get instance Err")
//		fmt.Println(err)
//	}
//	chainID, err := client.ChainID(context.TODO())
//	if err != nil {
//		fmt.Println("get chainId Err")
//		fmt.Println(err)
//	}
//	auth1, err := bind.NewKeyedTransactorWithChainID(firstPk, chainID)
//	if err != nil {
//		fmt.Println("bind auth1 Err")
//		fmt.Println(err)
//	}
//	auth2, err := bind.NewKeyedTransactorWithChainID(secondPk, chainID)
//	if err != nil {
//		fmt.Println("bind auth2 Err")
//		fmt.Println(err)
//	}
//
//	firstAddressNonce, err := client.PendingNonceAt(context.Background(), firstAddress)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	// 아래를 실행하기 전에 redis에 nonce를 저장해야 함!
//	//#############################################################
//
//	// db로부터 가져온 유저의 address
//	userAddress := "0x406A1Ee747CEEA7a8819Bb999c3B24cc88000461"
//	// user로부터 요청받은 mint하고싶은 양 -> 예를 들어 user가 흑철 10개를 바꾸고싶다고 하면 10 : 1 비율로 온체인 트랜잭션에 민트하는 amount는 1이 될 것이다.
//	tokenAmount := 2
//
//	data := util.CreditToWemixTx(userAddress, tokenAmount)
//
//	// multisig 컨트랙트의 submitTransaction 함수 실행 결과확인 및 txIdx 확인
//	// data는 util을 통해 만든 byte값
//	txIdx, err := SendSubmitTransaction(client, instance, auth1, dexContractAddr, big.NewInt(0), data, big.NewInt(int64(firstAddressNonce)))
//	if err != nil {
//		fmt.Println("get txIdx Err")
//		log.Fatal(err)
//	}
//
//	// gin 서버에서는 여기까지 실행 후, 위 txIdx를 가지고 실행되었는지 지속해서 요청을 보냄!
//	//########################################################
//	// 아래서부터는 Daemon(signer)에서 실행하는 로직
//
//	// redis를 통해 먼저 gin 서버에서 보낸 요청인지 확인해야함!
//	// 해당 txIdx에 대해서 confirmTransaction 실행
//	ch1 := make(chan bool)
//	ch2 := make(chan bool)
//	// 먼저 1번 계정으로
//	go SendConfirmTransaction(ch1, client, instance, auth1, big.NewInt(txIdx))
//	// 다음 2번 계정으로
//	go SendConfirmTransaction(ch2, client, instance, auth2, big.NewInt(txIdx))
//
//	confirmCount := 0
//
//	for {
//		if confirmCount == 2 {
//			break
//		}
//		ch1Value := <-ch1
//		ch2Value := <-ch2
//		if ch1Value {
//			confirmCount++
//		} else {
//			break
//		}
//		if ch2Value {
//			confirmCount++
//		} else {
//			break
//		}
//	}
//
//	// 위에서 txIdx에 대한 confirm까지 진행되었으므로, 다음은 excuteTransaction을 실행한다.
//	result := SendExcuteTransaction(client, instance, auth1, big.NewInt(txIdx))
//	if !result {
//		fmt.Println("auth1 excuteTransaction fail")
//		return
//	}
//
//	fmt.Println("success")
//}
//
//func SendSubmitTransaction(
//	client *ethclient.Client,
//	instance *multisig.Multisig,
//	auth *bind.TransactOpts,
//	dexContractAddr common.Address,
//	value *big.Int,
//	data []byte,
//	nonce *big.Int,
//) (int64, error) {
//	fmt.Println("Sendging...")
//	txSubmitTransaction, err := instance.SubmitTransaction(auth, dexContractAddr, value, data, nonce)
//	if err != nil {
//		fmt.Println("SendSubmitTransaction Err")
//		return 0, err
//	}
//
//	var receipt *types.Receipt
//	for {
//		receipt, err = client.TransactionReceipt(context.Background(), txSubmitTransaction.Hash())
//		if receipt != nil {
//			break
//		}
//	}
//	if err != nil {
//		fmt.Println("get submitTrasaction Receipt Err")
//		return 0, err
//		// 현재 보낸 트랜잭션의 상태 확인. 1이면 정상
//	} else if receipt.Status == 0 {
//		// 에러처리
//		fmt.Println("Status Err")
//		return 0, err
//	}
//
//	// txIdx를 찾아서 return
//	txData, err := receipt.MarshalJSON()
//	if err != nil {
//		fmt.Println("txData Err")
//		return 0, err
//	}
//	var datas map[string]interface{}
//	json.Unmarshal(txData, &datas)
//	txIdxHex := datas["logs"].([]interface{})[0].(map[string]interface{})["topics"].([]interface{})[2]
//	txIdxHex = txIdxHex.(string)[2:]
//	i, _ := strconv.ParseInt(txIdxHex.(string), 16, 64)
//	return i, nil
//}
//
//func SendConfirmTransaction(
//	ch chan bool,
//	client *ethclient.Client,
//	instance *multisig.Multisig,
//	auth *bind.TransactOpts,
//	txIdx *big.Int,
//) {
//	fmt.Println("Confirming...")
//	txConfirmTransaction, err := instance.ConfirmTransaction(auth, txIdx)
//	if err != nil {
//		fmt.Println("confirmTransaction Err", err)
//		ch <- false
//		return
//	}
//	var receipt *types.Receipt
//	for {
//		receipt, _ = client.TransactionReceipt(context.Background(), txConfirmTransaction.Hash())
//		if receipt != nil {
//			break
//		}
//	}
//	if err != nil {
//		fmt.Println("receipt Err", err)
//		ch <- false
//		return
//		// 현재 보낸 트랜잭션의 상태 확인. 1이면 정상
//	} else if receipt.Status == 0 {
//		// 에러처리
//		fmt.Println("receipt status Err", err)
//		ch <- false
//		return
//	}
//	ch <- true
//}
//
//func SendExcuteTransaction(
//	client *ethclient.Client,
//	instance *multisig.Multisig,
//	auth *bind.TransactOpts,
//	txIdx *big.Int,
//) bool {
//	fmt.Println("Executing...")
//	txExcuteTransaction, err := instance.ExecuteTransaction(auth, txIdx)
//	if err != nil {
//		fmt.Println("Execute Err", err)
//		return false
//	}
//	var receipt *types.Receipt
//	for {
//		receipt, _ = client.TransactionReceipt(context.Background(), txExcuteTransaction.Hash())
//		if receipt != nil {
//			break
//		}
//	}
//	if err != nil {
//		fmt.Println("Execute Receipt Err", err)
//		return false
//		// 현재 보낸 트랜잭션의 상태 확인. 1이면 정상
//	} else if receipt.Status == 0 {
//		// 에러처리
//		fmt.Println("Execute Status Err", err)
//		return false
//	}
//	return true
//}
