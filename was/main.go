package main

import (
	"github.com/codestates/WBA-BC-Project-02/common/cipher"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	"github.com/codestates/WBA-BC-Project-02/was/controller"
	"github.com/codestates/WBA-BC-Project-02/was/logger"
	"github.com/codestates/WBA-BC-Project-02/was/model"
	"github.com/codestates/WBA-BC-Project-02/was/router"
	"github.com/codestates/WBA-BC-Project-02/was/service"
	"github.com/joho/godotenv"
	"log"

	"github.com/codestates/WBA-BC-Project-02/common/enum"

	"github.com/codestates/WBA-BC-Project-02/was/common/app"
	"github.com/codestates/WBA-BC-Project-02/was/common/flag"
	"github.com/codestates/WBA-BC-Project-02/was/config"
)

var (
	App = app.NewApp()

	flags = []*flag.FlagCategory{
		flag.ServerConfigFlag,
		flag.LogConfigFlag,
		flag.DatabaseFlag,
		flag.JWTFlag,
		flag.RedisFlag,
		flag.ContractFlag,
	}

	mongoCollectionNames = []string{
		enum.UserCollectionName,
		enum.ContractCollectionName,
		enum.DexContractName,
	}
)

func init() {
	flag.FlagsLoad(flags)
	config.LoadConfigs(flag.Flags)

	//Decrypt
	CheckDevModAndLoadEnv()
	cipher.LoadCipherKey(config.ServerConfig.Mode)
	cipher.LoadCipherBlock()
	config.DecryptConfigs()
	//Redis
	if err := cache.LoadRedis(config.RedisConfig.DNS); err != nil {
		log.Fatal(err)
	}

	//logger
	logger.LoadLogger(config.LogConfig)

	// model
	if err := model.LoadMongoModel(config.DBConfig.URI, config.DBConfig.DBName, mongoCollectionNames); err != nil {
		log.Fatal(err)
	}
	model.CreateIndexesInModels()
	model.InjectModelsMongoDependency(model.MongoCollections)
	model.InsertInitContracts()

	// service
	service.InjectServicesDependency()

	// controller
	controller.InjectControllerDependency()

	// router
	router.LoadRouter(config.ServerConfig.Mode)

}

func main() {
	App.Run()
}

func CheckDevModAndLoadEnv() {
	if config.ServerConfig.Mode != enum.DevMode {
		return
	}
	if err := godotenv.Load(enum.DevModeENVFilePath); err != nil {
		log.Fatal("Error loading .env file ::", err.Error())
	}
}
