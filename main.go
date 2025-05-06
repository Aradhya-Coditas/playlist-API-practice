package main

import (
	"context"
	"omnenest-backend/src/constants"

	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/database/migrations"
	"omnenest-backend/src/models"

	"omnenest-backend/src/utils/configs"
	"omnenest-backend/src/utils/flags"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/postgres"

	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	initConfigs(ctx)

	// setup a logger
	initLogger(ctx)
	log := logger.GetLoggerWithoutContext()

	applicationConfig, err := configs.Get(genericConstants.ApplicationConfig)
	if err != nil {
		log.With(zap.Error(err)).Fatal(genericConstants.GetApplicationConfigError)
	}

	applicationName := applicationConfig.GetString(constants.ApplicationNameConfig)
	var dbConnectionClient *models.DBConnectionClient
	//initialize postgres client
	if applicationName == constants.MobileApp {
		postgresConfigError := postgres.InitPostgresDBConfig(ctx)
		if postgresConfigError != nil {
			log.With(zap.Error(postgresConfigError)).Fatal(genericConstants.PostgresDBInitializationError)
		}
		dbConnectionClient = postgres.GetPostGresClient()
		defer postgres.ClosePostgres(ctx)
	}

	migrationError := migrations.AutoMigrate(dbConnectionClient)
	if migrationError != nil {
		log.With(zap.Error(migrationError)).Fatal(genericConstants.CockroachDBModelsMigrationError)
	}

	// We need to call the below only when we move to any new environment
	// err := migrations.LoadExchangeConfigs(dbConnectionClient)
	// if err != nil {
	// 	log.With(zap.Error(err)).Fatal(genericConstants.CockroachDBLoadExchangeConfigsError)
	// }

}

func initConfigs(ctx context.Context) {
	// init configs
	configs.Init([]string{flags.RootConfigPath()})
}

func initLogger(ctx context.Context) {
	// Get a log config from yaml
	LoggerConfig, err := configs.Get(constants.LoggerConfig)
	if err != nil {
		panic(err)
	}
	// Setup logging from a config
	logger.SetupLogging(LoggerConfig.GetString(constants.LogLevelKey))
}
