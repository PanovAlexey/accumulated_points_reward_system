package main

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/config"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/service"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/databases"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/logging"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/repository"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/servers"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	config := config.NewConfig()
	logger := logging.GetLogger(config)

	defer logger.Close()

	db := getDatabaseConnector(logger, config)
	defer db.Close()

	migrationService := databases.GetMigrationService(db)
	migrationService.MigrateUp()

	userRegistrationRepository := repository.NewUserRepository(db)
	userRegistrationService := service.NewUserRegistrationService(userRegistrationRepository)

	orderStatusGetter := service.GetOrderStatusGetter()
	orderNumberValidator := service.GetLuhnAlgorithmChecker()
	orderRepository := repository.NewOrderRepository(db)
	orderLoaderService := service.NewOrderLoaderService(orderRepository, orderNumberValidator, orderStatusGetter)

	handler := http.NewHandler(logger, userRegistrationService, orderLoaderService)
	server := new(servers.Server)

	if err := server.Run(config, handler.InitRoutes()); err != nil {
		logger.Fatalf("error occurred while running http server: %s", err.Error())
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s", err.Error())
	}
}

func getDatabaseConnector(logger logging.LoggerInterface, config config.Config) *sqlx.DB {
	db, err := databases.NewPostgresDB(config.GetDatabaseDsn())
	if err != nil {
		logger.Fatalf("error occurred while initializing database connection: %s", err.Error())
	}

	return db
}
