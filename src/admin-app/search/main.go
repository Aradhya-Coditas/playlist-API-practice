package main

import (
	serviceConstants "admin-app/search/commons/constants"
	"admin-app/search/router"
	"context"
	"fmt"
	genericConstants "omnenest-backend/src/constants"
	loggerMiddleware "omnenest-backend/src/middleware/logger"

	"omnenest-backend/src/utils/configs"
	"omnenest-backend/src/utils/flags"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/postgres"
	"omnenest-backend/src/utils/tracer"
	"os"

	sdkTrace "go.opentelemetry.io/otel/sdk/trace"

	"go.uber.org/zap"
)

// @title omnenest-backend
// @version 1.0
// @description Omnenest backend for WebSocket Search micro-service (Middleware layer).
// @BasePath /v1
// @query.collection.format multi
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @x-extension-openapi {"example": "value on a json format"}
func main() {
	ctx := context.Background()
	initConfigs(ctx)

	// setup a logger
	initLogger(ctx)
	log := logger.GetLoggerWithoutContext()

	// set the host name
	configs.SetHostName(serviceConstants.ServiceName)

	err := postgres.InitPostgresDBConfig(ctx)
	if err != nil {
		log.With(zap.Error(err)).Fatal(genericConstants.PostgresDBInitializationError)
	}
	defer postgres.ClosePostgres(ctx)

	err = configs.InitApplicationConfigs(ctx)
	if err != nil {
		log.With(zap.Error(err)).Fatal(genericConstants.ConfigBindingFailedError)
		panic(fmt.Errorf(genericConstants.ConfigBindingFailedError))
	}

	// initialize traces
	applicationConfig := configs.GetApplicationConfig()
	isEnableTracing := applicationConfig.AppConfig.EnableOpenTelemetry
	if isEnableTracing {
		traceProvider := initTracer(ctx, log)
		defer func() {
			if err := traceProvider.Shutdown(context.Background()); err != nil {
				log.With(zap.Error(err)).Fatal(genericConstants.ErrorShoutDownTraceProvider)
			}
		}()
	}

	// initialize nestIniConfig
	err = configs.InitNestIniConfigs(ctx)
	if err != nil {
		log.With(zap.Error(err)).Fatal(genericConstants.ConfigBindingFailedError)
		panic(fmt.Errorf(genericConstants.ConfigBindingFailedError))
	}

	// start router
	startRouter(ctx)
}

func initConfigs(ctx context.Context) {
	configs.Init([]string{flags.BaseConfigPath(), flags.MockConfigPath()})
}

func initLogger(ctx context.Context) {
	LoggerConfig, err := configs.Get(genericConstants.LoggerConfig)
	if err != nil {
		panic(err)
	}
	logger.SetupLogging(LoggerConfig.GetString(genericConstants.LogLevelKey))
}

func initTracer(ctx context.Context, log logger.Logger) *sdkTrace.TracerProvider {
	// Get a Tracer config from yaml
	endpointUrl := os.Getenv(genericConstants.OLTPEndPointEnvConfig)
	if endpointUrl == "" {
		endpointUrl = genericConstants.OLTPHttpEndpointUrl
	}
	traceProvider, err := tracer.InitTracer(ctx, true, serviceConstants.ServiceName, endpointUrl)
	if err != nil {
		log.With(zap.Error(err)).Fatal(genericConstants.ErrorInitializeTraceProvider)
	}
	return traceProvider
}

func startRouter(ctx context.Context) {

	log := logger.GetLoggerWithoutContext()

	// get router
	router := router.GetRouter(loggerMiddleware.Logger())
	// now start router
	log.Info(fmt.Sprintf(genericConstants.RunningServerPort, serviceConstants.PortDefaultValue))

	// Load SLL Configuration
	// SSLCertificate := os.Getenv(genericConstants.SSLCertificateCRTConfig)
	// SSLKey := os.Getenv(genericConstants.SSLCertificateKeyConfig)

	// if SSLCertificate == "" || SSLKey == "" {
	// 	log.Fatal(genericConstants.SSLEnvironmentVariablesMissingError)
	// }
	// Run the Server
	// err := router.RunTLS(fmt.Sprintf(":%d", constants.PortDefaultValue), SSLCertificate, SSLKey)
	err := router.Run(fmt.Sprintf(":%d", serviceConstants.PortDefaultValue))
	if err != nil {
		log.With(zap.Error(err)).Fatal(genericConstants.ExternalServiceError)
	}
}
