package repository

import "github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"

type PaymentRepository interface {
	GetBalance(userID int64) (float64, error)
	Create(userID, orderID int64, sum float64) (entity.Payment, error)
	GetOrderIDToPaymentMap(orderIDList []int64) (map[int64]entity.Payment, error)
	GetTotalWithdrawn(userID int64) (float64, error)
}
