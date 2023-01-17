package db

import (
	"github.com/codestates/WBA-BC-Project-02/common/ciper"
)

type DB struct {
	URI        string
	DBName     string
	User       string
	PWD        string
	BackupPath string
}

func (d *DB) DecryptFields() error {
	block := ciper.GetCipherBlock()
	URI, err := ciper.AESDecrypt(block, d.URI)
	if err != nil {
		return err
	}
	d.URI = URI

	name, err := ciper.AESDecrypt(block, d.DBName)
	if err != nil {
		return err
	}
	d.DBName = name
	return nil
}
