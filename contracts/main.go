package main


func GenerateRawTransactionData() {
	mulsigAddr := ""
	toAddr := ""
	tokenAddr := common.HexToAddress("0xbc9F814D03f46B3D10C22e7490aC923de121d6A2")
	tokenAmount := ""

	client, err := ethclient.Dial("https://api.test.wemix.com")
	if err != nil {
		fmt.Println("client error")
	}

	
	if err != nil {
		fmt.Println(err)
	}

	pk, err := crypto.HexToECDSA("c04e38b264d4bf35625090e4764c912718f4c8be1bce5c3c0796365f9732d0b0")
	if err != nil {
		fmt.Println(err)
	}

	publicKey := pk.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err)
	}

	value := big.NewInt(int64(amount.(float64)) * 1000000000000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(value.Bytes(), 32)
	zValue := big.NewInt(0)

	var pdata []byte
	pdata = append(pdata, methodID...)
	pdata = append(pdata, paddedAddress...)
	pdata = append(pdata, paddedAmount...)

	gasLimit := uint64(200000)

	tx := types.NewTransaction(nonce, tokenAddr, zValue, gasLimit, gasPrice, pdata)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), pk)
	if err != nil {
		fmt.Println(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"msg" : "success",
	})
}

func main() {

}