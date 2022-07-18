package http

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/dto"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// getOrders endpoint for returning information about user orders
func (h *httpHandler) getOrders(c *gin.Context) {
	userCtxValue, isExist := c.Get(h.userRegistrationService.GetUserCtx())

	if !isExist {
		responses.NewErrorResponse(c, http.StatusInternalServerError, "it is no info about user in context.")
		h.logger.Error("it is no info about user in context.")

		return
	}

	orders, err := h.orderLoaderService.GetOrdersByUserID(int64(userCtxValue.(int)))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		h.logger.Error(err.Error())

		return
	}

	if orders == nil {
		responses.NewErrorResponse(c, http.StatusNoContent, "no orders")
		h.logger.Error("no orders", userCtxValue)

		return
	}

	orderIDToPaymentMap, err := h.paymentManagement.GetOrderIDToPaymentMap(*orders)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		h.logger.Error(err.Error())

		return
	}

	var orderOutputs []dto.OrderOutputDto

	for _, order := range *orders {
		orderOutput := dto.OrderOutputDto{
			Number:     order.Number,
			Status:     order.Status.String,
			UploadedAt: time.Now().Format(time.RFC3339),
		}

		if orderIDToPaymentMap[order.ID.Int64].Sum > 0 {
			orderOutput.Accrual = orderIDToPaymentMap[order.ID.Int64].Sum
		}

		orderOutputs = append(
			orderOutputs,
			orderOutput,
		)
	}

	c.JSON(http.StatusOK, orderOutputs)
}
