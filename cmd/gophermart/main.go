package main

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/config"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/logging"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/servers"
	"github.com/joho/godotenv"
)

func main() {
	logger := logging.GetLogger()
	defer logger.Close()

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	config := config.NewConfig()
	handler := http.NewHandler()
	server := new(servers.Server)

	if err := server.Run(config, handler.InitRoutes()); err != nil {
		logger.Fatalf("error occurred while running http server: %s", err.Error())
	}
}

func init() {
}
