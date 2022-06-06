package dto

type OrderInputDto struct {
	Order string  `json:"order"`
	Sum   float64 `json:"sum"`
}
