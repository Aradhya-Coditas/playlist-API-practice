package models

type ApplicationConfig struct {
	SwaggerConfig  SwaggerConfig
	Server         Server
	FilePath       FilePath
	Token          Token
	AppConfig      AppConfig
	WebSocket      WebSocket
	Authentication Authentication
	Orders         Orders
	DBMetrics      DBMetrics
	Watchlist      Watchlist
	DataLoad       DataLoad
	BackOfficeURLs BackOfficeURLs
}

type BackOfficeURLs struct {
	ClientLedgerURL    string
	GlobalPLURL        string
	HoldingsURL        string
	ClientFASummaryURL string
}
type DBMetrics struct {
	CockroachDB bool
	Postgres    bool
}

type AppConfig struct {
	UseMocks                    bool
	UseDBMocks                  bool
	EnableUIBFFEncDec           bool
	EnableNestEncryption        bool
	UseFrontendErrorFormat      bool
	EnableRateLimit             bool
	EnableMRVData               bool
	RateLimitIntervalInSecond   int
	RateLimitRequestPerInterval int
	BannerURLPrefix             string
	EnableOpenTelemetry         bool
	DefaultMWScripsCount        int
	EnableSSL                   bool
	CMOTSToken                  string
	EnableIPOCache              bool
	IsMonolith                  bool
	ApplicationName             string
}

type Token struct {
	AccessTokenSecretKey     string
	RefreshTokenSecretKey    string
	AccessTokenExpiryInDays  int
	RefreshTokenExpiryInDays int
	SecretKey                string
	EnableTokenCompression   bool
}

type FilePath struct {
	EncryptionKeysPath           string
	MiddlewareEncryptionKeysPath string
}

type Server struct {
	ServerPort int
}

type SwaggerConfig struct {
	SwaggerHost string
}

type WebSocket struct {
	WebSocketPort               int
	WebSocketHost               string
	WebSocketProtocol           string
	WebSocketInsecureSkipVerify bool
	WebSocketPingInterval       int
}

type Authentication struct {
	PasswordValidationRegex  string
	OtpLength                uint16
	AllowedMaximumLoginCount uint16
	OtpExpiryInMinutes       uint16
	ResendOTPTimeout         uint16
	AdminUserName            string
	AdminPassword            string
}

type Orders struct {
	BasketLimit               uint16
	BasketNameValidationRegex string
	MaximumBasketOrdersLimit  uint16
}

type Watchlist struct {
	WatchlistLimit               uint16
	ScripLimitPerWatchlist       uint16
	WatchlistNameValidationRegex string
}

type DataLoad struct {
	ScripMasterTimeZone            string
	ScripMasterTimeIntervalInHours uint16
	ScripMasterBatchSize           uint16
	CmotsYearMonthRegex            string
	CmotsSchedulerIntervalInDays   uint16
	CmotsSchedulerRunTime          string
	CmotsMaxRetryCount             uint16
}
