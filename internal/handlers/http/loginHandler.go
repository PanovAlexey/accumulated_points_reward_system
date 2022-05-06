package http

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/domain/dto"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *httpHandler) login(c *gin.Context) {
	var userAuth dto.UserAuthDto

	if err := c.BindJSON(&userAuth); err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.userRegistrationService.Auth(userAuth.Login, userAuth.Password)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := h.userRegistrationService.GenerateToken(int(user.Id.Int64))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
