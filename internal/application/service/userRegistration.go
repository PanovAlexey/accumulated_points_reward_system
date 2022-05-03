package service

import (
	"errors"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/repository"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain"
)

type UserRegistration struct {
	userRepository repository.UserRepository
}

func NewUserRegistrationService(userRepository repository.UserRepository) *UserRegistration {
	return &UserRegistration{
		userRepository: userRepository,
	}
}

func (service UserRegistration) Register(user domain.User) (domain.User, error) {
	isLoginExist, _ := service.userRepository.IsLoginExist(user.Login)

	if isLoginExist {
		return user, errors.New("user already exists") // @ToDo create custom error
	}

	return service.userRepository.CreateUser(user)
}

func (service UserRegistration) Auth(login, password string) (domain.User, error) {
	//@ToDo
	return domain.User{}, nil
}
