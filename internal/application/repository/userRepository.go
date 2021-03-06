package repository

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"
)

// UserRepository repository provides interface describing methods for working with users
type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUser(login, password string) (entity.User, error)
	GetUserByLogin(login string) (entity.User, error)
	IsLoginExist(login string) (bool, error)
}
