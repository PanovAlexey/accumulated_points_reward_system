package http

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *httpHandler) register(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		h.logger.Error(err.Error())
		return
	}

	_, err := h.userRegistrationService.Register(user)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.userRegistrationService.GenerateToken(user.Login, user.Password)

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
