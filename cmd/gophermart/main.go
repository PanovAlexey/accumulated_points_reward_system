package main

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/config"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/application/service"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/logging"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/repository"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/servers"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	config := config.NewConfig()
	logger := logging.GetLogger(config)

	defer logger.Close()

	userRegistrationRepository := repository.NewUserRepository()
	userRegistrationSerice := service.NewUserRegistrationService(userRegistrationRepository)

	handler := http.NewHandler(logger, userRegistrationSerice)
	server := new(servers.Server)

	if err := server.Run(config, handler.InitRoutes()); err != nil {
		logger.Fatalf("error occurred while running http server: %s", err.Error())
	}
}

func init() {
}
