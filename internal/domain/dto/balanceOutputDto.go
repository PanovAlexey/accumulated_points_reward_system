package dto

// BalanceOutputDto structure for displaying balance in API
type BalanceOutputDto struct {
	Current   float64 `json:"current"`
	Withdrawn float64 `json:"withdrawn"`
}
