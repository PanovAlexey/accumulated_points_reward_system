package repository

import "os/user"

type UserRepository interface {
	CreateUser(user user.User) (user.User, error)
	GetUser(userName, password string) (user.User, error)
}
