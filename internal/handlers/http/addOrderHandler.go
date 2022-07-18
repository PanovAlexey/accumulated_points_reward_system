package http

import (
	"errors"
	applicationErrors "github.com/PanovAlexey/accumulated_points_reward_system/internal/application/errors"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http/responses"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// addOrder endpoint for adding order information
func (h *httpHandler) addOrder(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()

	if err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		h.logger.Warn(err.Error())

		return
	}

	userCtxValue, isExist := c.Get(h.userRegistrationService.GetUserCtx())

	if !isExist {
		responses.NewErrorResponse(c, http.StatusInternalServerError, "it is no info about user in context.")
		h.logger.Error("it is no info about user in context.")

		return
	}

	orderNumber := string(body)
	order, err := h.orderLoaderService.PostOrder(orderNumber, int64(userCtxValue.(int)))

	if errors.Is(err, applicationErrors.ErrOrderNumberInvalid) {
		responses.NewErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		h.logger.Warn(err.Error())
		return
	}

	if errors.Is(err, applicationErrors.ErrOrderAlreadySent) {
		responses.NewErrorResponse(c, http.StatusOK, err.Error())
		h.logger.Warn(err.Error(), orderNumber)
		return
	}

	if errors.Is(err, applicationErrors.ErrOrderAlreadyExists) {
		responses.NewErrorResponse(c, http.StatusConflict, err.Error())
		h.logger.Warn(err.Error(), orderNumber)
		return
	}

	h.logger.Info("order successfully created "+orderNumber, order)

	c.JSON(http.StatusAccepted, map[string]interface{}{
		"order": order,
	})
}
