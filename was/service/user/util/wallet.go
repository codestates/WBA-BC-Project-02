package util

import (
	"crypto/elliptic"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type Wallet struct {
	Address    string
	PrivateKey string
	PublicKey  string
}

func NewMnemonic() (string, error) {
	entropy, err := hdwallet.NewEntropy(256)
	if err != nil {
		return "", err
	}

	mnemonicFromEntropy, err := hdwallet.NewMnemonicFromEntropy(entropy)
	if err != nil {
		return "", err
	}
	return mnemonicFromEntropy, nil
}

func NewWallet(mneminic string) (*Wallet, error) {
	wallet, err := createWallet(mneminic)
	if err != nil {
		return nil, err
	}

	account, err := createAccount(wallet)
	if err != nil {
		return nil, err
	}

	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		return nil, err
	}
	pubk := privateKey.PublicKey

	priKey := fmt.Sprintf("%x\n", crypto.FromECDSA(privateKey))
	pubKey := fmt.Sprintf("%x\n", elliptic.Marshal(pubk.Curve, pubk.X, pubk.Y))

	return &Wallet{
		Address:    account.Address.Hex(),
		PrivateKey: priKey,
		PublicKey:  pubKey,
	}, nil
}

func createWallet(mnemonic string) (*hdwallet.Wallet, error) {
	seed, err := hdwallet.NewSeedFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		return nil, err
	}
	return wallet, err
}

func createAccount(wallet *hdwallet.Wallet) (accounts.Account, error) {
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")

	account, err := wallet.Derive(path, false)
	if err != nil {
		return accounts.Account{}, err
	}
	return account, nil
}
