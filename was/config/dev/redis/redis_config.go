package redis

import "github.com/codestates/WBA-BC-Project-02/common/cipher"

type Redis struct {
	DNS string
	PWD string
}

func (r *Redis) DecryptFields() error {
	DNS, err := cipher.AESDecrypt(cipher.GetCipherBlock(), r.DNS)
	if err != nil {
		return err
	}
	r.DNS = DNS
	return nil
}
