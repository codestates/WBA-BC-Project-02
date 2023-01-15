package user

import "github.com/codestates/WBA-BC-Project-02/common/model/entity"

type UserModeler interface {
	InsertUser(user *entity.User) error
	FindUser(address string) (*entity.User, error)
	FindUserAndPWDUpdate(address, password string) (*entity.User, error)
}
