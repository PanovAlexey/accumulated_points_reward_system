package repository

import (
	"database/sql"
	"errors"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/repository"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/databases"
	"github.com/jmoiron/sqlx"
)

type orderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) repository.OrderRepository {
	return orderRepository{db: db}
}

func (repository orderRepository) CreateOrder(order entity.Order) (entity.Order, error) {
	rows, err := repository.db.NamedQuery(
		`INSERT INTO `+
			databases.OrdersTableNameConst+
			` (user_id, number, status) VALUES (:user_id, :number, :status) RETURNING id`,
		order,
	)

	if err == nil {
		var insertID int
		rows.Next()
		rows.Scan(&insertID)

		order.ID.Scan(insertID)

		if rows.Err() != nil {
			err = rows.Err()
		}
	} else {
		if rows.Err() != nil {
			err = errors.New(err.Error() + rows.Err().Error())
		}
	}

	return order, err
}

func (repository orderRepository) GetOrder(number int64) (*entity.Order, error) {
	var order entity.Order
	err := repository.db.Get(
		&order,
		"SELECT * FROM "+databases.OrdersTableNameConst+" WHERE number = $1 LIMIT 1",
		number,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &order, err
}
