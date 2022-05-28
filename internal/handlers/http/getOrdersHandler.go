package http

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/dto"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *httpHandler) getOrders(c *gin.Context) {
	userCtxValue, isExist := c.Get(h.userRegistrationService.GetUserCtx())

	if isExist == false {
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

	var orderOutputs []dto.OrderOutputDto

	for _, order := range *orders {
		orderOutputs = append(
			orderOutputs,
			dto.OrderOutputDto{
				Number:     order.Number,
				Status:     h.orderLoaderService.GetStatusNameByID(int(order.Status.Int64)),
				Accrual:    500, // @ToDo: show real value
				UploadedAt: time.Now().Format(time.RFC3339),
			},
		)
	}

	c.JSON(http.StatusOK, orderOutputs)
}
