package dto

// UserAuthDto structure for passing user data required for authorization
type UserAuthDto struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
