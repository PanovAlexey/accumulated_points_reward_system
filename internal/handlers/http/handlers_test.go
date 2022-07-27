package http

var testData = map[string]string{
	"user_login":      "test@test.test",
	"user_password":   "password",
	"user_auth_token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTM5NjA2MTAsImlhdCI6MTY1MzkxNzQxMCwidXNlcl9pZCI6NH0.UusEPIYIilvGkSf68BTNxioHs0LajSXT24OwnB7VE5A",
}

func getServerForTestAndPrepareDatabase() *httptest.Server {
	logger, userService, orderService, paymentService := getServicesForTests()
	handler := NewHandler(logger, userService, orderService, paymentService)
	server := httptest.NewServer(handler.InitRoutes())

	// get test user
	user, err := userService.GetUserByLogin(testData["user_login"])

	if err != nil {
		logger.Error(err.Error())
	}

	// delete all orders by test user
	orderService.DeleteOrdersByUserID(user.ID.Int64)

	return server
}

func getServicesForTests() (
	logging.LoggerInterface, *service.UserRegistration, *service.OrderLoader, *service.PaymentsManagement,
) {
	config := config.NewConfig()
	logger := logging.GetLogger(config)

	db, err := databases.NewPostgresDB(config.Storage.DatabaseDsn)

	if err != nil {
		logger.Fatalf("error occurred while initializing database connection: %s", err.Error())
	}

	userRegistrationRepository := repository.NewUserRepository(db)
	userRegistrationService := service.NewUserRegistrationService(userRegistrationRepository)

	orderStatusGetter := service.GetOrderStatusGetter()
	orderNumberValidator := luhn.GetLuhnAlgorithmChecker()
	orderRepository := repository.NewOrderRepository(db)
	orderLoaderService := service.NewOrderLoaderService(orderRepository, orderNumberValidator, orderStatusGetter)

	paymentRepository := repository.NewPaymentRepository(db)
	paymentManagement := service.NewPaymentManagement(paymentRepository)

	return logger, userRegistrationService, orderLoaderService, paymentManagement
}