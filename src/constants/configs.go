package constants

// Configuration Constants
const (
	RouterV1Config                          = "v1"
	RouterV2Config                          = "v2"
	RouterV3Config                          = "v3"
	ApplicationJSONTypeConfig               = "application/json"
	JsonConfig                              = "json"
	LoggerConfig                            = "logger"
	LoggerSensitiveTag                      = "sensitive"
	EncryptDecryptTagName                   = "db"
	EncryptDecryptTagValue                  = "crypt"
	ApplicationConfig                       = "application"
	MultiCastConfig                         = "multicast"
	RedisConfig                             = "redis"
	NestIniConfig                           = "nestIniConfig"
	ApiConfig                               = "apis"
	ModelConfig                             = "model"
	PostgresConfig                          = "postgres"
	OpenSearchConfig                        = "opensearch"
	SwaggerInfoHttpSchemeConfig             = "http"
	SwaggerInfoHttpsSchemeConfig            = "https"
	NestAPIMockConfig                       = "nestAPIResponse"
	CmotsAPIMockConfig                      = "cmotsAPIResponse"
	NestSocketMockConfig                    = "nestSocketResponse"
	NestAPICallConfig                       = "NestAPICall"
	EndPointConfig                          = "endPoint"
	ServerPortConfig                        = "serverPort.serverPort"
	ResponseConfig                          = "response"
	WarningTextConfig                       = "warningText"
	ErrorMessageConfig                      = "errorMessage"
	RawResponseConfig                       = "rawResponse"
	RequestPayloadConfig                    = "requestPayload"
	RequestURLConfig                        = "requestURL"
	AllowHeaderOriginConfig                 = "Origin"
	ExposeHeaderContentLengthConfig         = "Content-Length"
	SSLCertificateCRTConfig                 = "SSL_CERTIFICATE_CRT"
	SSLCertificateKeyConfig                 = "SSL_CERTIFICATE_KEY"
	UseMocksConfig                          = "appConfig.UseMocks"
	UseDBMocksConfig                        = "appConfig.UseDBMocks"
	IsMonolith                              = "appConfig.IsMonolith"
	EnableUIBFFEncDecConfig                 = "appConfig.EnableUIBFFEncDec"
	UseFrontendErrorFormatConfig            = "appConfig.UseFrontendErrorFormat"
	SwaggerHostKey                          = "swagger.swaggerHost"
	FilePathKey                             = "filePath.encryptionKeysPath"
	MiddlewareFilePathKey                   = "filePath.middlewareEncryptionKeysPath"
	AccessTokenSecretKey                    = "token.accessSecretKey"
	RefreshTokenSecretKey                   = "token.refreshSecretKey"
	AccessTokenExpiryInDaysKey              = "token.accessTokenExpiryInDays"
	RefreshTokenExpiryInDaysKey             = "token.refreshTokenExpiryInDays"
	EnableTokenCompressionKey               = "token.enableTokenCompression"
	EnableRateLimitConfig                   = "appConfig.EnableRateLimit"
	EnableMRVDataConfig                     = "appConfig.EnableMRVData"
	RateLimitIntervalInSecondConfig         = "appConfig.RateLimitIntervalInSecond"
	RateLimitRequestPerIntervalConfig       = "appConfig.RateLimitRequestPerInterval"
	BannerURLPrefixConfig                   = "appConfig.BannerURLPrefix"
	EnableOpenTelemetryConfig               = "appConfig.EnableOpenTelemetry"
	DefaultMWScripsCountConfig              = "appConfig.DefaultMWScripsCount"
	RateLimitIntervalInSecondEnvKey         = "RATE_LIMIT_INTERVAL_IN_SECOND"
	RateLimitRequestPerIntervalEnvKey       = "RATE_LIMIT_REQUEST_PER_INTERVAL"
	UIConfig                                = "UIConfig"
	OLTPEndPointEnvConfig                   = "OLTP_HTTP_ENDPOINT_URL"
	APPEnvironment                          = "APP_ENVIRONMENT"
	UseMocksEnvKey                          = "USE_MOCKS"
	UseDBMocksEnvKey                        = "USE_DB_MOCKS"
	IsMonolithEnvKey                        = "IS_MONOLITH"
	EnableUIBFFEncDecEnvKey                 = "ENABLE_UIBFF_ENCRYPT_DECRYPT"
	EnableOpenTelemetryEnvKey               = "ENABLE_OPEN_TELEMETRY"
	EnableRateLimitEnvKey                   = "ENABLE_RATE_LIMIT"
	DefaultMWScripsCountEnvKey              = "DEFAULT_MW_SCRIPTS_COUNT"
	EnableSSLConfig                         = "appConfig.EnableSSL"
	EnableSSLEnvKey                         = "ENABLE_SSL"
	CMOTSAPICallConfig                      = "CMOTSAPICall"
	CMOTSTokenConfig                        = "appConfig.CMOTSToken"
	CMOTSTokenEnvKey                        = "CMOTS_TOKEN"
	JKeyConfig                              = "jKey"
	CockroachDBConfig                       = "cockroachDB"
	NestCPPCallBackResponse                 = "nestCPPCallBackResponse"
	EnableIPOCacheEnvKey                    = "ENABLE_IPO_CACHE"
	EnableIPOCacheConfig                    = "appConfig.EnableIPOCache"
	NestCallBackWrapperConfig               = "NestCallBackWrapper"
	NestCPPCallBackFunctionConfig           = "nestCPPCallBackFunction"
	NestCPPCallBackFunctionParametersConfig = "nestCPPCallBackFunctionParameters"
	DBQueryConfig                           = "DBQuery"
	DBConditionsConfig                      = "dbConditions"
	DBResponseConfig                        = "dbResponse"
	DBRecordsFoundConfig                    = "dbRecordsFound"
	ConditionsConfig                        = "conditions"
	ApplicationNameConfig                   = "appConfig.ApplicationName"
	ApplicationNameEnvKey                   = "APPLICATION_NAME"
)

// Custom validaton constants
const (
	ValidateEnumConfig     = "ValidateEnum"
	RetentionDateConfig    = "RetentionDateValidation"
	ScannerTypeValueConfig = "ScannerTypeValueValidation"
	DataOfBirthConfig      = "DateOfBirthValidaton"
	PANConfig              = "PANValidation"
	BidLengthValidation    = "BidLengthValidation"
)

// Configuration Keys
const (
	GetInitialKey           = "omnenest.GetInitialKey"
	GetPreAuthenticationKey = "omnenest.GetPreAuthenticationKey"
)

// NEST Configuration Constants
const (
	EnableNestEncryptionConfig = "appConfig.EnableNestEncryption"
	KeySize                    = 512
	JKey                       = "jKey"
	JData                      = "jData"
)

const (
	NestAPIRestBaseURL         = "restbaseurl"
	NestAPIScannerBaseURL      = "scannerbaseurl"
	NestAPIGlobalSearchBaseURL = "globalsearchbaseurl"
	NestAPIIPOBaseURL          = "ipobaseurl"
	NestAPIMutualFundBaseURL   = "mutualfundbaseurl"
	NestAPITypeToURLMapping    = "nestapitypetourlmapping"
	CMOTSAPITypeToURLmapping   = "cmotsapitypetourlmapping"
	CMOTSAPIBaseURL            = "cmotsbaseurl"
)

// Postgres Configuration Constants
const (
	PostgresHostKey             = "host"
	PostgresPortKey             = "port"
	PostgresUserKey             = "user"
	PostgresPasswordKey         = "password"
	PostgresDBNameKey           = "dbName"
	PostgresSSLModeKey          = "sslMode"
	PostgresTimeZoneKey         = "TimeZone"
	PostgresIsMockConnectionKey = "isMockConnection"
	PostgresDriverName          = "postgres"
	PostgresHostEnv             = "POSTGRES_HOST"
	PostgresPortEnv             = "POSTGRES_PORT"
	PostgresUserEnv             = "POSTGRES_USER"
	PostgresPasswordEnv         = "POSTGRES_PASSWORD"
	PostgresDBNameEnv           = "POSTGRES_DB_NAME"
	PostgresSecretKey           = "secretKey"
)

// OpenSearch Configuration Constants
const (
	OpenSearchHostKey             = "host"
	OpenSearchPortKey             = "port"
	OpenSearchUserKey             = "user"
	OpenSearchPasswordKey         = "password"
	OpenSearchSSLModeKey          = "sslMode"
	OpenSearchTimeZoneKey         = "TimeZone"
	OpenSearchIsMockConnectionKey = "isMockConnection"
	OpenSearchCACertPathKey       = "caCertPath"
	OpenSearchHostEnv             = "OPENSEARCH_HOST"
	OpenSearchPortEnv             = "OPENSEARCH_PORT"
	OpenSearchUserEnv             = "OPENSEARCH_USER"
	OpenSearchPasswordEnv         = "OPENSEARCH_PASSWORD"
	OpenSearchCACertPathEnv       = "OPENSEARCH_CACERT_PATH"
	OpenSearchSSLModeEnv          = "OPENSEARCH_SSL_MODE"
)

// Redis Configuration Constants
const (
	RedisConfigUsername = "default"
	RedisConfigPassword = "hb9hlBCZ3EUykxX1VGfoP6yq2Fj9SKtL"
)

// Crypto Configuration Constants
const (
	KeyTypePublic          = "PUBLIC KEY"
	KeyTypePrivate         = "PRIVATE KEY"
	ObtainedHashedPassword = "Received hashed password"
)

// Mock Configuration Constants
const (
	MockResponseKey         = "mockresponse"
	MockEncryptionNeededKey = "encryptionneeded"
	MockErrorTextKey        = "errortext"
)

// REST Methods Constants
const (
	PostMethod    = "POST"
	GetMethod     = "GET"
	PatchMethod   = "PATCH"
	DeleteMethod  = "DELETE"
	PutMethod     = "PUT"
	OptionsMethod = "OPTIONS"
)

// Encryption Key File Names
const (
	PreAuthServerKey     = "preAuthServerPublicKey.pem"
	ClientPublicKey      = "clientPublicKey.pem"
	ClientPrivateKey     = "clientPrivateKey.pem"
	PreAuthServerKeyHash = "preAuthServerKeyHash"
	BffPrivateKey        = "bffPrivateKey.pem"
	BffPublicKey         = "bffPublicKey.pem"
)

// Redis Configuration Configuration
const (
	RedisURLKey    = "redis.HostUrl"
	PoolSize       = "redis.PoolSize"
	MinIdleConns   = "redis.MinIdleConns"
	MaxConnAge     = "redis.MaxConnAge"
	PoolTimeout    = "redis.PoolTimeout"
	ReadTimeout    = "redis.ReadTimeout"
	UserName       = "redis.UserName"
	PassWord       = "redis.PassWord"
	IsRedisCluster = "redis.IsRedisCluster"
)

// ScyllaDB Constants
const (
	ScyllaDBYamlFile        = "resources/scylladb.yml"
	ScyllaDBSelectionPolicy = "us-east-1"
	ScyllaDBKeySpace        = "omnenest_scylladb_test"
)

// Custom Validations Constants
const (
	CustomValidatorTag                   = "CustomValidation"
	DateOfBirthFieldName                 = "DateOfBirth"
	PANFieldName                         = "PAN"
	DateOfBirthFormat                    = `^\d{4}-\d{2}-\d{2}$`
	PANFormat                            = `^[A-Z]{5}[0-9]{4}[A-Z]$`
	DateOfBirthFormatMatch               = "2006-01-02"
	RetentionDateFormatMatch             = "2/1/2006"
	RetentionDateFieldName               = "RetentionTimestamp"
	RetentionDateFormat                  = `^(\d{1,2})/(\d{1,2})/\d{4}$` //`^(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-\d{4}$`
	DateFormatMatch                      = `(\d{2}/\d{2}/\d{4} \d{2}:\d{2}:\d{2})|(\d{2}/\d{2}/\d{4})`
	SpreadInstrumentTypeScripTokenFormat = `^\d+ \d+$`
	NonDigitSequence                     = `[^0-9]+`
	Digit                                = `\d+`
	WatchListLimitError                  = `^nest internal server error: number of scrips already is :  (\d+) the number of scrips chosen is :(\d+) the maximum scrips can be added is : (\d+) #end#$`
	WatchListMaxScripsMatch              = `maximum scrips can be added is : \d+`
	DecimalZeroOrComma                   = `\.0+$|,`
	BrokerRecommendationPattern          = `\s+([\d,]+(?:\.\d+)?)RS`
	MatchHttpUrl                         = `https?://[^\s]+|www\.[^\s]+`
	ReplaceSpecialCharsWithSpace         = `[^a-zA-Z0-9]`
)

// Research Call Response Field Extraction Tags
const (
	InitialPriceTag = "INITIATION"
	StopLossTag     = "SL"
	TargetTag       = "TARGET"
	PriceAtCallTag  = "PRICE AT CALL"
)

// regex patterns key
const (
	NonDigitSequenceKey               = "nonDigitSequence"
	DigitKey                          = "digit"
	InstrumentTypeScripTokenFormatKey = "instrumentTypeScripTokenFormat"
	DecimalZeroOrCommaKey             = "decimalZeroOrComma"
	InitialPriceTagKey                = "initialPriceTag"
	StopLossTagKey                    = "stopLossTag"
	TargetTagKey                      = "targetTag"
	PriceAtCallTagKey                 = "priceAtCallTag"
	DateFormatMatchKey                = `dateFormatMatch`
	WatchListLimitErrorKey            = `watchListLimitError`
	WatchListMaxScripsMatchKey        = `watchListMaxScripsMatch`
	MatchHttpUrlKey                   = "matchHttpUrl"
	ReplaceSpecialCharsWithSpaceKey   = "replaceSpecialCharsWithSpace"
	PanCardValidationKey              = "panCardValidation"
)

// Header JWT Context Configuration
const (
	UserID               = "userId"
	PanNumber            = "panNumber"
	ServerKeyPair        = "serverKeyPair"
	UserSessionToken     = "userSessionToken"
	TokenPayloadClaims   = "tokenPayload"
	BFFCtxPublicKey      = "bffPublicKey"
	BFFCtxPrivateKey     = "bffPrivateKey"
	DeviceCtxPublicKey   = "devicePublicKey"
	TokenExpiration      = "expiration"
	AccessToken          = "accessToken"
	RefreshToken         = "refreshToken"
	TokenType            = "tokenType"
	TokenPayload         = "tokenPayload"
	AccountID            = "accountId"
	IssueID              = "issueId"
	SourceKey            = "source"
	CriteriaAttributeKey = "criteriaAttribute"
	ProductAlias         = "productAlias"
	ClearingOrg          = "clearingOrg"
	BranchName           = "branchName"
	EquitySIPMode        = "1"
	EnabledExchangesKey  = "enabledExchanges"
	GttEnabledKey        = "gttEnabled"
)

// DNS String Configuration
const (
	DNSString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s"
)

// OpenSearch DSN String Configuration
const (
	OpenSearchDNSString = "https://%s:%s"
)

// Path for test configs file
const (
	TestConfig             = "../../../../setupTest/testConfigs"
	MockResponseConfigPath = "../../../../utils/mockResources"
)

// String Config
const (
	BlankSpace   = " "
	EmptySpace   = ""
	AndWord      = "and"
	AndWithSpace = " and "
	Comma        = ","
	Underscore   = "_"
	Dot          = "."
	LikeWildcard = "%"
	CommonPipeSeperator = "|"
)

// Numeric Config
const (
	ZeroNumericValue = 0
)

// Banners Env variable keys
const (
	BannerURLPrefixEnvKey = "BANNER_URL_PREFIX"
)

// NestAPIUrl Env variable keys
const (
	RestBaseUrlKey         = "REST_BASE_URL"
	ScannerBaseUrlKey      = "SCANNER_BASE_URL"
	GlobalSearchBaseUrlKey = "GLOBAL_SEARCH_BASE_URL"
	IPOBaseUrlKey          = "IPO_BASE_URL"
	CMOTSBaseURLKey        = "CMOTS_BASE_URL"
)

// WebSocket config variable keys
const (
	WebsocketPortKey     = "websocket.websocketPort"
	WebsocketHostKey     = "websocket.websocketHost"
	WebsocketProtocolKey = "websocket.websocketProtocol"
	WebsocketInsecure    = "websocket.websocketInsecureSkipVerify"
	WebsocketPing        = "websocket.websocketPingInterval"
	WebsocketPong        = "websocket.websocketPongTimeout"
	WebSocketIdleTimeout = "websocket.websocketIdleTimeout"
	WebsocketMaxMsgSize  = "websocket.websocketMaxMessageSize"
	WebsocketMaxConn     = "websocket.websocketMaxConnections"
)

// Backend server name
const (
	NESTBackendServerName  = "nestBackend"
	CMOTSBackendServerName = "cmotsBackend"
)

// Authentication config
const (
	OtpLength                      = "authentication.OtpLength"
	AllowedMaximumLoginCount       = "authentication.AllowedMaximumLoginCount"
	PasswordValidationRegex        = "authentication.PasswordValidationRegex"
	OtpExpiryInMinutes             = "authentication.OtpExpiryInMinutes"
	ResendOTPTimeout               = "authentication.ResendOTPTimeout"
	AdminUserName                  = "authentication.AdminUserName"
	AdminPassword                  = "authentication.AdminPassword"
	NestQAPICallbackTimeOut        = "authentication.NestQAPITimeOut"
	OtpLengthEnvKey                = "OTP_LENGTH"
	AllowedMaximumLoginCountEnvKey = "ALLOWED_MAXIMUM_LOGIN_COUNT"
	PasswordValidationRegexEnvKey  = "PASSWORD_VALIDATION_REGEX"
	OtpExpiryInMinutesEnvKey       = "OTP_EXPIRY_IN_MINUTES"
	ResendOTPTimeoutEnvKey         = "RESEND_OTP_TIMEOUT"
	AdminUserNameEnvKey            = "ADMIN_USER_NAME"
	AdminPasswordEnvKey            = "ADMIN_PASSWORD"
)

// path for ComInit file
const (
	ComInitPath = "../../configs/COMInitFile.ini"
)

// Orders Config
const (
	BasketLimit                     = "orders.BasketLimit"
	BasketLimitEnvKey               = "BASKET_LIMIT"
	BasketNameValidationRegex       = "orders.BasketNameValidationRegex"
	BasketNameValidationRegexEnvKey = "BASKET_NAME_VALIDATION_REGEX"
	MaximumBasketOrdersLimit        = "orders.MaximumBasketOrdersLimit"
	MaximumBasketOrdersLimitEnvKey  = "MAXIMUM_BASKET_ORDERS_LIMIT"
)

// Watchlist Config
const (
	WatchlistLimit                     = "watchlist.WatchlistLimit"
	WatchlistLimitEnvKey               = "WATCHLIST_LIMIT"
	ScripLimitPerWatchlist             = "watchlist.ScripLimitPerWatchlist"
	ScripLimitPerWatchlistEnvKey       = "SCRIP_TOKEN_LIMIT_PER_WATCHLIST"
	WatchlistNameValidationRegex       = "watchlist.WatchlistNameValidationRegex"
	WatchlistNameValidationRegexEnvKey = "WATCHLIST_NAME_VALIDATION_REGEX"
)

// NestInitConfig Env variable keys and yaml keys
const (
	MmlLocBrokAddrEnvKey    = "MML_LOC_BROK_ADDR"
	MmlDmnSrvrAddrEnvKey    = "MML_DMN_SRVR_ADDR"
	MmlDsFoAddrEnvKey       = "MML_DS_FO_ADDR"
	MmlLicSrvrAddrEnvKey    = "MML_LIC_SRVR_ADDR"
	AdminNameEnvKey         = "ADMIN_NAME"
	IntDDNameEnvKey         = "INT_DD_NAME"
	BcastDDNameEnvKey       = "BCAST_DD_NAME"
	IntReqDDNameEnvKey      = "INT_REQ_DD_NAME"
	RmsGetPrsntDDNameEnvKey = "RMS_GET_PRSNT_DD_NAME"
	TouchlineDDNameEnvKey   = "TOUCHLINE_DD_NAME"
	RmsDDNameEnvKey         = "RMS_DD_NAME"

	MmlLocBrokAddrKey    = "NEST_ENV_SETTINGS.MML_LOC_BROK_ADDR"
	MmlDmnSrvrAddrKey    = "NEST_ENV_SETTINGS.MML_DMN_SRVR_ADDR"
	MmlDsFoAddrKey       = "NEST_ENV_SETTINGS.MML_DS_FO_ADDR"
	MmlLicSrvrAddrKey    = "NEST_ENV_SETTINGS.MML_LIC_SRVR_ADDR"
	AdminNameKey         = "ADMIN_NAME.ADMIN_NAME"
	IntDDNameKey         = "INT_DD_NAME.DD_NAME"
	BcastDDNameKey       = "BCAST_DD_NAME.DD_NAME"
	IntReqDDNameKey      = "INT_REQ_DD_NAME.DD_NAME"
	RmsGetPrsntDDNameKey = "RMS_GET_PRSNT_DD_NAME.DD_NAME"
	TouchlineDDNameKey   = "TOUCHLINE_DD_NAME.DD_NAME"
	RmsDDNameKey         = "RMS_DD_NAME.DD_NAME"

	NestEnvSettingsKey = "NEST_ENV_SETTINGS"
	DDIniKey           = "DD_NAME"
)

// CockroachDB Configuration Constants
const (
	CockroachDBHostKey             = "host"
	CockroachDBPortKey             = "port"
	CockroachDBUserKey             = "user"
	CockroachDBPasswordKey         = "password"
	CockroachDBNameKey             = "dbName"
	CockroachDBSSLModeKey          = "sslMode"
	CockroachDBMaxIdleConnsKey     = "maxIdleConns"
	CockroachDBMaxOpenConnsKey     = "maxOpenConns"
	CockroachDBConnMaxLifetimeKey  = "connMaxLifetime"
	CockroachDBTimeZoneKey         = "TimeZone"
	CockroachDBIsMockConnectionKey = "isMockConnection"
	CockroachDBDriverName          = "cockroachdb"
	CockroachDBHostEnv             = "COCKROACHDB_HOST"
	CockroachDBPortEnv             = "COCKROACHDB_PORT"
	CockroachDBUserEnv             = "COCKROACHDB_USER"
	CockroachDBPasswordEnv         = "COCKROACHDB_PASSWORD"
	CockroachDBNameEnv             = "COCKROACHDB_NAME"
	CockroachDBSecretKey           = "secretKey"
)

// DB Metrics Configuration Constants
const (
	CockroachDBMetrics = "dbMetrics.cockroachDB"
	PostgresMetrics    = "dbMetrics.postgres"
)

// DataLoad Config
const (
	ScripMasterDataLoadTimeZone           = "data-load.ScripMasterTimeZone"
	ScripMasterDataLoadTimeZoneEnvKey     = "TIME_ZONE"
	ScripMasterDataLoadTimeIntervalInMin  = "data-load.ScripMasterTimeIntervalInHours"
	ScripMasterDataLoadTimeIntervalEnvKey = "SCRIP_MASTER_TIME_INTERVAL_IN_HOURS"
	ScripMasterDataLoadBatchSize          = "data-load.ScripMasterBatchSize"
	ScripMasterDataLoadBatchSizeEnvKey    = "SCRIP_MASTER_BATCH_SIZE"
	CmotsYearMonthRegex                   = "data-load.CmotsYearMonthRegex"
	CmotsYearMonthRegexEnvKey             = "CMOTS_YEAR_MONTH_REGEX"
	CmotsSchedulerIntervalInDays          = "data-load.CmotsSchedulerIntervalInDays"
	CmotsSchedulerIntervalInDaysEnvKey    = "CMOTS_SCHEDULER_INTERVAL_IN_DAYS"
	CmotsSchedulerRunTime                 = "data-load.CmotsSchedulerRunTime"
	CmotsSchedulerRunTimeEnvKey           = "CMOTS_SCHEDULER_RUN_TIME"
	CmotsMaxRetryCount                    = "data-load.CmotsMaxRetryCount"
	CmotsMaxRetryCountEnvKey              = "CMOTS_MAX_RETRY_COUNT"
)

// BackOfficeURLs Config
const (
	ClientLedgerURLEnvKey = "CLIENT_LEDGER_URL"
	GlobalPLURLEnvKey     = "GLOBAL_PL_URL"
	HoldingsURLEnvKey     = "HOLDINGS_URL"
	ClientFASummaryEnvKey = "CLIENT_FA_SUMMARY_URL"
	ClientLeger           = "back-office-urls.ClientLedger"
	GlobalPL              = "back-office-urls.GlobalPL"
	Holdings              = "back-office-urls.Holdings"
	ClientFASummary       = "back-office-urls.ClientFASummary"
)

// multicast config
const (
	MulticastPortKey    = "multicastPort"
	MulticastIPKey      = "multicastIp"
	MulticastUdpVersion = "UDPVersion"
	MaxDatagramSizeKey  = "maxDatagramSize"
	MulticastIPEnvKey   = "MULTICAST_IP"
	MulticastPortEnvKey = "MULTICAST_PORT"
)

// Application Name
const (
	MobileApp = "app"
	AdminApp  = "admin-app"
)
