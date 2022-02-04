package service

import (
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
	return service.userRepository.CreateUser(user)
}

func (service UserRegistration) Auth(userName, password string) (domain.User, error) {
	//@ToDo
	return domain.User{}, nil
}
