package dto

// BonusPointsSystemResponse structure containing the response of the bonus calculation system
type BonusPointsSystemResponse struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual"`
}
