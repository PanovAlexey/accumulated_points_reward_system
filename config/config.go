package config

import (
	"flag"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Server          serverConfig
	Application     applicationConfig
	Storage         storageConfig
	ExternalSystems externalSystemsConfig
}

type serverConfig struct {
	Address      string `envconfig:"run_address" default:"0.0.0.0:8080"`
	DebugAddress string `envconfig:"debug_address" default:"0.0.0.0:8081"`
}

type storageConfig struct {
	DatabaseDsn string `envconfig:"database_uri" default:""`
}

type applicationConfig struct {
	Environment string `envconfig:"environment" default:"dev"`
	LoggerDsn   string `envconfig:"logger_dsn" default:"https://1e8c898aac7c45259639d9a6eae5a926@o1210124.ingest.sentry.io/6345772"`
	IsDebug     bool   `envconfig:"is_debug" default:"false"`
}

type externalSystemsConfig struct {
	AccrualSystemAddress string `envconfig:"accrual_system_address" default:"http://localhost:8080/api/orders/"`
}

func NewConfig() Config {
	config := Config{}

	config = initConfigByEnv(config)
	config = initConfigByFlag(config)

	return config
}

func initConfigByEnv(c Config) Config {
	var serverConfig serverConfig
	err := envconfig.Process("", &serverConfig)

	if err != nil {
		log.Fatal(err.Error())
	}

	var applicationConfig applicationConfig
	err = envconfig.Process("", &applicationConfig)

	if err != nil {
		log.Fatal(err.Error())
	}

	var storageConfig storageConfig
	err = envconfig.Process("", &storageConfig)

	if err != nil {
		log.Fatal(err.Error())
	}

	var externalSystemsConfig externalSystemsConfig
	err = envconfig.Process("", &externalSystemsConfig)

	if err != nil {
		log.Fatal(err.Error())
	}

	c.Server = serverConfig
	c.Storage = storageConfig
	c.ExternalSystems = externalSystemsConfig
	c.Application = applicationConfig

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
		config.Server.Address = *serverAddress
	}

	if len(*databaseURI) > 0 {
		config.Storage.DatabaseDsn = *databaseURI
	}

	if len(*accrualSystemAddress) > 0 {
		config.ExternalSystems.AccrualSystemAddress = *accrualSystemAddress
	}

	return config
}
