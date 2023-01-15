package v1

import "github.com/gin-gonic/gin"

// AuthContracts ("app/v1/auth/contracts")
func AuthContracts(authUrl *gin.RouterGroup) {
	authUrl.GET("")
}
