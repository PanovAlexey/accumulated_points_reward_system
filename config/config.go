package config

import (
	"flag"
	"fmt"
	"os"
)

type serverConfig struct {
	address string
}

type config struct {
	server serverConfig
}

func NewConfig() config {
	config := config{}
	config = initConfigByEnv(config)
	config = initConfigByFlag(config)

	return config
}

func (c config) GetServerAddress() string {
	return c.server.address
}

func initConfigByEnv(c config) config {
	c.server.address = getEnv("RUN_ADDRESS", "0.0.0.0:8080")

	return c
}

func initConfigByFlag(config config) config {
	if flag.Parsed() {
		fmt.Println("Error occurred. Re-initializing the config")
		return config
	}

	serverAddress := flag.String("a", "", "RUN_ADDRESS")

	flag.Parse()

	if len(*serverAddress) > 0 {
		config.server.address = *serverAddress
	}

	return config
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
