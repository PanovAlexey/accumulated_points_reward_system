package config

import (
	"flag"
	"fmt"
	"os"
)

type serverConfig struct {
	address string
}

type storageConfig struct {
	databaseDsn string
}

type applicationConfig struct {
	environment string
	loggerDsn   string
	isDebug     bool
}

type externalSystemsConfig struct {
	accrualSystemAddress string
}

type Config struct {
	server          serverConfig
	application     applicationConfig
	storage         storageConfig
	externalSystems externalSystemsConfig
}

func NewConfig() Config {
	config := Config{}
	config = initConfigByEnv(config)
	config = initConfigByFlag(config)

	return config
}

func (c Config) GetServerAddress() string {
	return c.server.address
}

func (c Config) GetAppEnvironment() string {
	return c.application.environment
}

func (c Config) GetAppLoggerDsn() string {
	return c.application.loggerDsn
}

func (c Config) IsAppDebugMode() bool {
	return c.application.isDebug
}

func (c Config) GetDatabaseDsn() string {
	return c.storage.databaseDsn
}

func (c Config) GetAccrualSystemAddress() string {
	return c.externalSystems.accrualSystemAddress
}

func initConfigByEnv(c Config) Config {
	c.server.address = getEnv("RUN_ADDRESS", "0.0.0.0:8080")
	c.application.isDebug = getEnv("IS_DEBUG", "false") == "true"
	c.application.environment = getEnv("ENVIRONMENT", "dev")
	c.application.loggerDsn = getEnv("LOGGER_DSN", "https://1e8c898aac7c45259639d9a6eae5a926@o1210124.ingest.sentry.io/6345772")
	c.storage.databaseDsn = getEnv("DATABASE_URI", "")
	c.externalSystems.accrualSystemAddress = getEnv("ACCRUAL_SYSTEM_ADDRESS", "http://localhost:8080/api/orders/")

	return c
}

func initConfigByFlag(config Config) Config {
	if flag.Parsed() {
		fmt.Println("Error occurred. Re-initializing the config")
		return config
	}

	serverAddress := flag.String("a", "", "RUN_ADDRESS")
	databaseURI := flag.String("d", "", "DATABASE_URI")
	accrualSystemAddress := flag.String("r", "", "ACCRUAL_SYSTEM_ADDRESS")

	flag.Parse()

	if len(*serverAddress) > 0 {
		config.server.address = *serverAddress
	}

	if len(*databaseURI) > 0 {
		config.storage.databaseDsn = *databaseURI
	}

	if len(*accrualSystemAddress) > 0 {
		config.externalSystems.accrualSystemAddress = *accrualSystemAddress
	}

	return config
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
