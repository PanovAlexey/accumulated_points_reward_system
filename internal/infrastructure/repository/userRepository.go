package repository

import (
	"errors"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/repository"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/databases"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return userRepository{db: db}
}

func (repository userRepository) CreateUser(user entity.User) (entity.User, error) {
	rows, err := repository.db.NamedQuery(
		`INSERT INTO `+databases.UsersTableNameConst+` (login, password) VALUES (:login, :password) RETURNING id`,
		user,
	)

	if err == nil {
		var insertID int64
		rows.Next()
		err = rows.Scan(&insertID)

		user.ID.Int64 = insertID

		if err != nil {
			return user, err
		}

		if rows.Err() != nil {
			err = rows.Err()
		}
	} else {
		if rows.Err() != nil {
			err = errors.New(err.Error() + rows.Err().Error())
		}
	}

	return user, err
}

func (repository userRepository) IsLoginExist(login string) (bool, error) {
	user := entity.User{}
	err := repository.db.Get(
		&user,
		"SELECT * FROM "+databases.UsersTableNameConst+" WHERE login = $1 LIMIT 1",
		login,
	)

	return user.ID.Valid, err
}

func (repository userRepository) GetUser(login, passwordHash string) (entity.User, error) {
	var user entity.User
	err := repository.db.Get(
		&user,
		"SELECT * FROM "+databases.UsersTableNameConst+" WHERE login = $1 and password = $2 LIMIT 1",
		login,
		passwordHash,
	)

	return user, err
}

func (repository userRepository) GetUserByLogin(login string) (entity.User, error) {
	var user entity.User
	err := repository.db.Get(
		&user,
		"SELECT * FROM "+databases.UsersTableNameConst+" WHERE login = $1 LIMIT 1",
		login,
	)

	return user, err
}
