package router

import (
	"github.com/codestates/WBA-BC-Project-02/was/logger"
	v1 "github.com/codestates/WBA-BC-Project-02/was/router/handler/v1"
	"github.com/codestates/WBA-BC-Project-02/was/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

var instance *GinRoute

type GinRoute struct {
	engin *gin.Engine
}

func LoadRouter(mode string) {
	Route = newGinRoute(mode)
}

func newGinRoute(mode string) *GinRoute {
	if instance != nil {
		return instance
	}
	setMode(mode)
	instance = &GinRoute{
		engin: newEngine(),
	}
	return instance
}

func (r *GinRoute) Handle() http.Handler {
	gr := r.engin

	v1.GinHandle(gr)

	return r.engin
}

// newEngine generate gin engin and global middleware setting
func newEngine() *gin.Engine {
	grt := gin.New()
	setMiddleware(grt)
	return grt
}

func setMode(mode string) {
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

func setMiddleware(grt *gin.Engine) {
	grt.Use(middleware.GinLogger())
	grt.Use(middleware.GinRecovery(true))
	grt.Use(middleware.CORS())
}
