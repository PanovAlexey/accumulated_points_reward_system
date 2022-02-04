package repository

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUser(login, password string) (domain.User, error)
	IsLoginExist(login string) (bool, error)
}
