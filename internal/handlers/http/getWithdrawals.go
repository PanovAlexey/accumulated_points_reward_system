package http

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

// endpoint for returning information about withdrawals from the user's account
func (h *httpHandler) getWithdrawals(c *gin.Context) {
	userCtxValue, isExist := c.Get(h.userRegistrationService.GetUserCtx())

	if !isExist {
		responses.NewErrorResponse(c, http.StatusInternalServerError, "it is no info about user in context.")
		h.logger.Error("it is no info about user in context.")

		return
	}

	withdrawals, err := h.paymentManagement.GetWithdrawals(int64(userCtxValue.(int)))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		h.logger.Error(err.Error())

		return
	}

	if len(withdrawals) == 0 {
		responses.NewErrorResponse(c, http.StatusNoContent, "No withdrawals")

		return
	}

	c.JSON(http.StatusOK, withdrawals)
}
