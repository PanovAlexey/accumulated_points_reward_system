package entity

import (
	"database/sql"
)

type Payment struct {
	ID          sql.NullInt64 `json:"-" db:"id"`
	Order       sql.NullInt64 `json:"order_id" db:"order_id" binding:"required"`
	Sum         float64       `json:"sum"`
	ProcessedAt string        `json:"processed_at" db:"processed_at"`
	UserID      sql.NullInt64 `json:"user_id" db:"user_id" binding:"required"`
}

/*
func (w Withdrawal) MarshalJSON() ([]byte, error) {
	type WithdrawalAlias Withdrawal

	aliasValue := struct {
		WithdrawalAlias
		ProcessedAt string `json:"processed_at"`
	}{
		WithdrawalAlias: WithdrawalAlias(w),
		ProcessedAt:     w.ProcessedAt.Format(time.RFC3339),
	}

	return json.Marshal(aliasValue)
}
*/
