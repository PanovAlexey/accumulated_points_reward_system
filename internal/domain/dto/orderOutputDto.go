package dto

type OrderOutputDto struct {
	Number     string `json:"number" db:"number"`
	Status     string `json:"status"`
	Accrual    int    `json:"accrual"`
	UploadedAt string `json:"uploaded_at"`
}
