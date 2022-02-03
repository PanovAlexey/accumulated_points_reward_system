package repository

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/repository"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain"
)

type userRepository struct {
}

func NewUserRepository() repository.UserRepository {
	return userRepository{}
}

func (repository userRepository) CreateUser(user domain.User) (domain.User, error) {
	//@ToDo
	return user, nil
}

func (repository userRepository) GetUser(userName, password string) (domain.User, error) {
	//@ToDo
	return domain.User{}, nil
}
