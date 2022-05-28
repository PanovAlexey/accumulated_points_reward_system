package http

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *httpHandler) getOrders(c *gin.Context) {
	userCtxValue, isExist := c.Get(h.userRegistrationService.GetUserCtx())

	if isExist == false {
		responses.NewErrorResponse(c, http.StatusInternalServerError, "it is no info about user in context.")
		h.logger.Error("it is no info about user in context.")

		return
	}

	orders, err := h.orderLoaderService.GetOrdersByUserID(int64(userCtxValue.(int)))

	c.JSON(http.StatusOK, map[string]interface{}{
		"temp info": "endpoint in progress",
	})
}
