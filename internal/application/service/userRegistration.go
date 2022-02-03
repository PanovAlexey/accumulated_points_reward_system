package service

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/repository"
	"os/user"
)

type UserRegistration struct {
	userRepository repository.UserRepository
}

func NewUserRegistrationService(userRepository repository.UserRepository) *UserRegistration {
	return &UserRegistration{
		userRepository: userRepository,
	}
}

func (service UserRegistration) Register(user user.User) (user.User, error) {
	//@ToDo
	return user, nil
}

func (service UserRegistration) Auth(userName, password string) (user.User, error) {
	//@ToDo
	return user.User{}, nil
}
