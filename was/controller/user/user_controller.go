package user

import "github.com/gin-gonic/gin"

type UserController interface {
	CreateWallet(c *gin.Context)

	RecoverWallet(c *gin.Context)

	ReissueToken(c *gin.Context)

	GetUserSimpleInformation(c *gin.Context)

	GetUserInformation(c *gin.Context)

	IncreaseBlackIron(c *gin.Context)
}
