package repository

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/repository"
	"os/user"
)

type userRepository struct {
}

func NewUserRepository() repository.UserRepository {
	return userRepository{}
}

func (repository userRepository) CreateUser(user user.User) (user.User, error) {
	//@ToDo
	return user, nil
}

func (repository userRepository) GetUser(userName, password string) (user.User, error) {
	//@ToDo
	return user.User{}, nil
}
