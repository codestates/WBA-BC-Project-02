package config

import (
	configContract "github.com/codestates/WBA-BC-Project-02/common/config/dev"
	"github.com/codestates/WBA-BC-Project-02/was/common/flag"
	configDB "github.com/codestates/WBA-BC-Project-02/was/config/dev/db"
	configJWT "github.com/codestates/WBA-BC-Project-02/was/config/dev/jwt"
	"github.com/codestates/WBA-BC-Project-02/was/config/dev/redis"
	configLog "github.com/codestates/WBA-BC-Project-02/was/config/log"
	configServer "github.com/codestates/WBA-BC-Project-02/was/config/server"
	"github.com/naoina/toml"
	"log"
	"os"
)

var DBConfig *configDB.DB
var RedisConfig *redis.Redis
var JWTConfig *configJWT.JWT
var LogConfig *configLog.Log
var ServerConfig *configServer.Server
var ContractConfig *configContract.Contract

func LoadConfigs(pathMap map[string]*string) {
	dpath := pathMap[flag.DatabaseFlag.Name]
	DBConfig = NewConfig(*dpath, &configDB.DB{})

	rpath := pathMap[flag.RedisFlag.Name]
	RedisConfig = NewConfig(*rpath, &redis.Redis{})

	jpath := pathMap[flag.JWTFlag.Name]
	JWTConfig = NewConfig(*jpath, &configJWT.JWT{})

	lpath := pathMap[flag.LogConfigFlag.Name]
	LogConfig = NewConfig(*lpath, &configLog.Log{})

	spath := pathMap[flag.ServerConfigFlag.Name]
	ServerConfig = NewConfig(*spath, &configServer.Server{})

	cpath := pathMap[flag.ContractFlag.Name]
	ContractConfig = NewConfig(*cpath, &configContract.Contract{})

}

func NewConfig[T any](path string, t *T) *T {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal("start app... does not exists config file in ", path)
	}

	if err := toml.NewDecoder(file).Decode(t); err != nil {
		log.Fatal("start app... toml decode, fail ::", err.Error())
	}
	return t
}

func DecryptConfigs() {
	if err := JWTConfig.DecryptFields(); err != nil {
		log.Fatal(err)
	}
	if err := DBConfig.DecryptFields(); err != nil {
		log.Fatal(err)
	}
	if err := RedisConfig.DecryptFields(); err != nil {
		log.Fatal(err)
	}
	if err := ContractConfig.DecryptFields(); err != nil {
		log.Fatal(err)
	}
}
