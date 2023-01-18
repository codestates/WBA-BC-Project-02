package main

import (
	"crypto/aes"
	"fmt"
	"github.com/codestates/WBA-BC-Project-02/common/ciper"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestMMMM(t *testing.T) {
	testDecrypt()
}

func testDecrypt() {
	ENVFPath := "../.env"
	if err := godotenv.Load(ENVFPath); err != nil {
		log.Fatal("Error loading .env file")
	}
	key := os.Getenv("WEMIXON_DEV_KEY")
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err.Error())
	}

	target := "tczzOhlnAtweLPNP4rCkifhWQBXbWBEJjwZbyUJpi9M7jF+2nx+j/6Hsy9HERZlebr7Bm6cANlGC4mgQOjZPwl/scVRzbn9CDjMcnSec0Yw="
	decryptedText, err := ciper.AESDecrypt(block, target)
	if err != nil {
		fmt.Println("ttttt")
	}
	fmt.Println(decryptedText)
}
