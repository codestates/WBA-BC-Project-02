package dev

import "github.com/codestates/WBA-BC-Project-02/common/ciper"

type Contract struct {
	ServerPrivateKey string
	MultiSigAddr     string
	DexAddr          string
	CreditAddr       string
	DracoAddr        string
	TigAddr          string
	ServerAddr       string
	RawURL           string
}

func (c *Contract) DecryptFields() error {
	fk, err := ciper.AESDecrypt(ciper.GetCipherBlock(), c.ServerPrivateKey)
	if err != nil {
		return err
	}
	c.ServerPrivateKey = fk

	ms, err := ciper.AESDecrypt(ciper.GetCipherBlock(), c.MultiSigAddr)
	if err != nil {
		return err
	}
	c.MultiSigAddr = ms

	da, err := ciper.AESDecrypt(ciper.GetCipherBlock(), c.DexAddr)
	if err != nil {
		return err
	}
	c.DexAddr = da

	ca, err := ciper.AESDecrypt(ciper.GetCipherBlock(), c.CreditAddr)
	if err != nil {
		return err
	}
	c.CreditAddr = ca

	dca, err := ciper.AESDecrypt(ciper.GetCipherBlock(), c.DracoAddr)
	if err != nil {
		return err
	}
	c.DracoAddr = dca

	ta, err := ciper.AESDecrypt(ciper.GetCipherBlock(), c.TigAddr)
	if err != nil {
		return err
	}
	c.TigAddr = ta
	return nil
}
