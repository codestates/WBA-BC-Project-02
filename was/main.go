package main

import (
	"log"

	"github.com/codestates/WBA-BC-Project-02/common/enum"

	"github.com/codestates/WBA-BC-Project-02/common/ciper"
	commonModel "github.com/codestates/WBA-BC-Project-02/common/model"
	"github.com/codestates/WBA-BC-Project-02/was/common/app"
	"github.com/codestates/WBA-BC-Project-02/was/common/flag"
	"github.com/codestates/WBA-BC-Project-02/was/common/redis"
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
	ciper.LoadCipherKey(config.ServerConfig.Mode)
	ciper.LoadCipherBlock()
	config.DecryptConfigs()

	//Redis
	if err := redis.LoadRedisClient(config.RedisConfig.DNS); err != nil {
		log.Fatal(err)
	}

	//logger
	logger.LoadLogger(config.LogConfig)

	// model
	if err := model.LoadMongoModel(config.DBConfig.URI, config.DBConfig.DBName, mongoCollectionNames); err != nil {
		log.Fatal(err)
	}
	model.CreateIndexesInModels()
	model.InjectModelsMongoDependency(commonModel.MongoCollections)

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
