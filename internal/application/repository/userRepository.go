package repository

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUser(userName, password string) (domain.User, error)
}
