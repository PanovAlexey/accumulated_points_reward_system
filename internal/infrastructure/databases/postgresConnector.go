package databases

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const UsersTableNameConst = "users"
const OrdersTableNameConst = "orders"
const OrderStatusTableNameConst = "order_status"
const PaymentsTableNameConst = "payments"

func NewPostgresDB(databaseDsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", databaseDsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
