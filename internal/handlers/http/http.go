package http

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/service"
	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	userRegistrationService *service.UserRegistration
}

func NewHandler(userRegistrationService *service.UserRegistration) *httpHandler {
	return &httpHandler{
		userRegistrationService: userRegistrationService,
	}
}

func (h *httpHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/ping", h.HandlePing)

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/register", h.register)
			user.POST("/login", h.login)
		}
	}

	return router
}
