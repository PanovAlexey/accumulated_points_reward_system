package service

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"
	"strconv"
)

type OrderLoader struct {
	orderNumberValidator OrderValidator
}

func NewOrderLoaderService(validator OrderValidator) *OrderLoader {
	return &OrderLoader{
		orderNumberValidator: validator,
	}
}

func (service OrderLoader) Validate(numberOrder string) (int, error) {
	if numberOrderInt, err := strconv.Atoi(numberOrder); err == nil {
		return service.orderNumberValidator.Validate(numberOrderInt)
	}

	return numberOrderInt, nil
}

	return order, nil
}

func (service OrderLoader) GetOrderByNumber(number int) (*entity.Order, error) {
	return &entity.Order{}, nil
}

func (service OrderLoader) SaveOrder(number int) (*entity.Order, error) {
	return &entity.Order{}, nil
}
