package main

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/handlers/http"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/servers"
	"log"
)

func main() {
	handler := http.NewHandler()
	server := new(servers.Server)

	if err := server.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
