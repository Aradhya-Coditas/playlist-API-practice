package models

import (
	"database/sql"
	"fmt"
	genericConstants "omnenest-backend/src/constants"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                           uint64    `gorm:"primaryKey" json:"id"`
	Username                     string    `gorm:"column:username" json:"username"`
	NestPublicKey                string    `gorm:"column:nest_public_key" json:"nest_public_key"`
	FirstName                    string    `gorm:"column:first_name" json:"first_name"`
	LastName                     string    `gorm:"column:last_name" json:"last_name"`
	Email                        string    `gorm:"column:email" json:"email"`
	PhoneNumber                  string    `gorm:"column:phone_number" json:"phone_number"`
	BrokerID                     int64     `gorm:"column:broker_id" json:"broker_id"`
	UserAccessToken              string    `gorm:"column:user_access_token" json:"user_access_token"`
	SequenceID                   string    `gorm:"column:sequence_id" json:"sequence_id"`
	UserRefreshToken             string    `gorm:"column:user_refresh_token" json:"user_refresh_token"`
	IsAccountBlocked             bool      `gorm:"column:is_account_blocked" json:"is_account_blocked"`
	IsGuestUser                  bool      `gorm:"column:is_guest_user" json:"is_guest_user"`
	UnblockRequested             bool      `gorm:"column:unblock_requested" json:"unblock_requested"`
	CreatedAt                    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                    time.Time `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted                    bool      `gorm:"column:is_deleted" json:"is_deleted"`
	UserType                     string    `gorm:"column:user_type" json:"user_type"`
	AccountId                    string    `gorm:"column:account_id" json:"account_id"`
	EnabledExchange              string    `gorm:"column:enabled_exchange" json:"enabled_exchange"`
	EnabledOrderType             string    `gorm:"column:enabled_order_type" json:"enabled_order_type"`
	EnabledProductType           string    `gorm:"column:enabled_product_type" json:"enabled_product_type"`
	ProductAlias                 string    `gorm:"column:product_alias" json:"product_alias"`
	TransactionFlag              string    `gorm:"column:transaction_flag" json:"transaction_flag"`
	DefaultMarketWatchlistName   string    `gorm:"column:default_market_watchlist_name" json:"default_market_watchlist_name"`
	BranchID                     string    `gorm:"column:branch_id" json:"branch_id"`
	MarketWatchCount             string    `gorm:"column:market_watch_count" json:"market_watch_count"`
	WebLink                      string    `gorm:"column:web_link" json:"web_link"`
	PasswordSpecialCharacterFlag string    `gorm:"column:password_special_character_flag" json:"password_special_character_flag"`
	UserPrivileges               string    `gorm:"column:user_privileges" json:"user_privileges"`
	YsxExchangeFlag              string    `gorm:"column:ysx_exchange_flag" json:"ysx_exchange_flag"`
	CriteriaAttribute            string    `gorm:"column:criteria_attribute" json:"criteria_attribute"`
	ClearingOrg                  string    `gorm:"column:clearing_org" json:"clearing_org"`
	AttributeOrderType           string    `gorm:"column:attribute_order_type" json:"attribute_order_type"`
	EquitySIPOperationMode       string    `gorm:"column:equity_sip_operation_mode" json:"equity_sip_operation_mode"`
	IndicesOrdering              string    `gorm:"column:indices_ordering" json:"indices_ordering"`
}

type Devices struct {
	Id              int       `json:"Id" gorm:"primary_key;auto_increment"`
	Username        string    `gorm:"column:username" json:"username"`
	DeviceId        string    `json:"deviceId" gorm:"column:device_id"`
	CreatedAt       time.Time `json:"createdAt" gorm:"column:created_at"`
	BFFPublicKey    string    `gorm:"column:bff_public_key" json:"bff_public_key"`
	BFFPrivateKey   string    `gorm:"column:bff_private_key" json:"bff_private_key"`
	DevicePublicKey string    `gorm:"column:device_public_key" json:"device_public_key"`
}

type OldBroker struct {
	Id                            int64     `gorm:"primary_key"`
	BrokerName                    string    `gorm:"column:broker_name"`
	BrokerAddress                 string    `gorm:"column:broker_address"`
	ContactInformation            string    `gorm:"column:contact_information"`
	CustomString                  string    `gorm:"column:string"`
	GuestUserDuration             string    `gorm:"column:guest_user_duration"`
	CreatedAt                     time.Time `gorm:"column:created_at"`
	UpdatedAt                     time.Time `gorm:"column:updated_at"`
	AdvancedRetentionTypeLicensed bool      `gorm:"column:adv_ret_type_licensed"`
	FAQ                           string    `gorm:"column:faq"`
	Banners                       string    `gorm:"column:banner"`
	Properties                    string    `gorm:"column:properties"`
	SebiNumber                    string    `gorm:"column:sebi_number"`
}

type StaticExchangeConfig struct {
	Id                            int64  `gorm:"primary_key"`
	Exchange                      string `gorm:"column:exchange"`
	PriceType                     string `gorm:"column:price_type"`
	RetentionType                 string `gorm:"column:retention_type"`
	AdvancedRetentionTypeLicensed bool   `gorm:"column:adv_ret_type_licensed"`
	SpreadIndicator               bool   `gorm:"column:spread_indicator"`
	AuctionIndicator              bool   `gorm:"column:auction_indicator"`
}

type CompanyMaster struct {
	CompanyCode      uint32    `gorm:"primary_key;column:company_code"`
	BSECode          int       `gorm:"column:bse_code"`
	CompanyName      string    `gorm:"column:company_name"`
	CompanyShortName string    `gorm:"column:company_short_name"`
	CategoryName     string    `gorm:"column:category_name"`
	ISIN             string    `gorm:"column:isin"`
	BSEGroup         string    `gorm:"column:bse_group"`
	MCAPType         string    `gorm:"column:mcap_type"`
	SectorCode       int       `gorm:"column:sector_code"`
	SectorName       string    `gorm:"column:sector_name"`
	IndustryCode     int       `gorm:"column:industry_code"`
	IndustryName     string    `gorm:"column:industry_name"`
	BSEListedFlag    string    `gorm:"column:bse_listed_flag"`
	DisplayType      string    `gorm:"column:display_type"`
	BSEStatus        string    `gorm:"column:bse_status"`
	LastUpdatedAt    time.Time `gorm:"column:last_updated_at"`
}

type UsersInfo struct {
	Id                           uint64    `gorm:"primary_key;column:id;autoIncrement"`
	Username                     string    `gorm:"column:username" logger:"sensitive" `
	FirstName                    string    `gorm:"column:first_name" logger:"sensitive"`
	LastName                     string    `gorm:"column:last_name" logger:"sensitive"`
	Email                        string    `gorm:"column:email" logger:"sensitive"`
	BrokerID                     uint32    `gorm:"column:broker_id" logger:"sensitive" `
	PhoneNumber                  uint64    `gorm:"column:phone_number" logger:"sensitive"`
	UserAccessToken              string    `gorm:"column:user_access_token"`
	UserRefreshToken             string    `gorm:"column:user_refresh_token"`
	IsAccountBlocked             bool      `gorm:"column:is_account_blocked"`
	IsGuestUser                  bool      `gorm:"column:is_guest_user"`
	UnblockRequested             bool      `gorm:"column:unblock_requested"`
	CreatedAt                    time.Time `gorm:"column:created_at"`
	UpdatedAt                    time.Time `gorm:"column:updated_at"`
	IsDeleted                    bool      `gorm:"column:is_deleted"`
	UserType                     string    `gorm:"column:user_type"`
	AccountID                    string    `gorm:"column:account_id" logger:"sensitive"`
	EnabledExchange              string    `gorm:"column:enabled_exchange"`
	EnabledOrderType             string    `gorm:"column:enabled_order_type"`
	EnabledProductType           string    `gorm:"column:enabled_product_type"`
	ProductAlias                 string    `gorm:"column:product_alias"`
	TransactionFlag              bool      `gorm:"column:transaction_flag"`
	DefaultMarketWatchlistName   string    `gorm:"column:default_market_watchlist_name"`
	BranchID                     string    `gorm:"column:branch_id" logger:"sensitive"`
	MarketWatchCount             uint16    `gorm:"column:market_watch_count"`
	WebLink                      string    `gorm:"column:web_link"`
	PasswordSpecialCharacterFlag bool      `gorm:"column:password_special_character_flag"`
	UserPrivileges               string    `gorm:"column:user_privileges"`
	YsxExchangeFlag              bool      `gorm:"column:ysx_exchange_flag"`
	CriteriaAttribute            string    `gorm:"column:criteria_attribute"`
	ClearingOrg                  string    `gorm:"column:clearing_org"`
	AttributeOrderType           string    `gorm:"column:attribute_order_type"`
	EquitySIPOperationMode       string    `gorm:"column:equity_sip_operation_mode"`
	IndicesOrdering              string    `gorm:"column:indices_ordering"`
	RateLimitPlanID              uint8     `gorm:"column:rate_limit_plan_id"`
	Otp                          uint32    `json:"otp" gorm:"column:otp" logger:"sensitive"`
	OTPGeneratedAt               time.Time `gorm:"column:otp_generated_at"`
	Password                     string    `gorm:"column:password" logger:"sensitive"`
	ResendCount                  uint8     `gorm:"column:resend_count"`
	PancardNumber                string    `gorm:"column:pancard_number" logger:"sensitive"`
	InvalidLoginCount            uint16    `gorm:"column:invalid_login_count"`
	LastThreePasswords           string    `gorm:"column:last_three_passwords" logger:"sensitive"`
	IsOtpVerified                bool      `gorm:"column:is_otp_verified"`
	GttEnabled                   bool      `gorm:"-" json:"gtt_enabled"`
}

type CashFlow struct {
	Id              uint64        `json:"Id" gorm:"primary_key;auto_increment"`
	CFCompanyCode   uint32        `gorm:"column:company_code;not null;uniqueIndex:cf_idx_code_metric_type"`
	MonthYearDate   string        `gorm:"column:month_year_date;not null;uniqueIndex:cf_idx_code_metric_type"`
	Metric          string        `gorm:"column:metric;not null;uniqueIndex:cf_idx_code_metric_type"`
	Value           float64       `gorm:"column:value"`
	AggregationType string        `gorm:"column:aggregation_type;not null;uniqueIndex:cf_idx_code_metric_type"`
	LastUpdatedAt   time.Time     `gorm:"column:last_updated_at;not null;default:CURRENT_TIMESTAMP"`
	CompanyMaster   CompanyMaster `gorm:"foreignKey:CFCompanyCode;references:CompanyCode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Basket struct {
	Id         uint64    `gorm:"primary_key;column:id;auto_increment"`
	UserId     uint64    `gorm:"column:user_id;not null"`
	BasketName string    `gorm:"column:basket_name"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	IsDeleted  bool      `gorm:"column:is_deleted"`
	UsersInfo  UsersInfo `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ProfitAndLoss struct {
	Id              uint64        `json:"Id" gorm:"primary_key;auto_increment"`
	PLCompanyCode   uint32        `gorm:"column:company_code;not null;uniqueIndex:pl_idx_code_metric_type"`
	MonthYearDate   string        `gorm:"column:month_year_date;not null;uniqueIndex:pl_idx_code_metric_type"`
	Metric          string        `gorm:"column:metric;not null;uniqueIndex:pl_idx_code_metric_type"`
	Value           float64       `gorm:"column:value"`
	AggregationType string        `gorm:"column:aggregation_type;not null;uniqueIndex:pl_idx_code_metric_type"`
	LastUpdatedAt   time.Time     `gorm:"column:last_updated_at;not null;default:CURRENT_TIMESTAMP"`
	CompanyMaster   CompanyMaster `gorm:"foreignKey:PLCompanyCode;references:CompanyCode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type BasketOrder struct {
	Id                uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId            uint64    `gorm:"column:user_id;not null" logger:"sensitive"`
	BasketId          uint64    `gorm:"column:basket_id;not null" json:"basket_id"`
	ExchangeName      string    `gorm:"column:exchange;not null" json:"exchange"`
	TradingSymbol     string    `gorm:"column:trading_symbol;not null" json:"trading_symbol"`
	TransactionType   string    `gorm:"column:transaction_type;not null" json:"transaction_type"`
	ProductCode       string    `gorm:"column:product_code;not null" json:"product_code"`
	PriceType         string    `gorm:"column:price_type;not null" json:"price_type"`
	RetentionType     string    `gorm:"column:retention_type;not null" json:"retention_type"`
	Price             float32   `gorm:"column:price" json:"price"`
	TriggerPrice      float32   `gorm:"column:trigger_price" json:"trigger_price"`
	Quantity          uint32    `gorm:"column:quantity;not null" json:"quantity"`
	DisclosedQuantity uint32    `gorm:"column:disclosed_quantity" json:"disclosed_quantity"`
	GTDDate           time.Time `gorm:"column:gtd_date" json:"gtd_date"`
	CreatedAt         time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	IsDeleted         bool      `gorm:"column:is_deleted" json:"is_deleted"`
	UsersInfo         UsersInfo `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Basket            Basket    `gorm:"foreignKey:BasketId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type CmotsFinancials struct {
	Id              uint64        `json:"Id" gorm:"primary_key;auto_increment"`
	CompanyCodeCF   uint32        `gorm:"column:company_code;not null;uniqueIndex:cf_idx_code_metric_type"`
	MonthYearDate   string        `gorm:"column:month_year_date;not null;uniqueIndex:cf_idx_code_metric_type"`
	Metric          string        `gorm:"column:metric;not null;uniqueIndex:cf_idx_code_metric_type"`
	Value           float64       `gorm:"column:value"`
	AggregationType string        `gorm:"column:aggregation_type;not null;uniqueIndex:cf_idx_code_metric_type"`
	LastUpdatedAt   time.Time     `gorm:"column:last_updated_at;not null;default:CURRENT_TIMESTAMP"`
	CompanyMaster   CompanyMaster `gorm:"foreignKey:CompanyCodeCF;references:CompanyCode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type BalanceSheet struct {
	Id              uint64        `json:"Id" gorm:"primary_key;auto_increment"`
	BSCompanyCode   uint32        `gorm:"column:company_code;not null;uniqueIndex:bs_idx_code_metric_type"`
	MonthYearDate   string        `gorm:"column:month_year_date;not null;uniqueIndex:bs_idx_code_metric_type"`
	Metric          string        `gorm:"column:metric;not null;uniqueIndex:bs_idx_code_metric_type"`
	Value           float64       `gorm:"column:value"`
	AggregationType string        `gorm:"column:aggregation_type;not null;uniqueIndex:bs_idx_code_metric_type"`
	LastUpdatedAt   time.Time     `gorm:"column:last_updated_at;not null;default:CURRENT_TIMESTAMP"`
	CompanyMaster   CompanyMaster `gorm:"foreignKey:BSCompanyCode;references:CompanyCode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type CmotsEtlTracker struct {
	Id               uint64    `json:"Id" gorm:"primary_key;auto_increment"`
	OperationType    string    `gorm:"column:operation_type;type:varchar(20);not null"`
	CompanyCode      uint32    `gorm:"column:company_code;not null;uniqueIndex:etl_idx_code_api_type"`
	ApiName          string    `gorm:"column:api_name;type:varchar(50);not null;uniqueIndex:etl_idx_code_api_type"`
	ReportType       string    `gorm:"column:report_type;type:varchar(2);uniqueIndex:etl_idx_code_api_type"`
	Status           string    `gorm:"column:status;type:varchar(20);not null;uniqueIndex:etl_idx_code_api_type"`
	ResponseType     string    `gorm:"column:response_type;type:varchar(30);not null;uniqueIndex:etl_idx_code_api_type"`
	RetryAttempt     uint16    `gorm:"column:retry_attempt;default:0"`
	ErrorMessage     string    `gorm:"column:error_message;type:text"`
	RecordsProcessed uint32    `gorm:"column:records_processed;default:0"`
	StartTime        time.Time `gorm:"column:start_time;not null;uniqueIndex:etl_idx_code_api_type"`
	EndTime          time.Time `gorm:"column:end_time"`
	ExecutionDate    string    `gorm:"column:execution_date;not null"`
}

type CompanyYearlyRatios struct {
	Id                      uint64        `json:"Id" gorm:"primary_key;auto_increment"`
	YRCompanyCode           uint32        `gorm:"column:company_code;not null;uniqueIndex:yr_idx_code_year_type"`
	MonthYearDate           string        `gorm:"column:month_year_date;not null;uniqueIndex:yr_idx_code_year_type"`
	MarketCapitalization    float64       `gorm:"column:market_capitalization"`
	EnterpriseValue         float64       `gorm:"column:enterprise_value"`
	PriceEarningsRatio      float32       `gorm:"column:price_earnings_ratio"`
	PriceToBookRatio        float32       `gorm:"column:price_to_book_value"`
	DividendYieldPercent    float32       `gorm:"column:dividend_yield"`
	DividendPayoutRatio     float64       `gorm:"column:dividend_payout_ratio"`
	EarningsPerShare        float64       `gorm:"column:earnings_per_share"`
	BookValuePerShare       float64       `gorm:"column:book_value_per_share"`
	ReturnOnAssets          float64       `gorm:"column:return_on_assets"`
	ReturnOnEquity          float64       `gorm:"column:return_on_equity"`
	ReturnOnCapitalEmployed float64       `gorm:"column:return_on_capital_employed"`
	EBIT                    float64       `gorm:"column:ebit"`
	EBITDA                  float64       `gorm:"column:ebitda"`
	EnterpriseValueToSales  float64       `gorm:"column:enterprise_value_to_sales"`
	EnterpriseValueToEBITDA float64       `gorm:"column:enterprise_value_to_ebitda"`
	NetIncomeMargin         float64       `gorm:"column:net_income_margin"`
	GrossIncomeMargin       float64       `gorm:"column:gross_income_margin"`
	AssetTurnover           float64       `gorm:"column:asset_turnover"`
	CurrentRatio            float32       `gorm:"column:current_ratio"`
	DebtToEquityRatio       float32       `gorm:"column:debt_to_equity_ratio"`
	NetDebtToEBITDA         float64       `gorm:"column:net_debt_to_ebitda"`
	EBITDAMargin            float64       `gorm:"column:ebitda_margin"`
	TotalShareholdersEquity float64       `gorm:"column:total_shareholders_equity"`
	ShortTermDebt           float64       `gorm:"column:short_term_debt"`
	LongTermDebt            float64       `gorm:"column:long_term_debt"`
	DilutedEarningsPerShare float64       `gorm:"column:diluted_earnings_per_share"`
	NetSales                float64       `gorm:"column:net_sales"`
	NetProfit               float64       `gorm:"column:net_profit"`
	AnnualDividend          float64       `gorm:"column:annual_dividend"`
	CostOfGoodsSold         float64       `gorm:"column:cost_of_goods_sold"`
	FaceValuePerShare       float64       `gorm:"face_value_per_share"`
	LastUpdatedAt           time.Time     `gorm:"column:last_updated_at;not null;default:CURRENT_TIMESTAMP"`
	AggregationType         string        `gorm:"column:aggregation_type;not null;uniqueIndex:yr_idx_code_year_type"`
	CompanyMaster           CompanyMaster `gorm:"foreignKey:YRCompanyCode;references:CompanyCode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ShareHoldingPattern struct {
	Id                             uint64        `json:"Id" gorm:"primary_key;auto_increment"`
	SHCompanyCode                  uint32        `gorm:"column:company_code;not null;uniqueIndex:sh_idx_code_year_type"`
	Promoters                      float64       `gorm:"column:promoters;not null"`
	ForeignInstitutionalInvestors  float64       `gorm:"column:foreign_institutional_investors;not null"`
	Others                         float64       `gorm:"column:others;not null"`
	PublicInvestors                float64       `gorm:"column:public_investors;not null"`
	DomesticInstitutionalInvestors float64       `gorm:"column:domestic_institutional_investors;not null"`
	MonthYearDate                  string        `gorm:"column:month_year_date;not null;uniqueIndex:sh_idx_code_year_type"`
	LastUpdatedAt                  time.Time     `gorm:"column:last_updated_at;not null;default:CURRENT_TIMESTAMP"`
	CompanyMaster                  CompanyMaster `gorm:"foreignKey:SHCompanyCode;references:CompanyCode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ScripMaster struct {
	Id                       string          `gorm:"column:id;primaryKey"`
	ScripToken               string          `gorm:"not null;uniqueIndex:idx_scrip_token_exchange_segment;column:scrip_token"`
	Group                    string          `gorm:"column:group"`
	ExchangeSegment          string          `gorm:"not null;uniqueIndex:idx_scrip_token_exchange_segment;column:exchange_segment"`
	InstrumentType           string          `gorm:"column:instrument_type"`
	SymbolName               string          `gorm:"column:symbol_name"`
	TradingSymbol            string          `gorm:"column:trading_symbol"`
	OptionType               string          `gorm:"column:option_type"`
	UniqueKey                string          `gorm:"column:unique_key"`
	ISIN                     string          `gorm:"column:isin"`
	AssetCode                string          `gorm:"column:asset_code"`
	LotSize                  uint32          `gorm:"column:lot_size"`
	TickSize                 sql.NullFloat64 `gorm:"column:tick_size"`
	ExpiryDate               sql.NullString  `gorm:"column:expiry_date"`
	Multiplier               uint32          `gorm:"column:multiplier"`
	DecimalPrecision         uint16          `gorm:"column:decimal_precision"`
	StrikePrice              sql.NullFloat64 `gorm:"column:strike_price"`
	DisplayStrikePrice       sql.NullFloat64 `gorm:"column:display_strike_price"`
	Description              string          `gorm:"column:description"`
	SubGroup                 string          `gorm:"column:sub_group"`
	AmcCode                  string          `gorm:"column:amc_code"`
	ContractID               string          `gorm:"column:contract_id"`
	CombinedScripToken       string          `gorm:"column:combined_scrip_token"`
	ScripReserved            string          `gorm:"column:scrip_reserved"`
	ExchangeSymbolName       string          `gorm:"column:exchange_symbol_name"`
	RmsMarketProtection      float32         `gorm:"column:rms_market_protection"`
	ContractType             string          `gorm:"column:contract_type"`
	ExchangeName             string          `gorm:"column:exchange"`
	LastUpdatedAt            time.Time       `gorm:"column:last_updated_at;autoUpdateTime"`
	CompanyCode              uint32          `gorm:"column:company_code"`
	DisplayExpiryDate        sql.NullString  `gorm:"column:display_expiry_date"`
	SymbolLeg2               string          `gorm:"column:symbol_leg2"`
	CreatedAt                time.Time       `gorm:"column:created_at"`
	LastTradedPrice          float32         `gorm:"column:last_traded_price"`
	AverageTradePrice        float32         `gorm:"column:average_trade_price"`
	VolumeTradedToday        uint64          `gorm:"column:volume_traded_today"`
	OpenPrice                float32         `gorm:"column:open_price"`
	HighPrice                float32         `gorm:"column:high_price"`
	LowPrice                 float32         `gorm:"column:low_price"`
	PreviousClosePrice       float32         `gorm:"column:previous_close_price"`
	HighLimitPriceProtection float32         `gorm:"column:high_limit_price_protection"`
	LowLimitPriceProtection  float32         `gorm:"column:low_limit_price_protection"`
	YearlyHighPrice          float32         `gorm:"column:yearly_high_price"`
	YearlyLowPrice           float32         `gorm:"column:yearly_low_price"`
	TotalTradedValue         float32         `gorm:"column:total_traded_value"`
	MarketType               string          `gorm:"column:market_type"`
	LowerCircuitPrice        float32         `gorm:"column:lower_circuit_price"`
	UpperCircuitPrice        float32         `gorm:"column:upper_circuit_price"`
	TotalAskQuantity         uint32          `gorm:"column:total_ask_quantity"`
	TotalBidQuantity         uint32          `gorm:"column:total_bid_quantity"`
	OpenInterest             float32         `gorm:"column:open_interest"`
	LastTradeQuantity        uint32          `gorm:"column:last_trade_quantity"`
	LastTradeTime            time.Time       `gorm:"column:last_trade_time"`
}

type Brokers struct {
	Id         uint32 `gorm:"primaryKey;column:id;autoIncrement"`
	BrokerName string `gorm:"column:broker_name;size:255;not null"`
}

type BrokerWatchlists struct {
	Id            uint64    `gorm:"primaryKey;column:id;autoIncrement"`
	BrokerId      uint32    `gorm:"column:broker_id;not null"`
	WatchlistName string    `gorm:"column:watchlist_name;size:255;not null"`
	LastUpdatedAt time.Time `gorm:"column:last_updated_at;not null"`
	ScripCount    uint16    `gorm:"column:scrip_count"`
	Broker        Brokers   `gorm:"foreignKey:BrokerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Watchlists struct {
	Id            uint64    `gorm:"primaryKey;column:id;autoIncrement"`
	UserId        uint64    `gorm:"column:user_id;not null"`
	WatchlistName string    `gorm:"column:watchlist_name;size:255;not null"`
	LastUpdatedAt time.Time `gorm:"column:last_updated_at;not null"`
	ScripCount    uint16    `gorm:"column:scrip_count"`
	UsersInfo     UsersInfo `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type WatchlistScrips struct {
	Id          uint64      `gorm:"primaryKey;column:id;autoIncrement"`
	WatchlistId uint64      `gorm:"column:watchlist_id;not null"`
	ScripId     string      `gorm:"column:scrip_id;not null"`
	Order       uint16      `gorm:"column:order"`
	Watchlists  Watchlists  `gorm:"foreignKey:WatchlistId;references:Id;constraint:OnDelete:CASCADE"`
	Scrip       ScripMaster `gorm:"foreignKey:ScripId;references:Id;constraint:OnDelete:CASCADE"`
}

type CompanyProfile struct {
	Id                   uint64        `gorm:"column:id;primaryKey;autoIncrement"`
	CPCompanyCode        uint32        `gorm:"column:company_code;unique;not null"`
	CompanyName          string        `gorm:"column:company_name;not null"`
	ISIN                 string        `gorm:"column:isin;not null"`
	Chairman             string        `gorm:"column:chairman"`
	ManagingDirector     string        `gorm:"column:managing_director"`
	IncorporationDate    string        `gorm:"column:incorporation_date"`
	RegisteredAddress1   string        `gorm:"column:registered_address_1"`
	RegisteredAddress2   string        `gorm:"column:registered_address_2"`
	RegisteredDistrict   string        `gorm:"column:registered_district"`
	RegisteredState      string        `gorm:"column:registered_state"`
	RegisteredPIN        string        `gorm:"column:registered_pin"`
	Telephone1           string        `gorm:"column:telephone_1"`
	Telephone2           string        `gorm:"column:telephone_2"`
	Fax1                 string        `gorm:"column:fax_1"`
	Fax2                 string        `gorm:"column:fax_2"`
	Auditors             string        `gorm:"column:auditors"`
	FaceValuePerShare    float32       `gorm:"column:face_value_per_share"`
	MarketLot            uint32        `gorm:"column:market_lot"`
	CompanySecretary     string        `gorm:"column:company_secretary"`
	Email                string        `gorm:"column:email"`
	Website              string        `gorm:"column:website"`
	DirectorName         string        `gorm:"column:director_name"`
	DirectorDesignation  string        `gorm:"column:director_designation"`
	HeadOfficeAddress1   string        `gorm:"column:head_office_address_1"`
	HeadOfficeAddress2   string        `gorm:"column:head_office_address_2"`
	HeadOfficeCity       string        `gorm:"column:head_office_city"`
	HeadOfficeStateCode  string        `gorm:"column:head_office_state_code"`
	HeadOfficePIN        uint16        `gorm:"column:head_office_pin"`
	HeadOfficeTelephone1 string        `gorm:"column:head_office_telephone_1"`
	HeadOfficeEmail      string        `gorm:"column:head_office_email"`
	HeadOfficeCountry    string        `gorm:"column:head_office_country"`
	RegistrarCode        string        `gorm:"column:registrar_code"`
	RegistrarName        string        `gorm:"column:registrar_name"`
	RegistrarAddress1    string        `gorm:"column:registrar_address_1"`
	RegistrarAddress2    string        `gorm:"column:registrar_address_2"`
	RegistrarAddress3    string        `gorm:"column:registrar_address_3"`
	RegistrarAddress4    string        `gorm:"column:registrar_address_4"`
	RegistrarTelephone   string        `gorm:"column:registrar_telephone"`
	RegistrarFax         string        `gorm:"column:registrar_fax"`
	RegistrarEmail       string        `gorm:"column:registrar_email"`
	RegistrarWebsite     string        `gorm:"column:registrar_website"`
	IndustryName         string        `gorm:"column:industry_name"`
	LastUpdatedAt        time.Time     `gorm:"column:last_updated_at;not null;default:CURRENT_TIMESTAMP"`
	CompanyMaster        CompanyMaster `gorm:"foreignKey:CPCompanyCode;references:CompanyCode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
type Song struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Title  string `gorm:"column:title;not null" json:"title"`
	Artist string `gorm:"column:artist" json:"artist"`
}

type Playlist struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"column:name;unique;not null" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	UserID      int    `gorm:"column:user_id;not null" json:"user_id"`
	Songs       []Song `gorm:"many2many:playlist_songs" json:"songs"`
}

type PlaylistSongs struct {
    PlaylistID int `gorm:"column:playlist_id;not null"`
    SongID     int `gorm:"column:song_id;not null"`
}

type DBConnectionClient struct {
	GormDb *gorm.DB
	SqlDb  *sql.DB
}

func (User) TableName() string {
	return genericConstants.UserTable
}

func (Devices) TableName() string {
	return genericConstants.DeviceTable
}

func (Brokers) TableName() string {
	return genericConstants.BrokerTable
}

func (BrokerWatchlists) TableName() string {
	return genericConstants.BrokerWatchlistTable
}

func (StaticExchangeConfig) TableName() string {
	return genericConstants.StaticExchangeConfigTables
}

func (CompanyMaster) TableName() string {
	return genericConstants.CompanyMasterTable
}

func (UsersInfo) TableName() string {
	return genericConstants.UsersInfoTable
}

func (CashFlow) TableName() string {
	return genericConstants.CashFlowTable
}

func (Basket) TableName() string {
	return genericConstants.BasketTable
}

func (Watchlists) TableName() string {
	return genericConstants.WatchlistsTable
}

func (BasketOrder) TableName() string {
	return genericConstants.BasketOrderTable
}

func (ProfitAndLoss) TableName() string {
	return genericConstants.ProfitAndLossTable
}

func (CompanyYearlyRatios) TableName() string {
	return genericConstants.CompanyYearlyRatiosTable
}

func (CmotsEtlTracker) TableName() string {
	return genericConstants.CmotsEtlTrackerTable
}

func (CmotsFinancials) TableName() string {
	return genericConstants.CmotsFinancialsTable
}

func (BalanceSheet) TableName() string {
	return genericConstants.BalanceSheetTable
}

func (ShareHoldingPattern) TableName() string {
	return genericConstants.ShareHoldingPatternTable
}

func (ScripMaster) TableName() string {
	return genericConstants.ScripMasterTable
}

func (CompanyProfile) TableName() string {
	return genericConstants.CompanyProfileTable
}

func (WatchlistScrips) TableName() string {
	return genericConstants.WatchlistScripsTable
}

func InitDBConstraints(db *gorm.DB) error {
	//Query for Watchlist API
	err := db.Exec(genericConstants.UniqueIndexForWatchlistQuery).Error
	if err != nil {
		return err
	}

	err = db.Exec(genericConstants.UniqueIndexForWatchlistScripsQuery).Error
	if err != nil {
		return err
	}

	err = db.Exec(genericConstants.UniqueIndexForBasketQuery).Error
	if err != nil {
		return err
	}

	return nil
}

func InitDBIndexes(db *gorm.DB) error {
	indexQueries := []string{
		genericConstants.IndexForBasketOrderOnBasketId,
		genericConstants.IndexOnBasketsIdActive,
		genericConstants.IndexForScripmasterOnExchangeAndContractType,
		genericConstants.IndexForScripmaster,
		genericConstants.IndexForWatchlistsOnUserId,
		genericConstants.IndexForWatchlistsOnUserIdId,
		genericConstants.IndexForWatchlistsOnScripsWatchlistId,
		genericConstants.IndexForWatchlistScripsOnWatchlistIdScripOrder,
		genericConstants.IndexForScripMasterOnExchangeNameToken,
		genericConstants.IndexForScripMasterOnExchangeSegmentToken,
		genericConstants.IndexForWatchlistsOnUserIdLastUpdatedAt,
		genericConstants.IndexForScripMasterOnExchangeSegementTradingSymbol,
	}

	for _, query := range indexQueries {
		if err := db.Exec(query).Error; err != nil {
			return fmt.Errorf(genericConstants.InitDBIndexesError, query, err)
		}
	}

	return nil
}
