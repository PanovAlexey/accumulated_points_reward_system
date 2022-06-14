package repository

import (
	"database/sql"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/repository"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/databases"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strconv"
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
			` (user_id, number, status, uploaded_at) VALUES (:user_id, :number, :status, :uploaded_at) RETURNING id`,
		order,
	)

	if err == nil {
		var insertID int
		rows.Next()
		err = rows.Scan(&insertID)

		if err != nil {
			return order, err
		}

		err = order.ID.Scan(insertID)

		if rows.Err() != nil {
			err = rows.Err()
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

func (repository orderRepository) GetOrdersByUserID(userID int64) (*[]entity.Order, error) {
	var orders []entity.Order

	err := repository.db.Select(
		&orders,
		"SELECT * FROM "+databases.OrdersTableNameConst+" WHERE user_id = $1",
		userID,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &orders, nil
}

func (repository orderRepository) SetOrderStatusID(orderID int64, statusID int) error {
	_, err := repository.db.Exec(
		"UPDATE " + databases.OrdersTableNameConst +
			" SET status=" + strconv.Itoa(statusID) +
			" WHERE id=" + strconv.FormatInt(orderID, 10),
	)

	return err
}

func (repository orderRepository) GetOrdersByStatusesID(statusesID []int) (*[]entity.Order, error) {
	var orders []entity.Order

	if len(statusesID) == 0 {
		return &orders, nil
	}

	stmt, err := repository.db.Prepare("SELECT * FROM " + databases.OrdersTableNameConst + " WHERE status = ANY($1)")

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(pq.Array(statusesID))

	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	defer rows.Close()

	for rows.Next() {
		var o entity.Order
		err = rows.Scan(&o.ID, &o.Status, &o.Number, &o.UserID, &o.UploadedAt)

		if err != nil {
			return &orders, err
		}

		orders = append(orders, o)
	}

	return &orders, nil
}
