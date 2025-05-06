package configs

import (
	"context"
	"omnenest-backend/src/constants"
	"omnenest-backend/src/models"
	"os"
	"strconv"
)

var applicationConfig *models.ApplicationConfig

func InitApplicationConfigs(ctx context.Context) error {
	// Get a Apis config from yaml
	applicationViperConfig, err := Get(constants.ApplicationConfig)
	if err != nil {
		return err
	}

	applicationConfig = &models.ApplicationConfig{
		SwaggerConfig: models.SwaggerConfig{
			SwaggerHost: applicationViperConfig.GetString(constants.SwaggerHostKey),
		},
		Server: models.Server{
			ServerPort: applicationViperConfig.GetInt(constants.ServerPortConfig)},
		AppConfig: models.AppConfig{
			UseMocks:                    getEnvBool(constants.UseMocksEnvKey, applicationViperConfig.GetBool(constants.UseMocksConfig)),
			UseDBMocks:                  getEnvBool(constants.UseDBMocksEnvKey, applicationViperConfig.GetBool(constants.UseDBMocksConfig)),
			IsMonolith:                  getEnvBool(constants.IsMonolithEnvKey, applicationViperConfig.GetBool(constants.IsMonolith)),
			EnableUIBFFEncDec:           getEnvBool(constants.EnableUIBFFEncDecEnvKey, applicationViperConfig.GetBool(constants.EnableUIBFFEncDecConfig)),
			EnableNestEncryption:        applicationViperConfig.GetBool(constants.EnableNestEncryptionConfig),
			UseFrontendErrorFormat:      applicationViperConfig.GetBool(constants.UseFrontendErrorFormatConfig),
			EnableRateLimit:             getEnvBool(constants.EnableRateLimitEnvKey, applicationViperConfig.GetBool(constants.EnableRateLimitConfig)),
			EnableMRVData:               applicationViperConfig.GetBool(constants.EnableMRVDataConfig),
			RateLimitIntervalInSecond:   getEnvInt(constants.RateLimitIntervalInSecondEnvKey, applicationViperConfig.GetInt(constants.RateLimitIntervalInSecondConfig)),
			RateLimitRequestPerInterval: getEnvInt(constants.RateLimitRequestPerIntervalEnvKey, applicationViperConfig.GetInt(constants.RateLimitRequestPerIntervalConfig)),
			BannerURLPrefix:             getEnv(constants.BannerURLPrefixEnvKey, applicationViperConfig.GetString(constants.BannerURLPrefixConfig)),
			EnableOpenTelemetry:         getEnvBool(constants.EnableOpenTelemetryEnvKey, applicationViperConfig.GetBool(constants.EnableOpenTelemetryConfig)),
			DefaultMWScripsCount:        getEnvInt(constants.DefaultMWScripsCountEnvKey, applicationViperConfig.GetInt(constants.DefaultMWScripsCountConfig)),
			EnableSSL:                   getEnvBool(constants.EnableSSLEnvKey, applicationViperConfig.GetBool(constants.EnableSSLConfig)),
			CMOTSToken:                  getEnv(constants.CMOTSTokenEnvKey, applicationViperConfig.GetString(constants.CMOTSTokenConfig)),
			EnableIPOCache:              getEnvBool(constants.EnableIPOCacheEnvKey, applicationViperConfig.GetBool(constants.EnableIPOCacheConfig)),
			ApplicationName:             getEnv(constants.ApplicationNameEnvKey, applicationViperConfig.GetString(constants.ApplicationNameConfig)),
		},
		Token: models.Token{
			AccessTokenSecretKey:     applicationViperConfig.GetString(constants.AccessTokenSecretKey),
			RefreshTokenSecretKey:    applicationViperConfig.GetString(constants.RefreshTokenSecretKey),
			AccessTokenExpiryInDays:  applicationViperConfig.GetInt(constants.AccessTokenExpiryInDaysKey),
			RefreshTokenExpiryInDays: applicationViperConfig.GetInt(constants.RefreshTokenExpiryInDaysKey),
			EnableTokenCompression:   applicationViperConfig.GetBool(constants.EnableTokenCompressionKey),
		},
		FilePath: models.FilePath{
			EncryptionKeysPath:           applicationViperConfig.GetString(constants.FilePathKey),
			MiddlewareEncryptionKeysPath: applicationViperConfig.GetString(constants.MiddlewareFilePathKey),
		},
		WebSocket: models.WebSocket{
			WebSocketPort:               applicationViperConfig.GetInt(constants.WebsocketPortKey),
			WebSocketHost:               applicationViperConfig.GetString(constants.WebsocketHostKey),
			WebSocketProtocol:           applicationViperConfig.GetString(constants.WebsocketProtocolKey),
			WebSocketInsecureSkipVerify: applicationViperConfig.GetBool(constants.WebsocketInsecure),
			WebSocketPingInterval:       applicationViperConfig.GetInt(constants.WebsocketPing),
		},
		Authentication: models.Authentication{
			PasswordValidationRegex:  getEnv(constants.PasswordValidationRegexEnvKey, applicationViperConfig.GetString(constants.PasswordValidationRegex)),
			OtpLength:                getEnvUint16(constants.OtpLengthEnvKey, applicationViperConfig.GetUint16(constants.OtpLength)),
			AllowedMaximumLoginCount: getEnvUint16(constants.AllowedMaximumLoginCountEnvKey, applicationViperConfig.GetUint16(constants.AllowedMaximumLoginCount)),
			OtpExpiryInMinutes:       getEnvUint16(constants.OtpExpiryInMinutesEnvKey, applicationViperConfig.GetUint16(constants.OtpExpiryInMinutes)),
			ResendOTPTimeout:         getEnvUint16(constants.ResendOTPTimeoutEnvKey, applicationViperConfig.GetUint16(constants.ResendOTPTimeout)),
			AdminUserName:            getEnv(constants.AdminUserNameEnvKey, applicationViperConfig.GetString(constants.AdminUserName)),
			AdminPassword:            getEnv(constants.AdminPasswordEnvKey, applicationViperConfig.GetString(constants.AdminPassword)),
		},
		Orders: models.Orders{
			BasketLimit:               getEnvUint16(constants.BasketLimitEnvKey, applicationViperConfig.GetUint16(constants.BasketLimit)),
			BasketNameValidationRegex: getEnv(constants.BasketNameValidationRegexEnvKey, applicationViperConfig.GetString(constants.BasketNameValidationRegex)),
			MaximumBasketOrdersLimit:  getEnvUint16(constants.MaximumBasketOrdersLimitEnvKey, applicationViperConfig.GetUint16(constants.MaximumBasketOrdersLimit)),
		},
		DBMetrics: models.DBMetrics{
			CockroachDB: applicationViperConfig.GetBool(constants.CockroachDBMetrics),
			Postgres:    applicationViperConfig.GetBool(constants.PostgresMetrics),
		},
		Watchlist: models.Watchlist{
			WatchlistLimit:               getEnvUint16(constants.WatchlistLimitEnvKey, applicationViperConfig.GetUint16(constants.WatchlistLimit)),
			ScripLimitPerWatchlist:       getEnvUint16(constants.ScripLimitPerWatchlistEnvKey, applicationViperConfig.GetUint16(constants.ScripLimitPerWatchlist)),
			WatchlistNameValidationRegex: getEnv(constants.WatchlistNameValidationRegexEnvKey, applicationViperConfig.GetString(constants.WatchlistNameValidationRegex)),
		},
		DataLoad: models.DataLoad{
			ScripMasterTimeZone:            getEnv(constants.ScripMasterDataLoadTimeZoneEnvKey, applicationViperConfig.GetString(constants.ScripMasterDataLoadTimeZone)),
			ScripMasterTimeIntervalInHours: getEnvUint16(constants.ScripMasterDataLoadTimeIntervalEnvKey, applicationViperConfig.GetUint16(constants.ScripMasterDataLoadTimeIntervalInMin)),
			ScripMasterBatchSize:           getEnvUint16(constants.ScripMasterDataLoadBatchSizeEnvKey, applicationViperConfig.GetUint16(constants.ScripMasterDataLoadBatchSize)),
			CmotsYearMonthRegex:            getEnv(constants.CmotsYearMonthRegexEnvKey, applicationViperConfig.GetString(constants.CmotsYearMonthRegex)),
			CmotsSchedulerIntervalInDays:   getEnvUint16(constants.CmotsSchedulerIntervalInDaysEnvKey, applicationViperConfig.GetUint16(constants.CmotsSchedulerIntervalInDays)),
			CmotsSchedulerRunTime:          getEnv(constants.CmotsSchedulerRunTimeEnvKey, applicationViperConfig.GetString(constants.CmotsSchedulerRunTime)),
			CmotsMaxRetryCount:             getEnvUint16(constants.CmotsMaxRetryCountEnvKey, applicationViperConfig.GetUint16(constants.CmotsMaxRetryCount)),
		},
		BackOfficeURLs: models.BackOfficeURLs{
			ClientLedgerURL:    getEnv(constants.ClientLedgerURLEnvKey, applicationViperConfig.GetString(constants.ClientLeger)),
			GlobalPLURL:        getEnv(constants.GlobalPLURLEnvKey, applicationViperConfig.GetString(constants.GlobalPL)),
			HoldingsURL:        getEnv(constants.HoldingsURLEnvKey, applicationViperConfig.GetString(constants.Holdings)),
			ClientFASummaryURL: getEnv(constants.ClientFASummaryEnvKey, applicationViperConfig.GetString(constants.ClientFASummary)),
		},
	}
	return nil
}

func GetApplicationConfig() *models.ApplicationConfig {
	return applicationConfig
}

func getEnvInt(name string, defaultValue int) int {
	valueStr := os.Getenv(name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvUint16(name string, defaultValue uint16) uint16 {
	valueStr := os.Getenv(name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return uint16(value)
	}
	return defaultValue
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvBool(name string, defaultValue bool) bool {
	valueStr := os.Getenv(name)
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}
