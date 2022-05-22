package service

import (
	"fmt"
	applicationErrors "github.com/PanovAlexey/accumulated_points_reward_system/internal/application/errors"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/repository"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"
	"strconv"
)

type OrderLoader struct {
	orderRepository      repository.OrderRepository
	orderNumberValidator OrderValidator
}

func NewOrderLoaderService(orderRepository repository.OrderRepository, validator OrderValidator) *OrderLoader {
	return &OrderLoader{
		orderRepository:      orderRepository,
		orderNumberValidator: validator,
	}
}

func (service OrderLoader) validate(numberOrder string) (int64, error) {
	if numberOrderInt, err := strconv.ParseInt(numberOrder, 10, 64); err == nil {
		return numberOrderInt, service.orderNumberValidator.Validate(numberOrderInt)
	} else {
		return 0, err
	}
}

func (service OrderLoader) PostOrder(number string, userID int64) (*entity.Order, error) {
	numberInt, err := service.validate(number)

	if err != nil {
		return nil, fmt.Errorf("%v: %w", number, applicationErrors.ErrorOrderNumberInvalid)
	}

	order, err := service.GetOrderByNumber(numberInt)

	if err != nil {
		return nil, err
	}

	if order == nil {
		newOrder := entity.NewOrder(numberInt, 1, userID) // @ToDO change 1 to real status ID
		order = &newOrder
	} else {
		if order.UserID.Int64 == userID {
			return order, fmt.Errorf("%v: %w", number, applicationErrors.ErrorOrderAlreadySent)
		} else {
			return order, fmt.Errorf("%v: %w", number, applicationErrors.ErrorOrderAlreadyExists)
		}
	}

	order, err = service.SaveOrder(numberInt, userID)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (service OrderLoader) GetOrderByNumber(number int64) (*entity.Order, error) {
	order, err := service.orderRepository.GetOrder(number)

	return order, err
}

func (service OrderLoader) SaveOrder(number int64, userID int64) (*entity.Order, error) {
	order, err := service.orderRepository.CreateOrder(entity.NewOrder(number, 1, userID)) // @ToDo: change 1 to real value

	return &order, err
}
