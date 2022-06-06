package http

import (
	"errors"
	applicationErrors "github.com/PanovAlexey/accumulated_points_reward_system/internal/application/errors"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/dto"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *httpHandler) postWithdraw(c *gin.Context) {
	userCtxValue, isExist := c.Get(h.userRegistrationService.GetUserCtx())

	if isExist == false {
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

	order, err := h.orderLoaderService.PostOrder(orderInputDto.Order, int64(userCtxValue.(int)))

	if err != nil {
		if errors.Is(err, applicationErrors.ErrorOrderNumberInvalid) {
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
