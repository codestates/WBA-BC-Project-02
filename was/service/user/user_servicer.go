package user

import (
	"github.com/codestates/WBA-BC-Project-02/was/protocol/useer/response"
)

type UserServicer interface {
	RegisterUser() (string, error)

	ModifyUser(ID string) (int, error)

	FindUser(ID string) error

	DeleteUser(ID string) (int, error)

	Login() (*response.Token, error)
}
