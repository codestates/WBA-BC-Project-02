package db

import (
	"github.com/codestates/WBA-BC-Project-02/common/cipher"
)

type DB struct {
	URI        string
	DBName     string
	User       string
	PWD        string
	BackupPath string
}

func (d *DB) DecryptFields() error {
	block := cipher.GetCipherBlock()
	URI, err := cipher.AESDecrypt(block, d.URI)
	if err != nil {
		return err
	}
	d.URI = URI

	name, err := cipher.AESDecrypt(block, d.DBName)
	if err != nil {
		return err
	}
	d.DBName = name
	return nil
}
