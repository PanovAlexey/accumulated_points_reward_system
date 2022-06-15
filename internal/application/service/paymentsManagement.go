package service

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/repository"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/dto"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"
)

type PaymentsManagement struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentManagement(paymentRepository repository.PaymentRepository) *PaymentsManagement {
	return &PaymentsManagement{
		paymentRepository: paymentRepository,
	}
}

func (service PaymentsManagement) Debt(order entity.Order, sum float64) (entity.Payment, error) {
	var payment entity.Payment

	err := payment.Order.Scan(order.ID.Int64)

	if err != nil {
		return payment, err
	}

	err = payment.UserID.Scan(order.UserID.Int64)

	if err != nil {
		return payment, err
	}

	payment.Sum = sum

	return service.paymentRepository.Create(payment)
}

func (service PaymentsManagement) Credit(order entity.Order, sum float64) (entity.Payment, error) {
	if sum > 0 {
		sum = sum * (-1)
	}

	var payment entity.Payment

	err := payment.Order.Scan(order.ID.Int64)

	if err != nil {
		return payment, err
	}

	err = payment.UserID.Scan(order.UserID.Int64)

	if err != nil {
		return payment, err
	}

	payment.Sum = sum

	return service.paymentRepository.Create(payment)
}

func (service PaymentsManagement) GetOrderIDToPaymentMap(orders []entity.Order) (map[int64]entity.Payment, error) {
	orderIDList := make([]int64, 0)

	for _, order := range orders {
		orderIDList = append(orderIDList, order.ID.Int64)
	}

	return service.paymentRepository.GetOrderIDToPaymentMap(orderIDList)
}

func (service PaymentsManagement) GetUserBalance(userID int64) (float64, error) {
	return service.paymentRepository.GetBalance(userID)
}

func (service PaymentsManagement) GetTotalWithdrawn(userID int64) (float64, error) {
	totalWithdrawn, err := service.paymentRepository.GetTotalWithdrawn(userID)

	if totalWithdrawn < 0 {
		totalWithdrawn = totalWithdrawn * (-1)
	}

	return totalWithdrawn, err
}

func (service PaymentsManagement) GetWithdrawals(userID int64) ([]dto.WithdrawalsOutputDto, error) {
	return service.paymentRepository.GetWithdrawnPayments(userID)
}
