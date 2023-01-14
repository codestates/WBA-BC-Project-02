package main

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	"github.com/joho/godotenv"
	"log"

	"github.com/codestates/WBA-BC-Project-02/common/enum"

	"github.com/codestates/WBA-BC-Project-02/common/ciper"
	"github.com/codestates/WBA-BC-Project-02/was/common/app"
	"github.com/codestates/WBA-BC-Project-02/was/common/flag"
	"github.com/codestates/WBA-BC-Project-02/was/config"
	"github.com/codestates/WBA-BC-Project-02/was/controller"
	"github.com/codestates/WBA-BC-Project-02/was/logger"
	"github.com/codestates/WBA-BC-Project-02/was/model"
	"github.com/codestates/WBA-BC-Project-02/was/router"
	"github.com/codestates/WBA-BC-Project-02/was/service"
)

var (
	App = app.NewApp()

	flags = []*flag.FlagCategory{
		flag.ServerConfigFlag,
		flag.LogConfigFlag,
		flag.DatabaseFlag,
		flag.JWTFlag,
		flag.RedisFlag,
	}

	mongoCollectionNames = []string{
		enum.UserCollectionName,
		enum.ContractCollectionName,
	}
)

func init() {
	flag.FlagsLoad(flags)
	config.LoadConfigs(flag.Flags)

	//Decrypt
	checkDevModAndLoadEnv()
	ciper.LoadCipherKey(config.ServerConfig.Mode)
	ciper.LoadCipherBlock()
	config.DecryptConfigs()

	//Redis
	if err := cache.LoadRedisClient(config.RedisConfig.DNS); err != nil {
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

func checkDevModAndLoadEnv() {
	if config.ServerConfig.Mode != enum.DevMode {
		return
	}
	if err := godotenv.Load(enum.DevModeENVFilePath); err != nil {
		log.Fatal("Error loading .env file ::", err.Error())
	}
}
