package gin

import (
	"github.com/codestates/WBA-BC-Project-02/was/controller"
	"github.com/codestates/WBA-BC-Project-02/was/logger"
	v1 "github.com/codestates/WBA-BC-Project-02/was/router/gin/handler/v1"
	"github.com/codestates/WBA-BC-Project-02/was/router/gin/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

var instance *GinRoute

type GinRoute struct {
	engin *gin.Engine
}

func NewGinRoute(mode string) *GinRoute {
	if instance != nil {
		return instance
	}
	setGinMode(mode)
	instance = &GinRoute{
		engin: gin.New(),
	}
	setMiddleware()
	return instance
}

func (r *GinRoute) Handle() http.Handler {
	gr := r.engin

	version1 := gr.Group("app/v1", middleware.UserAgent())
	{
		version1.GET("/info", controller.InfoControl.GetInformation)

		u := version1.Group("/users")
		{
			v1.Users(u)
		}
	}

	return r.engin
}

func setGinMode(mode string) {
	switch mode {
	case "dev":
		logger.AppLog.Info("Start gin mod", gin.DebugMode)
		gin.SetMode(gin.DebugMode)
	case "prod":
		logger.AppLog.Info("Start gin mod", gin.ReleaseMode)
		gin.SetMode(gin.ReleaseMode)
	case "test":
		logger.AppLog.Info("Start gin mod", gin.TestMode)
		gin.SetMode(gin.TestMode)
	default:
		logger.AppLog.Info("Start gin mod", gin.DebugMode)
		gin.SetMode(gin.DebugMode)
	}
}

func setMiddleware() {
	instance.engin.Use(middleware.GinLogger())
	instance.engin.Use(middleware.GinRecovery(true))
	instance.engin.Use(middleware.CORS())
}
