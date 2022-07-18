package dto

// OrderInputDto structure to get information about a new order
type OrderInputDto struct {
	Order string  `json:"order"`
	Sum   float64 `json:"sum"`
}
