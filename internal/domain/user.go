package domain

type User struct {
	Id       int    `json:"-" db:"id"`
	Login    string         `json:"login" db:"login" binding:"required"`
	Password string         `json:"password" binding:"required"`
}
