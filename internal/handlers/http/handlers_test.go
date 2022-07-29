package http

import (
	"bytes"
	"github.com/PanovAlexey/accumulated_points_reward_system/config"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/service"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/service/luhn"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/databases"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/logging"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testData = map[string]string{
	"user_login":      "test@test.test",
	"user_password":   "password",
	"user_auth_token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTM5NjA2MTAsImlhdCI6MTY1MzkxNzQxMCwidXNlcl9pZCI6NH0.UusEPIYIilvGkSf68BTNxioHs0LajSXT24OwnB7VE5A",
}

func Test_handleAddAndGetRequests(t *testing.T) {
	server := getServerForTestAndPrepareDatabase()

	defer server.Close()

	type want struct {
		code              int
		response          string
		locationHeader    string
		contentTypeHeader string
		contentEncoding   string
	}

	tests := []struct {
		name    string
		urlPath string
		method  string
		headers map[string]string
		body    []byte
		want    want
	}{}

	for _, testData := range tests {
		response, bodyString := testRequest(
			t, server, testData.method, testData.urlPath, testData.body, testData.headers,
		)

		if response != nil {
			defer response.Body.Close()
		}

		if testData.want.code > 0 {
			assert.Equal(t, testData.want.code, response.StatusCode)
		}

		if len(testData.want.response) > 0 {
			assert.Equal(t, testData.want.response, bodyString)
		}

		if len(testData.want.contentTypeHeader) > 0 {
			assert.Equal(t, testData.want.contentTypeHeader, response.Header.Get("Content-Type"))
		}

		if len(testData.want.contentEncoding) > 0 {
			assert.Equal(t, testData.want.contentEncoding, response.Header.Get("Content-Encoding"))
		}
	}
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

func testRequest(
	t *testing.T, ts *httptest.Server, method, path string, body []byte, headers map[string]string,
) (*http.Response, string) {
	bodyIoReader := bytes.NewBuffer(body)
	request, err := http.NewRequest(method, ts.URL+path, bodyIoReader)
	require.NoError(t, err)

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{
		CheckRedirect: func(request *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	response, err := client.Do(request)
	require.NoError(t, err)

	responseBody, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	require.NoError(t, err)

	return response, string(responseBody)
}
