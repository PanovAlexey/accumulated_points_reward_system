package databases

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const Users_table_name = "users"

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
