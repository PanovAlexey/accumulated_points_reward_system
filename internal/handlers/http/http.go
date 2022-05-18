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
	orderLoaderService      *service.OrderLoader
}

func NewHandler(
	logger logging.LoggerInterface,
	userRegistrationService *service.UserRegistration,
	orderLoaderService *service.OrderLoader,
) *httpHandler {
	return &httpHandler{
		logger:                  logger,
		userRegistrationService: userRegistrationService,
		orderLoaderService:      orderLoaderService,
	}
}

func (h *httpHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.GET("/ping", h.HandlePing)

	api := router.Group("/api")
	{
		api.Use(middleware.JSON())
		user := api.Group("/user")
		{
			user.POST("/register", h.register)
			user.POST("/login", h.login)

			user.Use(middleware.Authorization(h.userRegistrationService))
			user.POST("/orders", h.addOrder)
			user.GET("/orders", h.getOrders)
		}
	}

	return router
}
