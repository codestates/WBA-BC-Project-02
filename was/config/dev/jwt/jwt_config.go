package jwt

import "github.com/codestates/WBA-BC-Project-02/common/cipher"

type JWT struct {
	AccessKey  string
	RefreshKey string
}

func (j *JWT) DecryptFields() error {
	ac, err := cipher.AESDecrypt(cipher.GetCipherBlock(), j.AccessKey)
	if err != nil {
		return err
	}
	j.AccessKey = ac

	re, err := cipher.AESDecrypt(cipher.GetCipherBlock(), j.RefreshKey)
	if err != nil {
		return err
	}
	j.RefreshKey = re
	return nil
}
