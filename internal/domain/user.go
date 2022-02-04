package domain

import "database/sql"

type User struct {
	Id       *sql.NullInt64 `json:"-" db:"id"`
	Login    string         `json:"login" db:"login" binding:"required"`
	Password string         `json:"password" binding:"required"`
}
