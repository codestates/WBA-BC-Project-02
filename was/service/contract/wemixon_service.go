package contract

import (
	"context"
	"errors"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache/contract"
	"github.com/codestates/WBA-BC-Project-02/was/common/contract/util"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/config"
	"github.com/codestates/WBA-BC-Project-02/was/model/user"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type wemixonService struct {
	userModel user.UserModeler
}

var wemixonInstance *wemixonService

func NewWemixonService(mod user.UserModeler) *wemixonService {
	if wemixonInstance != nil {
		return wemixonInstance
	}
	wemixonInstance = &wemixonService{userModel: mod}
	return wemixonInstance
}

type Method func(string, int) []byte

func ActContract(userAddress string, amount int, method Method) (int, error) {
	client, nonce, err := getClientAndNonce()
	if err != nil {
		return 0, err
	}

	data := method(userAddress, amount)

	if err := saveNonceInRedis(nonce); err != nil {
		return 0, err
	}

	successCount, err := sendTx(client, int64(nonce), data)
	if err != nil || successCount == 0 {
		return 0, errors.New("트랜잭션을 실패했습니다")
	}

	return successCount, nil
}

func getClientAndNonce() (*ethclient.Client, uint64, error) {
	client, err := getConnection()
	if err != nil {
		return nil, 0, err
	}

	nonce, err := getNonce(client)
	if err != nil {
		return nil, 0, err
	}
	return client, nonce, nil
}

func getConnection() (*ethclient.Client, error) {
	return ethclient.Dial(config.ContractConfig.RawURL)
}

func getNonce(client *ethclient.Client) (uint64, error) {
	serverAddr := common.HexToAddress(config.ContractConfig.ServerAddr)
	serverAddrNonce, err := client.PendingNonceAt(context.Background(), serverAddr)
	if err != nil {
		return 0, err
	}
	return serverAddrNonce, nil
}

func saveNonceInRedis(nonceValue uint64) error {
	nonceValues := make([]uint64, 0)
	nonceValues = append(nonceValues, nonceValue)
	nonce := &contract.Nonce{NonceValues: nonceValues}
	if err := cache.Redis.Cache(enum.NonceCacheKey, nonce, 0); err != nil {
		return err
	}
	return nil
}

// sendTx 1 은 1개의 트랜잭션 성공이라는 성공의미, 0 은 실패
func sendTx(client *ethclient.Client, nonce int64, data []byte) (int, error) {
	err := util.SendTransaction(client, int64(nonce), data)
	if err != nil {
		return 0, err
	}
	return 1, nil
}
