package config

import (
	"os"
	"github.com/naoina/toml"
)

type Config struct {
	Info struct {
		FirstPk string
		SecondPk string
		MultisigAddr string
		BotToken string
		ChannelID string
	}
}

func GetConfig(fpath string) *Config {
	c := new(Config)
	if file, err := os.Open(fpath); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}
