package repository

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"
)

// OrderRepository repository provides interface describing methods for working with orders
type OrderRepository interface {
	CreateOrder(order entity.Order) (entity.Order, error)
	GetOrder(number int64) (*entity.Order, error)
	GetOrdersByUserID(userID int64) (*[]entity.Order, error)
	DeleteOrdersByUserID(userID int64) error
	SetOrderStatus(orderID int64, status string) error
	GetOrdersByStatuses([]string) (*[]entity.Order, error)
}
