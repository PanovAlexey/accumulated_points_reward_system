package repository

import (
	"database/sql"
	"errors"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/repository"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/databases"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
)

type paymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) repository.PaymentRepository {
	return paymentRepository{db: db}
}

func (repository paymentRepository) GetBalance(userID int64) (float64, error) {
	var balance sql.NullFloat64

	err := repository.db.QueryRow(
		"SELECT SUM(sum) FROM "+databases.PaymentsTableNameConst+" WHERE user_id = ($1) GROUP BY user_id",
		userID,
	).Scan(&balance)

	if err != nil {
		return 0, err
	}

	return balance.Float64, nil
}

func (repository paymentRepository) Create(userID, orderID int64, sum float64) (entity.Payment, error) {
	var payment entity.Payment

	payment.Order.Scan(orderID)
	payment.UserID.Scan(userID)
	payment.Sum = sum
	payment.ProcessedAt = time.Now().Format(time.RFC3339)

	rows, err := repository.db.NamedQuery(
		`INSERT INTO `+
			databases.PaymentsTableNameConst+
			` (user_id, order_id, processed_at, sum) VALUES (:user_id, :order_id, :processed_at, :sum) RETURNING id`,
		payment,
	)

	if err == nil {
		var insertID int
		rows.Next()
		rows.Scan(&insertID)

		payment.ID.Scan(insertID)

		if rows.Err() != nil {
			err = rows.Err()
		}
	} else {
		if rows.Err() != nil {
			err = errors.New(err.Error() + rows.Err().Error())
		}
	}

	return payment, err
}

func (repository paymentRepository) GetOrderIDToPaymentMap(orderIDList []int64) (map[int64]entity.Payment, error) {
	var ordersIDString string
	orderIDToPaymentMap := make(map[int64]entity.Payment)

	for i := 0; i < len(orderIDList); i++ {
		ordersIDString = ordersIDString + " order_id = " + strconv.FormatInt(orderIDList[i], 10)

		if i+1 < len(orderIDList) {
			ordersIDString = ordersIDString + " OR"
		}
	}

	var payments []entity.Payment
	err := repository.db.Select(
		&payments,
		"SELECT id, user_id, order_id, processed_at, sum FROM "+databases.PaymentsTableNameConst+" WHERE "+ordersIDString,
	)

	if err != nil {
		return nil, err
	}

	for _, payment := range payments {
		orderIDToPaymentMap[payment.Order.Int64] = payment
	}

	return orderIDToPaymentMap, nil
}
