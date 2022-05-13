package entity

import (
	"database/sql"
)

type Order struct {
	ID     sql.NullInt64 `json:"-" db:"id"`
	Number string        `json:"number" db:"number" binding:"required"`
	Status sql.NullInt64 `json:"status" db:"status" binding:"required"`
	UserID sql.NullInt64 `json:"user_id" db:"user_id" binding:"required"`
}
