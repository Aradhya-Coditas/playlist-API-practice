package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"omnenest-backend/src/constants"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/models"
	"omnenest-backend/src/utils"
	"omnenest-backend/src/utils/configs"
	"omnenest-backend/src/utils/logger"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	gormLogger "gorm.io/gorm/logger"

	"github.com/DATA-DOG/go-sqlmock"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
)

var postgresClient *models.DBConnectionClient

// InitPostgresDBConfig initializes the Postgres database configuration and establishes a connection to the database.
func InitPostgresDBConfig(ctx context.Context) error {

	log := logger.GetLoggerWithoutContext()
	PostgresConfig, err := configs.Get(genericConstants.PostgresConfig)
	if err != nil {
		return fmt.Errorf(genericConstants.GetPostgresConfigError, err)
	}

	var postgresConfig models.PostgresConfig

	postgresConfig.Host = utils.GetEnv(genericConstants.PostgresHostEnv, PostgresConfig.GetString(genericConstants.PostgresHostKey))
	postgresConfig.Port = utils.GetEnv(genericConstants.PostgresPortEnv, PostgresConfig.GetString(genericConstants.PostgresPortKey))
	postgresConfig.User = utils.GetEnv(genericConstants.PostgresUserEnv, PostgresConfig.GetString(genericConstants.PostgresUserKey))
	postgresConfig.Password = utils.GetEnv(genericConstants.PostgresPasswordEnv, PostgresConfig.GetString(genericConstants.PostgresPasswordKey))
	postgresConfig.DBName = utils.GetEnv(genericConstants.PostgresDBNameEnv, PostgresConfig.GetString(genericConstants.PostgresDBNameKey))
	postgresConfig.SSLMode = PostgresConfig.GetString(genericConstants.PostgresSSLModeKey)
	postgresConfig.TimeZone = PostgresConfig.GetString(genericConstants.PostgresTimeZoneKey)
	postgresConfig.IsMockConnection = PostgresConfig.GetBool(genericConstants.PostgresIsMockConnectionKey)

	if postgresConfig.IsMockConnection {
		// Connection Mock for postgres
		err = ConnectMockDatabase(ctx, postgresConfig, log)
		if err != nil {
			return fmt.Errorf(genericConstants.PostgresMockConnectionError, err)
		}
	} else {
		// Connection for postgres
		err = ConnectPostgresDatabase(ctx, postgresConfig, log)
		if err != nil {
			return fmt.Errorf(genericConstants.PostgresConnectionError, err)
		}
	}
	return nil
}

// ConnectPostgresDatabase establishes a connection to the Postgres database using the provided configuration.
func ConnectPostgresDatabase(ctx context.Context, postgresConfig models.PostgresConfig, log logger.Logger) error {
	// Construct the DSN string
	dsn := fmt.Sprintf(genericConstants.DNSString, postgresConfig.Host, postgresConfig.Port, postgresConfig.User, postgresConfig.Password, postgresConfig.DBName, postgresConfig.SSLMode, postgresConfig.TimeZone)
	applicationConfig, err := configs.Get(genericConstants.ApplicationConfig)
	if err != nil {
		panic(fmt.Errorf(genericConstants.ConfigBindingFailedError))
	}
	enablePostgresMetrics := applicationConfig.GetBool(genericConstants.PostgresConfig)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		return err
	}

	if enablePostgresMetrics {
		// Adds the Prometheus plugin to the Postgres database connection
		db.Use(prometheus.New(prometheus.Config{
			DBName:          postgresConfig.DBName,
			RefreshInterval: 15,
			StartServer:     false,
			MetricsCollector: []prometheus.MetricsCollector{
				&prometheus.Postgres{VariableNames: []string{"Threads_running"}},
			},
		}))
	}
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = sqlDB.PingContext(ctx)
	if err != nil {
		return err
	}
	log.Info(genericConstants.PostgresConnectionSuccessful)

	SetPostgresClient(db, sqlDB)

	return nil
}

// ConnectMockDatabase establishes a connection to a mock Postgres database using the provided configuration.
func ConnectMockDatabase(ctx context.Context, postgresConfig models.PostgresConfig, log logger.Logger) error {
	mockDb, _, err := sqlmock.New()
	if err != nil {
		return err
	}
	dialConnection := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: constants.PostgresDriverName,
	})
	db, err := gorm.Open(dialConnection, &gorm.Config{})
	if err != nil {
		return err
	}
	SetPostgresClient(db, mockDb)
	log.Info(genericConstants.PostgresMockConnectionSuccessful)
	return nil
}

func SetPostgresClient(db *gorm.DB, sqlDB *sql.DB) {
	postgresClient = &models.DBConnectionClient{GormDb: db, SqlDb: sqlDB}
}

func ClosePostgres(ctx context.Context) {
	log := logger.GetLoggerWithoutContext()
	if postgresClient != nil {
		err := postgresClient.SqlDb.Close()
		if err != nil {
			log.Error(genericConstants.ClosePostgresClientError, zap.Error(err))
		}
	}
}

func GetPostGresClient() *models.DBConnectionClient {
	return postgresClient
}
