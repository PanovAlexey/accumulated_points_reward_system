package service

import (
	"encoding/json"
	"fmt"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/dto"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/logging"
	"io"
	"net/http"
	"strconv"
	"time"
)

const orderProcessingFrequencyInterval = 1 * time.Second

// SynchronizationWithScoringSystemService updates order statuses, receiving up-to-date data from an external system
type SynchronizationWithScoringSystemService struct {
	orderService         OrderLoader
	paymentsManagement   PaymentsManagement
	logger               logging.LoggerInterface
	accrualSystemAddress string
}

func NewSynchronizationWithScoringSystemService(
	orderService OrderLoader,
	paymentsManagement PaymentsManagement,
	logger logging.LoggerInterface,
	accrualSystemAddress string,
) SynchronizationWithScoringSystemService {
	return SynchronizationWithScoringSystemService{
		orderService:         orderService,
		paymentsManagement:   paymentsManagement,
		logger:               logger,
		accrualSystemAddress: accrualSystemAddress,
	}
}

func (service SynchronizationWithScoringSystemService) Init() {
	for {
		service.step()
		time.Sleep(orderProcessingFrequencyInterval)
	}
}

func (service SynchronizationWithScoringSystemService) step() {
	orders, err := service.orderService.getOrdersInUnfinishedStatus()

	if err != nil {
		service.logger.Warn("error when getting orders in unfinished status from storage " + err.Error())
	}

	for _, order := range *orders {
		bonusPointsSystemResponse, err := service.GetOrderStatusInScoringSystem(order)

		if err != nil {
			service.logger.Warn("error when getting order status from external system " + err.Error())
		} else {
			service.handleOrderByStatusInScoringSystem(order, *bonusPointsSystemResponse)
		}
	}
}

func (service SynchronizationWithScoringSystemService) GetOrderStatusInScoringSystem(order entity.Order) (*dto.BonusPointsSystemResponse, error) {
	response, err := http.Get(service.accrualSystemAddress + "/api/orders/" + order.Number)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	bonusPointsSystemResponse := dto.BonusPointsSystemResponse{}
	err = json.Unmarshal(responseBody, &bonusPointsSystemResponse)

	if err != nil {
		return nil, err
	}

	return &bonusPointsSystemResponse, nil
}

func (service SynchronizationWithScoringSystemService) handleOrderByStatusInScoringSystem(
	order entity.Order, response dto.BonusPointsSystemResponse,
) {
	if response.Status == "INVALID" {
		err := service.orderService.SetInvalidStatus(order)

		if err != nil {
			service.logger.Error("error when trying to set INVALID order status to " + order.Number + err.Error())
			return
		}
	}

	if response.Status == "PROCESSING" {
		err := service.orderService.SetProcessingStatus(order)

		if err != nil {
			service.logger.Error("error when trying to set PROCESSING order status to " + order.Number + err.Error())
			return
		}
	}

	if response.Status == "PROCESSED" {
		err := service.orderService.SetProcessedStatus(order)

		if err != nil {
			service.logger.Error("error when trying to set PROCESSED order status to " + order.Number + err.Error())
			return
		}

		if response.Accrual > 0 {
			payment, err := service.paymentsManagement.Debt(order, response.Accrual)

			if err != nil {
				service.logger.Error("error when trying to accrue bonus points " + err.Error())
			} else {
				service.logger.Info("user " +
					strconv.FormatInt(order.UserID.Int64, 10) +
					" is credited with " + fmt.Sprintf("%.2f", payment.Sum) + " points .",
				)
			}
		}
	}
}
