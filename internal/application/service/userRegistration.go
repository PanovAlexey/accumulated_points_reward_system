package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/repository"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	salt       = "wertyuiopasdfghjkl"
	tokenTTL   = 12 * time.Hour
	signingKey = "qweqr78939870424&(#$@"
)

type UserRegistration struct {
	userRepository repository.UserRepository
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
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

	user.Password = service.generatePasswordHash(user.Password)

	return service.userRepository.CreateUser(user)
}

func (service UserRegistration) Auth(login, password string) (domain.User, error) {
	//@ToDo
	return domain.User{}, nil
}
