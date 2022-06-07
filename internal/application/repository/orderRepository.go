package repository

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"
)

type OrderRepository interface {
	CreateOrder(order entity.Order) (entity.Order, error)
	GetOrder(number int64) (*entity.Order, error)
	GetOrdersByUserID(userID int64) (*[]entity.Order, error)
	SetOrderStatusID(orderID int64, statusID int) error
	GetOrdersByStatusesID([]int) (*[]entity.Order, error)
}
