package http

import (
	"errors"
	"fmt"
	applicationErrors "github.com/PanovAlexey/accumulated_points_reward_system/internal/application/errors"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/dto"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// postWithdraw endpoint for withdrawing bonuses on account of the order.
func (h *httpHandler) postWithdraw(c *gin.Context) {
	userCtxValue, isExist := c.Get(h.userRegistrationService.GetUserCtx())

	if !isExist {
		responses.NewErrorResponse(c, http.StatusInternalServerError, "it is no info about user in context.")
		h.logger.Error("it is no info about user in context.")

		return
	}

	var orderInputDto dto.OrderInputDto

	if err := c.BindJSON(&orderInputDto); err != nil {
		responses.NewErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		h.logger.Warn(err.Error())
		return
	}

	currentBalance, err := h.paymentManagement.GetUserBalance(int64(userCtxValue.(int)))

	if err != nil {
		errorMessage := "error getting balance for user " + strconv.Itoa(userCtxValue.(int))
		responses.NewErrorResponse(
			c,
			http.StatusInternalServerError,
			errorMessage,
		)
		h.logger.Error(errorMessage)
		return
	}

	if currentBalance < orderInputDto.Sum {
		errorMessage := "Error: user " +
			strconv.Itoa(userCtxValue.(int)) +
			" not enough money. It is impossible to withdraw " +
			fmt.Sprintf("%.2f", orderInputDto.Sum) +
			". The balance is " + fmt.Sprintf("%.2f", currentBalance) + "."

		responses.NewErrorResponse(c, http.StatusPaymentRequired, errorMessage)
		h.logger.Warn(errorMessage)
		return
	}

	order, err := h.orderLoaderService.PostOrder(orderInputDto.Order, int64(userCtxValue.(int)))

	if err != nil {
		if errors.Is(err, applicationErrors.ErrOrderNumberInvalid) {
			responses.NewErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
			h.logger.Warn(err.Error())
			return
		} else {
			responses.NewErrorResponse(
				c,
				http.StatusInternalServerError,
				"error creating order for withdraw. Order number "+orderInputDto.Order,
			)
			h.logger.Error("error creating order for withdraw. Order number " + orderInputDto.Order)

			return
		}
	}

	payment, err := h.paymentManagement.Credit(*order, orderInputDto.Sum)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		h.logger.Error(err.Error())

		return
	}

	c.JSON(http.StatusOK, payment)
}
