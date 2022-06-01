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
	paymentManagement       *service.PaymentsManagement
}

func NewHandler(
	logger logging.LoggerInterface,
	userRegistrationService *service.UserRegistration,
	orderLoaderService *service.OrderLoader,
	paymentManagement *service.PaymentsManagement,
) *httpHandler {
	return &httpHandler{
		logger:                  logger,
		userRegistrationService: userRegistrationService,
		orderLoaderService:      orderLoaderService,
		paymentManagement:       paymentManagement,
	}
}

func (h *httpHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.GET("/ping", h.HandlePing)

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/register", middleware.RequestJSON(), middleware.ResponseJSON(), h.register)
			user.POST("/login", middleware.RequestJSON(), middleware.ResponseJSON(), h.login)

			user.POST("/orders", middleware.Authorization(h.userRegistrationService), middleware.ResponseJSON(), h.addOrder)
			user.GET("/orders", middleware.Authorization(h.userRegistrationService), middleware.ResponseJSON(), h.getOrders)

			user.GET("/balance", middleware.Authorization(h.userRegistrationService), middleware.ResponseJSON(), h.getBalance)
		}
	}

	return router
}
