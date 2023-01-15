package user

import "github.com/gin-gonic/gin"

type UserController interface {
	CreateWallet(c *gin.Context)

	RecoverWallet(c *gin.Context)
}
