package ciper

import (
	"crypto/aes"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestGenerateEncrypteAndDecrypt(t *testing.T) {
	testText := "0xE56f41b1d58D510E1bbE87411a1150b6b2f543cc"
	ENVFPath := "../../.env"
	keyName := "WEMIXON_DEV_KEY"

	if err := godotenv.Load(ENVFPath); err != nil {
		log.Fatal("Error loading .env file")
	}

	AESDEVKey := os.Getenv(keyName)
	log.Println("AES Encrypt key is ::", AESDEVKey)

	block, err := aes.NewCipher([]byte(AESDEVKey))
	if err != nil {
		log.Fatal(err.Error())
	}

	encryptedText, err := AESEncrypt(block, []byte(testText))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Encrypted text is ::", encryptedText)

	decryptedText, err := AESDecrypt(block, encryptedText)
	if err != nil {
		return
	}

	log.Println("Decrypted text is ::", decryptedText)
}
