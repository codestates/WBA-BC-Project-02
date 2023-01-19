package user

import (
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache/login"
)

type UserModeler interface {
	InsertUser(user *entity.User) error

	FindUser(address string) (*entity.User, error)

	FindUserNonTx(address string) (*entity.User, error)

	FindUserAndPWDUpdate(address, password string) (*entity.User, error)

	FindUserAndIncreaseIron(address string) (*entity.User, error)

	FindUserAndSetIron(address string, blackIron int) (*entity.User, error)

	FindUserAndSet(loginInfo *login.Information) (*entity.User, error)
}
