package repository

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/repository"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/databases"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return userRepository{db: db}
}

func (repository userRepository) CreateUser(user domain.User) (domain.User, error) {
	_, err := repository.db.NamedExec(
		`INSERT INTO `+databases.Users_table_name+` (login, password) VALUES (:login, :password)`,
		user,
	)

	return user, err
}

func (repository userRepository) IsLoginExist(login string) (bool, error) {
	user := domain.User{}
	err := repository.db.Get(
		&user,
		"SELECT * FROM "+databases.Users_table_name+" WHERE login = $1 LIMIT 1",
		login,
	)

	return user.Id.Valid, err
}

func (repository userRepository) GetUser(login, password string) (domain.User, error) {
	user := domain.User{}
	err := repository.db.Get(
		&user,
		"SELECT * FROM "+databases.Users_table_name+" WHERE login = $1 and password = $2 LIMIT 1",
		"login",
		"password",
	)

	return domain.User{}, err
}
