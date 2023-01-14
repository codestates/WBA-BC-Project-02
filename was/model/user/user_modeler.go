package user

type UserModeler interface {
	PostUser() (string, error)

	UpdateUser() (int64, error)

	SelectUser(id string) error

	SelectUserByNicName(nicName string) error

	DeleteUser(id string) (int64, error)
}
