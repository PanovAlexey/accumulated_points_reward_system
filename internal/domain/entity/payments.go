package entity

import (
	"database/sql"
	"time"
)

type Payments struct {
	ID          sql.NullInt64 `json:"-" db:"id"`
	Order       string        `json:"order" db:"order" binding:"required"`
	Sum         float64       `json:"sum"`
	ProcessedAt time.Time     `json:"processed_at"`
	UserID      sql.NullInt64 `json:"user_id" db:"user_id" binding:"required"`
}
