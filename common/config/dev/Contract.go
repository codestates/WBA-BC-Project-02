package dev

import "github.com/codestates/WBA-BC-Project-02/common/ciper"

type Contract struct {
	MultiSigAddr     string
	DexAddr          string
	CreditAddr       string
	DracoAddr        string
	TigAddr          string
	ServerAddr       string
	ServerPrivateKey string
	SecondPrivateKey string
	RawURL           string
	ChannelToken     string
	BotToken         string
}

func (c *Contract) DecryptFields() error {
	fk, err := ciper.AESDecrypt(ciper.GetCipherBlock(), c.ServerPrivateKey)
	if err != nil {
		return err
	}
	c.ServerPrivateKey = fk

	ct, err := ciper.AESDecrypt(ciper.GetCipherBlock(), c.ChannelToken)
	if err != nil {
		return err
	}
	c.ChannelToken = ct

	spk, err := ciper.AESDecrypt(ciper.GetCipherBlock(), c.SecondPrivateKey)
	if err != nil {
		return err
	}
	c.SecondPrivateKey = spk

	btk, err := ciper.AESDecrypt(ciper.GetCipherBlock(), c.BotToken)
	if err != nil {
		return err
	}
	c.BotToken = btk

	return nil
}
