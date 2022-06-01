package http

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/dto"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *httpHandler) getBalance(c *gin.Context) {
	userCtxValue, isExist := c.Get(h.userRegistrationService.GetUserCtx())

	if isExist == false {
		responses.NewErrorResponse(c, http.StatusInternalServerError, "it is no info about user in context.")
		h.logger.Error("it is no info about user in context.")

		return
	}

	currentBalance, err := h.paymentManagement.GetUserBalance(int64(userCtxValue.(int)))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		h.logger.Error(err.Error())

		return
	}

	totalWithdrawn, err := h.paymentManagement.GetTotalWithdrawn(int64(userCtxValue.(int)))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		h.logger.Error(err.Error())

		return
	}

	balanceOutput := dto.BalanceOutputDto{
		Current:   currentBalance,
		Withdrawn: totalWithdrawn,
	}

	c.JSON(http.StatusOK, balanceOutput)
}
