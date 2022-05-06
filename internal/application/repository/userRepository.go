package repository

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUser(login, password string) (entity.User, error)
	IsLoginExist(login string) (bool, error)
}
