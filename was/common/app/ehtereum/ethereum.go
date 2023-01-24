package ehtereum

import (
	"crypto/ecdsa"
	"github.com/codestates/WBA-BC-Project-02/common/enum"
	wasCommon "github.com/codestates/WBA-BC-Project-02/was/common"
	"github.com/codestates/WBA-BC-Project-02/was/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

var EtheClient *ethclient.Client

func LoadEthereumClient() {
	if EtheClient != nil {
		return
	}

	client, err := ethclient.Dial("https://api.test.wemix.com")
	if err != nil {
		logger.AppLog.Error("client error")
	}

	EtheClient = client
}

type TransferAction struct {
	ownerAddress common.Address // 토큰 컨트랙트의 오너 -> 토큰에 대한 정보들 볼 수 있음
	tokenAddress common.Address // 토큰 컨트랙트 주소
	fromAddress  common.Address // 트랜잭션을 일으키는 자신의 주소
	toAddress    common.Address // 보낼 주소
	value        *big.Int       // 토큰 전송은 ETH 를 전송할 필요가 없으므로 값을 0으로
	amount       *big.Int       // 전송할 토큰양
	gasLimit     uint64
	gasPrice     *big.Int
	privateKey   *ecdsa.PrivateKey
	padData      []byte // transferFnSignature 를 통해 생성한 데이터, 예) []byte("transfer(address,uint256)") -> toAddress, amount  를 넣어 생성해야함
	transaction  *types.Transaction
}

func NewTransferAction() *TransferAction {
	return &TransferAction{}
}

func (e *TransferAction) SetOwnerAddress(address string) *TransferAction {
	e.ownerAddress = common.HexToAddress(address)
	return e
}

func (e *TransferAction) SetTokenAddress(address string) *TransferAction {
	e.ownerAddress = common.HexToAddress(address)
	return e
}

func (e *TransferAction) SetFromAddress(address string) *TransferAction {
	e.fromAddress = common.HexToAddress(address)
	return e
}

func (e *TransferAction) SetToAddress(address string) *TransferAction {
	e.toAddress = common.HexToAddress(address)
	return e
}

func (e *TransferAction) SetValue(value int64) *TransferAction {
	e.value = big.NewInt(value)
	return e
}

func (e *TransferAction) SetAmount(amount int64) *TransferAction {
	e.amount = big.NewInt(amount)
	return e
}

func (e *TransferAction) LoadGasPrice() *TransferAction {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	gasPrice, err := EtheClient.SuggestGasPrice(ctx)
	if err != nil {
		logger.AppLog.Error(err)
	}

	e.gasPrice = gasPrice

	return e
}

func (e *TransferAction) SetPadData(data ...[]byte) *TransferAction {
	var pdata []byte
	for _, d := range data {
		pdata = append(pdata, d...)
	}
	e.padData = pdata
	return e
}

func (e *TransferAction) SendTokenTransaction() (string, error) {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	nonce, err := EtheClient.PendingNonceAt(ctx, e.fromAddress)
	if err != nil {
		return enum.BlankSTR, err
	}

	e.transaction = types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &e.tokenAddress,
		Value:    e.value,
		Gas:      e.gasLimit,
		GasPrice: e.gasPrice,
		Data:     e.padData,
	})

	chainID, err := EtheClient.NetworkID(ctx)
	if err != nil {
		return enum.BlankSTR, err
	}

	// 트랜잭션 서명
	signedTx, err := types.SignTx(e.transaction, types.NewEIP155Signer(chainID), e.privateKey)
	if err != nil {
		return enum.BlankSTR, err
	}

	// 트랜잭션 전송
	err = EtheClient.SendTransaction(ctx, signedTx)
	if err != nil {
		return enum.BlankSTR, err
	}

	//tx.hash를 이용해 전송결과를 확인
	//예)0x016430c748dad98865afb61038537f3ab8f504b56910769d328e7d857be7886a
	return signedTx.Hash().Hex(), nil
}
