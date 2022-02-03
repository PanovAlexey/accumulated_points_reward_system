package http

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/service"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http/middleware"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/logging"
	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	logger                  logging.LoggerInterface
	userRegistrationService *service.UserRegistration
}

func NewHandler(logger logging.LoggerInterface, userRegistrationService *service.UserRegistration) *httpHandler {
	return &httpHandler{
		logger:                  logger,
		userRegistrationService: userRegistrationService,
	}
}

func (h *httpHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/ping", h.HandlePing)

	api := router.Group("/api")
	{
		api.Use(middleware.JSON())
		user := api.Group("/user")
		{
			user.POST("/register", h.register)
			user.POST("/login", h.login)
		}
	}

	return router
}
