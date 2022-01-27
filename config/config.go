package config

import "os"

type serverConfig struct {
	port    string
	address string
}

type config struct {
	server serverConfig
}

func NewConfig() config {
	config := config{}
	config = initConfigByEnv(config)

	return config
}

func (c config) GetServerPort() string {
	return c.server.port
}

func (c config) GetServerAddress() string {
	return c.server.address
}

func initConfigByEnv(c config) config {
	c.server.port = getEnv("SERVER_PORT", "8080")
	c.server.address = getEnv("SERVER_ADDRESS", "localhost")

	return c
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
