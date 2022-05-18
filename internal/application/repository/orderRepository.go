package repository

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"
)

type OrderRepository interface {
	CreateOrder(order entity.Order) (entity.Order, error)
	GetOrder(number string) (entity.Order, error)
}
