package router

import "github.com/codestates/WBA-BC-Project-02/was/router/gin"

var Route Router

func LoadRouter(mode string) {
	Route = gin.NewGinRoute(mode)
}
