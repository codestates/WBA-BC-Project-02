package dev

import "github.com/codestates/WBA-BC-Project-02/common/cipher"

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
	fk, err := cipher.AESDecrypt(cipher.GetCipherBlock(), c.ServerPrivateKey)
	if err != nil {
		return err
	}
	c.ServerPrivateKey = fk

	ct, err := cipher.AESDecrypt(cipher.GetCipherBlock(), c.ChannelToken)
	if err != nil {
		return err
	}
	c.ChannelToken = ct

	spk, err := cipher.AESDecrypt(cipher.GetCipherBlock(), c.SecondPrivateKey)
	if err != nil {
		return err
	}
	c.SecondPrivateKey = spk

	btk, err := cipher.AESDecrypt(cipher.GetCipherBlock(), c.BotToken)
	if err != nil {
		return err
	}
	c.BotToken = btk

	return nil
}
