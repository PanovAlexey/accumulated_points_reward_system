package http

import (
	"errors"
	applicationErrors "github.com/PanovAlexey/accumulated_points_reward_system/internal/application/errors"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http/responses"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (h *httpHandler) addOrder(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		h.logger.Warn(err.Error())

		return
	}

	orderNumber := string(body)
	order, err := h.orderLoaderService.PostOrder(orderNumber)

	if errors.Is(err, applicationErrors.ErrorOrderNumberInvalid) {
		responses.NewErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		h.logger.Warn(err.Error())
	}

	if errors.Is(err, applicationErrors.ErrorOrderAlreadySent) {
		responses.NewErrorResponse(c, http.StatusOK, err.Error())
		h.logger.Warn(err.Error(), orderNumber)
	}

	if errors.Is(err, applicationErrors.ErrorOrderAlreadyExists) {
		responses.NewErrorResponse(c, http.StatusConflict, err.Error())
		h.logger.Warn(err.Error(), orderNumber)
	}

	c.JSON(http.StatusAccepted, map[string]interface{}{
		"order": order,
	})
}
