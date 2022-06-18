package entity

import (
	"database/sql"
	"strconv"
	"time"
)

type Order struct {
	ID         sql.NullInt64  `json:"-" db:"id"`
	Number     string         `json:"number" db:"number" binding:"required"`
	Status     sql.NullString `json:"status" db:"status" binding:"required"`
	UserID     sql.NullInt64  `json:"user_id" db:"user_id" binding:"required"`
	UploadedAt string         `json:"uploaded_at" db:"uploaded_at" binding:"required"`
}

func NewOrder(number int64, status string, userID int64) *Order {
	var userIDNullInt sql.NullInt64
	var statusNullString sql.NullString

	err := userIDNullInt.Scan(userID)

	if err != nil {
		return nil
	}

	err = statusNullString.Scan(status)

	if err != nil {
		return nil
	}

	return &Order{
		Number:     strconv.FormatInt(number, 10),
		Status:     statusNullString,
		UserID:     userIDNullInt,
		UploadedAt: time.Now().Format(time.RFC3339),
	}
}
