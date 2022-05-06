package repository

import (
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
		`INSERT INTO `+databases.Users_table_name+` (login, password) VALUES (:login, :password) RETURNING id`,
		user,
	)

	if err == nil {
		var insertID int
		rows.Next()
		rows.Scan(&insertID)

		user.Id.Scan(insertID)
	}

	return user, err
}

func (repository userRepository) IsLoginExist(login string) (bool, error) {
	user := entity.User{}
	err := repository.db.Get(
		&user,
		"SELECT * FROM "+databases.Users_table_name+" WHERE login = $1 LIMIT 1",
		login,
	)

	return user.Id.Valid, err
}

func (repository userRepository) GetUser(login, passwordHash string) (entity.User, error) {
	var user entity.User
	err := repository.db.Get(
		&user,
		"SELECT * FROM "+databases.Users_table_name+" WHERE login = $1 and password = $2 LIMIT 1",
		login,
		passwordHash,
	)

	return user, err
}
