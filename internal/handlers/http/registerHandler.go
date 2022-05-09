package http

import (
	"errors"
	applicationErrors "github.com/PanovAlexey/accumulated_points_reward_system/internal/application/errors"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/entity"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *httpHandler) register(c *gin.Context) {
	var user entity.User

	if err := c.BindJSON(&user); err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		h.logger.Warn(err.Error())
		return
	}

	registeredUser, err := h.userRegistrationService.Register(user)

	if err != nil {
		if errors.Is(err, applicationErrors.ErrorAlreadyExists) {
			responses.NewErrorResponse(c, http.StatusConflict, err.Error())
			h.logger.Warn(err.Error())

			return
		}

		responses.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		h.logger.Error(err.Error())

		return
	}

	token, err := h.userRegistrationService.GenerateToken(int(registeredUser.ID.Int64))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		h.logger.Warn(err.Error())
	}

	h.logger.Info("User has successfully registered.", token, user.ID.Int64)

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
