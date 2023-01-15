package ciper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/codestates/WBA-BC-Project-02/common/enum"
	"io"
	"log"
	"os"
)

var cipherBlock cipher.Block

var cipherKey string

// LoadCipherKey set AES cipher key, AES cipher key is private variable
func LoadCipherKey(mode string) {
	if cipherKey != "" {
		return
	}

	switch mode {
	case enum.DevMode:
		setCipherKey(enum.CipherDevKey)
	case enum.ProdMode:
		setCipherKey(enum.CipherProdKey)
	}
}

func setCipherKey(keyName string) {
	cipherKey = os.Getenv(keyName)
}

func LoadCipherBlock() {
	if cipherKey == "" {
		log.Fatal("please call LoadCipherKey() function")
	}

	block, err := aes.NewCipher([]byte(cipherKey)) // AES 대칭키 암호화 블록 생성
	if err != nil {
		log.Fatal(err.Error())
	}
	cipherBlock = block
}

func GetCipherBlock() cipher.Block {
	return cipherBlock
}

func AESEncrypt(block cipher.Block, plainText []byte) (string, error) {
	if mod := len(plainText) % aes.BlockSize; mod != 0 {
		padding := make([]byte, aes.BlockSize-mod)
		plainText = append(plainText, padding...)
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil { // 랜덤 값을 초기화 벡터에 넣어줌
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func AESDecrypt(b cipher.Block, cipherText string) (string, error) {
	decodedCipherText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	if len(decodedCipherText)%aes.BlockSize != 0 {
		return "", errors.New("the length of the encrypted data must be a multiple of the block size")
	}

	iv := decodedCipherText[:aes.BlockSize]
	decodedCipherText = decodedCipherText[aes.BlockSize:]

	plainText := make([]byte, len(decodedCipherText))
	mode := cipher.NewCBCDecrypter(b, iv)
	mode.CryptBlocks(plainText, decodedCipherText)

	removedPaddingText := bytes.Trim(plainText, "\x00")

	return string(removedPaddingText), nil
}
