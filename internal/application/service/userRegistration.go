package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	applicationErrors "github.com/PanovAlexey/accumulated_points_reward_system/internal/application/errors"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/repository"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	userCtx    = "userID"
	salt       = "wertyuiopasdfghjkl"
	tokenTTL   = 12 * time.Hour
	signingKey = "qweqr78939870424&(#$@"
)

// UserRegistration provides the ability to register new users, authorize existing ones and work with secret tokens.
type UserRegistration struct {
	userRepository repository.UserRepository
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func NewUserRegistrationService(userRepository repository.UserRepository) *UserRegistration {
	return &UserRegistration{
		userRepository: userRepository,
	}
}

func (service UserRegistration) GetUserCtx() string {
	return userCtx
}

func (service UserRegistration) Register(user entity.User) (entity.User, error) {
	isLoginExist, _ := service.userRepository.IsLoginExist(user.Login)

	if isLoginExist {
		return user, fmt.Errorf("%v: %w", user.Login, applicationErrors.ErrUserAlreadyExists)
	}

	passwordHash, err := service.generatePasswordHash(user.Password)

	if err != nil {
		return entity.User{}, err
	}

	user.Password = passwordHash

	return service.userRepository.CreateUser(user)
}

func (service UserRegistration) Auth(login, password string) (entity.User, error) {
	passwordHash, err := service.generatePasswordHash(password)

	if err != nil {
		return entity.User{}, err
	}

	return service.userRepository.GetUser(login, passwordHash)
}

func (service UserRegistration) GetUserByLogin(login string) (entity.User, error) {
	return service.userRepository.GetUserByLogin(login)
}

func (service UserRegistration) GenerateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
	})

	signedString, err := token.SignedString([]byte(signingKey))

	return "Bearer " + signedString, err
}

func (service UserRegistration) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func (service UserRegistration) generatePasswordHash(password string) (string, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt))), err
}
