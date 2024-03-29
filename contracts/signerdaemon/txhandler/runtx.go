package txhandler

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/codestates/WBA-BC-Project-02/contracts/multisig"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	// "golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum/core/types"
)

func RunTx(firstPkS string, secondPkS string, multisigAddrS string, txIdx *big.Int) {
	client, err := ethclient.Dial("https://api.test.wemix.com")
	if err != nil {
		fmt.Println("client connect err")
		log.Fatal(err)
	}

	firstPk, err := crypto.HexToECDSA(firstPkS)
	if err != nil {
		fmt.Println("first Pk err")
		log.Fatal(err)
	}
	secondPk, err := crypto.HexToECDSA(secondPkS)
	if err != nil {
		fmt.Println("second Pk err")
		log.Fatal(err)
	}
	multisigAddr := common.HexToAddress(multisigAddrS)
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

	// redis를 통해 먼저 gin 서버에서 보낸 요청인지 확인해야함!
	// 해당 txIdx에 대해서 confirmTransaction 실행
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	// 먼저 1번 계정으로
	go SendConfirmTransaction(ch1, client, instance, auth1, txIdx)
	// 다음 2번 계정으로
	go SendConfirmTransaction(ch2, client, instance, auth2, txIdx)

	confirmCount := 0

	for {
		if confirmCount == 2 {
			break
		}
		ch1Value := <-ch1
		ch2Value := <-ch2
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
	result := SendExcuteTransaction(client, instance, auth1, txIdx)
	if !result {
		fmt.Println("auth1 excuteTransaction fail")
		return
	}

	fmt.Println("success")
}


func SendConfirmTransaction(
	ch chan bool,
	client *ethclient.Client,
	instance *multisig.Multisig,
	auth *bind.TransactOpts,
	txIdx *big.Int,
) {
	fmt.Println("Confirming...")
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
	fmt.Println("Executing...")
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
