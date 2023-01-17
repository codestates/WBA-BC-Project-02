package temp

//
//import (
//	"context"
//	"fmt"
//	"github.com/codestates/WBA-BC-Project-02/was/config"
//
//	"github.com/codestates/WBA-BC-Project-02/temp/util"
//
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/ethclient"
//)
//
//func main() {
//	// 연결하는 과정
//	client, err := ethclient.Dial(config.ContractConfig.RawURL) // 그대로 사용
//	if err != nil {
//		panic(err)
//	}
//	firstAddress := common.HexToAddress(config.ContractConfig.ServerAddr) // 서버 계정 주소 이녀석은 toml에서 가져와야 함
//	firstAddressNonce, err := client.PendingNonceAt(context.Background(), firstAddress)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// 원하는 함수를 선택하여 data를 만들어냄
//	// userAddress 등 인자값은 http프로토콜 요청을 통해 받아온 가변 데이터
//	// 기본적으로 address는 그냥 string type으로 넣으면 됨
//	userAddress := "0x406A1Ee747CEEA7a8819Bb999c3B24cc88000461" // http로 받아온 유저 정보를 통해 address를 가져옴
//	tokenAmount := 10                                           // mint요청의 경우 "만" 10 : 1의 비율로 변경해서 넣어줌 그게 아니면 사용자가 보낸 요청 값
//	data := util.BuyCreditbyDracoTx(userAddress, tokenAmount)
//
//	//######################################################################################################
//	// 여기서 redis에 firstAddressNonce값을 저장하는 로직을 추가해주세요
//
//	// client객체와 nonce, data를 담아서 tx를 보냄
//	err = util.SendTransaction(client, int64(firstAddressNonce), data)
//	if err != nil {
//		// 유저한테 실패 반환
//		return
//	}
//	// 유저한테 성공 반환
//}
