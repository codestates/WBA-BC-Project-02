package jwt

import "github.com/Hooneats/Syeong_server/common/ciper"

type JWT struct {
	AccessKey  string
	RefreshKey string
}

func (j *JWT) DecryptFields() error {
	ac, err := ciper.AESDecrypt(ciper.GetCipherBlock(), j.AccessKey)
	if err != nil {
		return err
	}
	j.AccessKey = ac

	re, err := ciper.AESDecrypt(ciper.GetCipherBlock(), j.RefreshKey)
	if err != nil {
		return err
	}
	j.RefreshKey = re
	return nil
}
