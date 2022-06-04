package dto

type WithdrawalsOutputDto struct {
	Order       float64 `json:"order" db:"order_id"`
	Sum         float64 `json:"sum" db:"sum"`
	ProcessedAt string  `json:"processed_at" db:"processed_at"`
}
