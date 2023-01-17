package ciper

import (
	"crypto/aes"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestGenerateEncrypteAndDecrypt(t *testing.T) {
	testText := "0xEC3b843a78D2d9430Acb73279Aed35fC0DEE427B"
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
