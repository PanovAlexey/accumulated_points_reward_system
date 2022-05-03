package repository

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/repository"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain"
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

	return user, nil
}

func (repository userRepository) GetUser(login, password string) (domain.User, error) {
	//@ToDo
	return domain.User{}, nil
}
