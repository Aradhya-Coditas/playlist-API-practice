package constants

const (
	ConnectionSuccessful                = "Connected successfully"
	RunningServerPort                   = "Running Server on port : %v"
	PostgresConnectionSuccessful        = "Successfully connected to Postgres"
	PostgresMockConnectionSuccessful    = "Successfully connected to Mock Postgres"
	RedisConnectionSuccessful           = "Successfully connected to Redis"
	NestAPICallSuccessful               = "NestAPICall successful"
	CMOTSAPICallSuccessful              = "CMOTSAPICall successful"
	NestAPICallWithEncryptedData        = "NestAPICall request with encrypted data"
	NestInternalServerErrorStartTag     = "nest internal server error: " // This is the start tag for the nest internal server error message
	NestInternalServerErrorEndTag       = " #end#"                       // This is the end tag for the nest internal server error message
	GenericErrorKey                     = "generic"                      // This is the key for generic error to be passed as the field name in response
	BFFResponseSuccessMessage           = "success"                      // This is the success message for the BFF response
	ContentType                         = "Content-Type"
	ValidationErrorKey                  = "validation"
	AvailableRateLimitKey               = "Available-Limit"
	Password                            = "password"
	OTP                                 = "otp"
	CockroachDBConnectionSuccessful     = "Successfully connected to CockroachDB"
	CockroachDBMockConnectionSuccessful = "Successfully connected to Mock CockroachDB"
)

// Nest Request Constants
const (
	NestMRVTrueValue  = "Y"
	NestMRVFalseValue = "N"
)

// Response Message Constants
const (
	ResponseMessageKey    = "message"
	ApiOkStatusMessage    = "Ok"
	ApiNotOkStatusMessage = "Not_Ok"
	Status500             = "500"
	Authorization         = "Authorization"
	Bearer                = "Bearer "
	Success               = true
)

// Request Message Constants
const (
	NestInputZeroValue      = "0"
	NestInputZeroValueFloat = 0.0
	NestInputNAValue        = "NA"
	OrderTypeMarketValue    = "MKT"
	OrderTypeSLM            = "SL-M"
	OrderTypeSPValue        = "SP"
	OrderTypeLimitValue     = "L"
	OrderTypeSLValue        = "SL"
)

// Validation Constants
const (
	GreaterThanValue = "gt"
	LessThanValue    = "lt"
)

// Column Name Constants
const (
	Id                             = "id"
	Username                       = "username"
	EmailId                        = "email"
	PhoneNumber                    = "phone_number"
	SequenceId                     = "sequence_id"
	IsAccountBlocked               = "is_account_blocked"
	UnblockRequested               = "unblock_requested"
	UpdatedAt                      = "updated_at"
	UserAccessToken                = "user_access_token"
	UserRefreshToken               = "user_refresh_token"
	FirstName                      = "first_name"
	LastName                       = "last_name"
	DeviceId                       = "device_id"
	BFFPublicKey                   = "bff_public_key"
	BFFPrivateKey                  = "bff_private_key"
	DevicePublicKey                = "device_public_key"
	AccountId                      = "account_id"
	BrokerId                       = "broker_id"
	BrokerName                     = "broker_name"
	UserType                       = "user_type"
	EnabledExchange                = "enabled_exchange"
	EnabledOrderType               = "enabled_order_type"
	EnabledProductType             = "enabled_product_type"
	ProductAliasKey                = "product_alias"
	TransactionFlag                = "transaction_flag"
	DefaultMarketWatchlistName     = "default_market_watchlist_name"
	BranchId                       = "branch_id"
	MarketWatchCount               = "market_watch_count"
	WebLink                        = "web_link"
	PasswordSpecialCharacterFlag   = "password_special_character_flag"
	UserPrivileges                 = "user_privileges"
	YSXExchangeFlag                = "ysx_exchange_flag"
	CriteriaAttribute              = "criteria_attribute"
	InterOpClearingOrg             = "clearing_org"
	AttributeOrderType             = "attribute_order_type"
	EquitySIPOperationMode         = "equity_sip_operation_mode"
	PriceTypeColumn                = "price_type"
	ExchangeColumn                 = "exchange"
	RetentionTypeColumn            = "retention_type"
	TitleColumn                    = "title"
	ContentSnippetColumn           = "content_snippet"
	LinkColumn                     = "link"
	PubDateColumn                  = "pub_date"
	SourceColumn                   = "source"
	IndicesOrderingColumn          = "indices_ordering"
	FAQColumn                      = "faq"
	Banners                        = "banner"
	Properties                     = "properties"
	SebiNumber                     = "sebi_number"
	SectorName                     = "sector_name"
	Classification                 = "mcap_type"
	IndustryName                   = "industry_name"
	ISIN                           = "isin"
	CompanyCode                    = "company_code"
	MarketCapitalization           = "market_capitalization"
	PriceEarningsRatio             = "price_earnings_ratio"
	PriceToBookValue               = "price_to_book_value"
	DividendYield                  = "dividend_yield"
	BookValuePerShare              = "book_value_per_share"
	ReturnOnEquity                 = "return_on_equity"
	ReturnOnCapitalEmployed        = "return_on_capital_employed"
	FaceValuePerShare              = "face_value_per_share"
	BSECode                        = "bse_code"
	PanCardNumber                  = "pancard_number"
	AggregationType                = "aggregation_type"
	IsDeleted                      = "is_deleted"
	UserId                         = "user_id"
	WatchlistName                  = "watchlist_name"
	MonthYearDate                  = "month_year_date"
	LastUpdatedAt                  = "last_updated_at"
	Metric                         = "metric"
	Value                          = "value"
	Promoters                      = "promoters"
	ForeignInstitutionalInvestors  = "foreign_institutional_investors"
	Others                         = "others"
	PublicInvestors                = "public_investors"
	DomesticInstitutionalInvestors = "domestic_institutional_investors"
	Status                         = "status"
	RetryCount                     = "retry_count"
	ApiName                        = "api_name"
	ReportType                     = "report_type"
)

// Scrip Master column names constants
const (
	ID                         = "id"
	ScripToken                 = "scrip_token"
	Group                      = "group"
	ExchangeSegmentScripMaster = "exchange_segment"
	InstrumentType             = "instrument_type"
	SymbolName                 = "symbol_name"
	TradingSymbol              = "trading_symbol"
	OptionType                 = "option_type"
	UniqueKey                  = "unique_key"
	AssetCode                  = "asset_code"
	LotSize                    = "lot_size"
	TickSize                   = "tick_size"
	Multiplier                 = "multiplier"
	DecimalPrecision           = "decimal_precision"
	StrikePrice                = "strike_price"
	DisplayStrikePrice         = "display_strike_price"
	Description                = "description"
	SubGroup                   = "sub_group"
	AMCCOde                    = "amc_code"
	ContractID                 = "contract_id"
	CombinedScripToken         = "combined_scrip_token"
	ScripReserved              = "scrip_reserved"
	ExchangeSymbolName         = "exchange_symbol_name"
	RMSMarketProtection        = "rms_market_protection"
	ContractType               = "contract_type"
	Exchange                   = "exchange"
	DisplayExpiryDate          = "display_expiry_date"
	SymbolLeg2                 = "symbol_leg2"
)

// Basket Constants
const (
	BasketName = "basket_name"
	BasketID   = "basket_id"
)

// SQL Query Entry Keys
const (
	DeviceIdSQLEntry = "device_id = ?"
	JDataSQLEntry    = "?jData="
	JKeySQLEntry     = "&jKey="
)

// NEST API URL Keys
const (
	OmnenestKey    = "omnenest"
	NestAPITypeKey = "nestapitypetourlmapping"
)

// CMOTS API URL Keys
const (
	CMOTSKey        = "cmots"
	CMOTSAPITypeKey = "cmotsapitypetourlmapping"
)

// Table Names Constants
const (
	UserTable                  = "users"
	DeviceTable                = "devices"
	BrokerTable                = "brokers"
	BrokerWatchlistTable       = "broker_watchlists"
	StaticExchangeConfigTables = "static_exchange_config"
	NewsAndFeedTable           = "news_and_feeds"
	CompanyMasterTable         = "company_master"
	UsersInfoTable             = "users_info"
	BasketOrderTable           = "basket_orders"
	CashFlowTable              = "cash_flow"
	BasketTable                = "baskets"
	BasketOrdersTable          = "basket_orders"
	ProfitAndLossTable         = "profit_and_loss"
	CompanyYearlyRatiosTable   = "company_yearly_ratios"
	ScripMasterTable           = "scrip_master"
	ShareHoldingPatternTable   = "share_holding_pattern"
	WatchlistsTable            = "watchlists"
	WatchlistScripsTable       = "watchlist_scrips"
	UsersInfoTableName         = "users_info"
	CmotsEtlTrackerTable       = "cmots_etl_tracker"
	CmotsFinancialsTable       = "cmots_financials"
	BalanceSheetTable          = "balance_sheet"
	CompanyProfileTable        = "company_profile"
)

// JSON Names
const (
	ExchangeKey        = "exchange"
	ScripTokenKey      = "scripToken"
	TransactionTypeKey = "transactionType"
	QuantityKey        = "quantity"
	TradingSymbolKey   = "tradingSymbol"
)

// NEST API Constants
const (
	VendorCode = ""
	NullString = "NULL"
)

// NEST API short values
const (
	OrderTypeLimitShort  = "L"
	OrderTypeMarketShort = "MKT"
	ActionBuyShort       = "B"
	ActionSellShort      = "S"
)

// Field Names
const (
	Action                      = "Action"
	OrderType                   = "OrderType"
	OrderStatus                 = "OrderStatus"
	InputOrderStatus            = "InputOrderStatus"
	OrderDate                   = "OrderDate"
	ValidityDate                = "ValidityDate"
	TransactionType             = "TransactionType"
	Leg1TransactionType         = "Leg1TransactionType"
	Leg2TransactionType         = "Leg2TransactionType"
	Leg3TransactionType         = "Leg3TransactionType"
	Leg4TransactionType         = "Leg4TransactionType"
	PriceType                   = "PriceType"
	PositionType                = "PositionType"
	ProductCode                 = "ProductCode"
	RetentionType               = "RetentionType"
	ExchangeName                = "ExchangeName"
	OriginalPriceType           = "OriginalPriceType"
	SendAlertsOn                = "SendAlertsOn"
	BFFOrderStatus              = "BFFOrderStatus"
	InstrumentName              = "InstrumentName"
	Option                      = "Option"
	AlertType                   = "AlertType"
	SourceField                 = "Source"
	Segment                     = "Segment"
	ScannerType                 = "ScannerType"
	ScannerTypeValue            = "ScannerTypeValue"
	ConversionTypeField         = "ConversionType"
	CurrentProductCode          = "CurrentProductCode"
	TargetProductCode           = "TargetProductCode"
	AfterMarketOrderFlag        = "AfterMarketOrderFlag"
	IPOStatus                   = "IPOStatus"
	OFSStatus                   = "OFSStatus"
	ExchangeSegment             = "ExchangeSegment"
	ClientSubCategoryCode       = "ClientSubCategoryCode"
	BFFTriggerStatus            = "BFFTriggerStatus"
	PaymentMode                 = "PaymentMode"
	DematAccountNumberFieldName = "DematAccountNumber"
	BFFBidStatus                = "BFFBidStatus"
	BidHistory                  = "BidHistory"
	ExpiryDate                  = "ExpiryDate"
	SubCategory                 = "SubCategory"
	IPOCategory                 = "IPOCategory"
	OFSCategory                 = "OFSCategory"
	Yrc                         = "Yrc"
	BiddingType                 = "BiddingType"
)

// Date and Time Constant
const (
	DateOldLayout1                 = "02-Jan-2006 15:04:05"
	DateOldLayout2                 = "02-01-2006"
	DateOldLayout3                 = "02-Jan-2006"
	DateOldLayout4                 = "02 Jan, 2006"
	DateOldLayout5                 = "2006-01-02"
	DateOldLayout6                 = "02 Jan,2006"
	DateOldLayout7                 = "02/01/2006"
	DateOldLayout8                 = "2 Jan, 2006"
	DateOldLayout9                 = "2-Jan-2006"
	DateOldLayout10                = "02-Jan-06 15:04:05"
	DateOldLayout11                = "02 Jan, 2006 - 15:04:5"
	DateOldLayout12                = "200601"
	DateOldLayout13                = "Y200603"
	DateNewLayout                  = "02/01/2006 15:04:05"
	TimeTrimChars                  = " 00:00:00"
	PrometheusDateTimeFormat       = "06010215"
	PrometheusDateTimeMinuteFormat = "0601021504"
)

// Bff Response Constant Fields
const (
	OrderTypeMarket             = "Market"
	OrderTypeLimit              = "Limit"
	ShortCapsBuy                = "B"
	ShortCapsSell               = "S"
	ShortCapsNil                = "SO"
	ActionBuy                   = "Buy"
	ActionSell                  = "Sell"
	OrderStatusExecuted         = "Executed"
	OrderStatusOpen             = "Open"
	OrderStatusRejected         = "Rejected"
	OrderStatusComplete         = "complete"
	OrderStatusCancelled        = "cancelled"
	OrderStatusOpenStr          = "open"
	OrderStatusPendingStr       = "open pending"
	OrderStatusModifyPendingStr = "modify pending"
	OrderStatusRejectedStr      = "rejected"
	OrderStatusRejectStr        = "reject"
	OrderStatusFailureStr       = "failure"
	SLBMExchange                = "SLBM"
	PositionTypeDay             = "DAY"
)

// Date Format Fields
const (
	DateInField = "date"
)

// Opensearch Constants
const (
	IndexName = "stocks_index"
)

var BFFToNestRequestMapping = map[string]string{

	//Transaction Type Mapping
	"S":            "S",
	"SELL":         "S",
	"B":            "B",
	"BUY":          "B",
	"SO":           "SO",
	"SQUARED-OFF":  "SO",
	"BORROW":       "B",
	"LEND":         "L",
	"RP":           "RP",
	"REPAY":        "RP",
	"RC":           "RC",
	"RECALL":       "RC",
	"SUBSCRIPTION": "B",
	"REDEMPTION":   "S",
	//Specific Transaction Type Mapping
	"FRESH":      "1",
	"ADDITIONAL": "2",
	"SIP":        "3",
	"XSIP":       "4",
	"ISIP":       "5",
	//SIP Security Type Mapping
	"SIPORDER":    "S",
	"NSIPORDER":   "N",
	"ISIPORDER":   "ISIP",
	"XSIPORDER":   "XSIP",
	"SWPORDER":    "SWP",
	"SPREADORDER": "SPREAD",
	//SIP Frequency Number Type Mapping
	"ANNUALLY":     "A",
	"SEMIANNUALLY": "SA",
	"QUARTERLY":    "Q",
	"MONTHLY":      "M",
	"WEEKLY":       "W",
	//Price Type Mapping
	"MARKET":    "MKT",
	"M":         "MKT",
	"MKT":       "MKT",
	"LIMIT":     "L",
	"L":         "L",
	"SL-L":      "SL-L",
	"SL-M":      "SL-M",
	"SP":        "SP",
	"SP-M":      "SP-M",
	"TWO LEG":   "2L",
	"2L":        "2L",
	"THREE LEG": "3L",
	"3L":        "3L",
	"FOUR LEG":  "4L",
	"4L":        "4L",
	//Exchange Name Mapping
	"NSE":     "NSE",
	"BSE":     "BSE",
	"NSE IPO": "NSE IPO",
	"BSE IPO": "BSE IPO",
	"NSE OFS": "NSE OFS",
	"BSE OFS": "BSE OFS",
	"NCDEX":   "NCDEX",
	"NFO":     "NFO",
	"MCX":     "MCX",
	"CDS":     "CDS",
	"ICEX":    "ICEX",
	"NMCE":    "NMCE",
	"DGCX":    "DGCX",
	"MCXSX":   "MCXSX",
	"BFO":     "BFO",
	"NSEL":    "NSEL",
	"MCXSXCM": "MCXSXCM",
	"MCXSXFO": "MCXSXFO",
	"NDM":     "NDM",
	"BCD":     "BCD",
	"BSEMF":   "BSEMF",
	"NCO":     "NCO",
	"BCO":     "BCO",
	"SLBM":    "SLBM",
	//Position Type Mapping
	"DAY": "DAY",
	"NET": "NET",
	//Retention Type Mapping
	"IOC":   "IOC",
	"FOK":   "FOK",
	"GTC":   "GTC",
	"GTD":   "GTD",
	"GTDys": "GTDys",
	"GTDYS": "GTDys",
	"GTT":   "GTT",
	"OPG":   "OPG",
	"EOS":   "EOS",
	"COL":   "COL",
	//Product Code Mapping
	"NRML": "NRML",
	"CNC":  "CNC",
	"MIS":  "MIS",
	"CO":   "CO",
	"BO":   "BO",
	//Order Status Mapping
	"ACCEPTED":                               "accepted",
	"REJECTED":                               "rejected",
	"OPEN":                                   "open",
	"COMPLETE":                               "complete",
	"PUT ORDER REQ RECEIVED":                 "put order req received",
	"VALIDATION PENDING":                     "validation pending",
	"OPEN PENDING":                           "open pending",
	"MODIFY VALIDATION PENDING":              "modify validation pending",
	"MODIFY PENDING":                         "modify pending",
	"MODIFIED":                               "modified",
	"NOT MODIFIED":                           "not modified",
	"CANCEL PENDING":                         "cancel pending",
	"CANCELLED":                              "cancelled",
	"NOT CANCELLED":                          "not cancelled",
	"FROZEN":                                 "frozen",
	"AFTER MARKET ORDER REQ RECEIVED":        "after market order req received",
	"MODIFY AFTER MARKET ORDER REQ RECEIVED": "modify after market order req received",
	"CANCELLED AFTER MARKET ORDER":           "cancelled after market order",
	"LAPSED":                                 "Lapsed",
	"TRIGGER PENDING":                        "trigger pending", //If nest is accepting input as "Trigger pending", it needs to be handled here//
	"MODIFY ORDER REQ RECEIVED":              "pending",         //need to check all nest order status for ipo modification
	//Source Mapping
	"INDXALRT": "indxalrt",
	"SLTP":     "SLTP",
	"TTVOLFD":  "TTVOLFD",
	"TTVALFD":  "TTVALFD",
	"SVWAP":    "SVWAP",
	//ScannerType Mapping
	"OPEN-LOW":       "open_low",
	"OPEN-HIGH":      "open_high",
	"PRICE-SHOCKER":  "ltp",
	"RISING-FALLING": "riseFall",
	//ScannerTypeValue Mapping
	"LOW":                  "low",
	"HIGH":                 "high",
	"52HIGH":               "52-high",
	"52LOW":                "52-low",
	"1%-UPPERCIRCUIT":      "1%uppercircuit",
	"1%-LOWERCIRCUIT":      "1%lowercircuit",
	"UPPERCIRCUIT":         "uppercircuit",
	"LOWERCIRCUIT":         "lowercircuit",
	"PRICEUP-VOLUMEUP":     "priceup_volumeup",
	"PRICEUP-VOLUMEDOWN":   "priceup_volumedown",
	"PRICEDOWN-VOLUMEUP":   "pricedown_volumeup",
	"PRICEDOWN-VOLUMEDOWN": "pricedown_volumedown",
	//AlertType
	"STOCK": "stock",
	"INDEX": "index",
	"ORDER": "order",
	"Y":     "YES",
	"YES":   "YES",
	"N":     "NO",
	"NO":    "NO",
	//Status
	"ONGOING":  "C",
	"CLOSED":   "P",
	"UPCOMING": "F",
	"ALL":      "A",
	//Payment Mode
	"INVESTMENT-ACCOUNT": "IA",
	"DIRECT-PAYMENT":     "DP",
	//ClientSubCategory Mapping
	"SHAREHOLDER":                     "SHA",
	"EMPLOYEE":                        "EMP",
	"POLICY HOLDER":                   "POL",
	"INDIVIDUAL":                      "IND",
	"OTHERS":                          "OTH",
	"CORPORATES":                      "CO",
	"FINANCIAL INSTITUTIONS":          "FI",
	"FOREIGN INSTITUTIONAL INVESTORS": "FII",
	"IC":                              "IC",
	"MUTUAL FUNDS":                    "MF",
	"NOH":                             "NOH",
	//Order Type
	"RETAIL":     "Retail",
	"HNI":        "Non-institutional",
	"NON-RETAIL": "Non-institutional",
}

var NestToBFFResponseMapping = map[string]string{

	//Order Type Mapping
	"L":   "Limit",
	"MKT": "Market",
	//Transaction Type Mapping
	"B":  "Buy",
	"S":  "Sell",
	"SO": "Squared-off",
	//Order Status Mapping
	"complete":                               "Executed",
	"open":                                   "Open",
	"rejected":                               "Rejected",
	"reject":                                 "Rejected",
	"put order req received":                 "ReqRecd",
	"validation pending":                     "InValidation",
	"open pending":                           "OpnPend",
	"modify validation pending":              "ModValidPend",
	"modify pending":                         "ModPending",
	"modified":                               "Modified",
	"not modified":                           "NotModified",
	"cancel pending":                         "CanclPending",
	"cancelled":                              "Cancelled",
	"not cancelled":                          "Not Cancelled",
	"frozen":                                 "Frozen",
	"after market order req received":        "AmoReqRecd",
	"modify after market order req received": "ModAmoReqRecd",
	"modify order req received":              "ModOrdReqRecd",
	"cancelled after market order":           "CancelledAmo",
	"Lapsed":                                 "Lapsed",
	"Trigger pending":                        "TrigPending",
	"trigger pending":                        "TrigPending",
	"pay gate pending":                       "PayGatePend",
	"pay gate approved":                      "PayGateApprd",
	"pay gate rejected":                      "PayGateRejtd",
	"validation complete":                    "Validated",
	"validation rejected":                    "Validatn-Rejtd",
	"amo triggered":                          "AMO-Triggered",
	"offline waiting":                        "OfflineWait",
	"offline modify waiting":                 "OfflineModWait",
	"in transit":                             "InTransit",
	"adapter failure":                        "Adapter Failure",
	"failure":                                "Failure",
	"modify pay gate pending":                "ModPayGatePend",
	"modify pay gate approved":               "ModPayGateApprd",
	"modify pay gate rejected":               "ModPayGateRejtd",
	"modify failure":                         "ModFailed",
	"cancel order req received":              "CanclOrdReqRecd",
	"cancel failure":                         "CanclFailed",
	"asba pending":                           "ASBA-Pend",
	"asba generated":                         "ASBA-Gentd",
	"asba confirmed":                         "ASBA-Conf",
	//Exchange Name Mapping
	"NSE":     "NSE",
	"BSE":     "BSE",
	"NCDEX":   "NCDEX",
	"NFO":     "NFO",
	"MCX":     "MCX",
	"CDS":     "CDS",
	"ICEX":    "ICEX",
	"NMCE":    "NMCE",
	"DGCX":    "DGCX",
	"MCXSX":   "MCXSX",
	"BFO":     "BFO",
	"NSEL":    "NSEL",
	"MCXSXCM": "MCXSXCM",
	"MCXSXFO": "MCXSXFO",
	"NDM":     "NDM",
	"BCD":     "BCD",
	"BSEMF":   "BSEMF",
	"NCO":     "NCO",
	"BCO":     "BCO",
	"SLBM":    "SLBM",
	//Price Type Mapping
	"SL-L": "SL-L",
	"SL-M": "SL-M",
	"SP":   "SP",
	"SP-M": "SP-M",
	"2L":   "TWO LEG",
	"3L":   "THREE LEG",
	"4L":   "FOUR LEG",
	//Product Code Mapping
	"NRML": "NRML",
	"CNC":  "CNC",
	"MIS":  "MIS",
	"CO":   "CO",
	"BO":   "BO",
	//SendAlertsOn Mapping
	"1": "SENTON_EMAIL",
	"2": "SENTON_SMS",
	"3": "SENTON_EMAIL_SMS",
	//AlertType Mapping
	"SECURITY LAST TRADE PRICE":          "stock",
	"TOTAL TRADED VOLUME FOR THE DAY":    "stock",
	"TOTAL TRADED VALUE FOR THE DAY":     "stock",
	"SECURITY VOLUME WEIGHTED AVG PRICE": "stock",
	"INDEX ALERT":                        "index",
	"INDEX VALUE":                        "index",
	//IPO Status
	"UPCOMING": "Upcoming",
	"ONGOING":  "Ongoing",
	"CLOSED":   "Closed",
	// biddingType
	"IPO": "Main Board",
	"SME": "SME",
	"FPO": "FPO",
}

var ClientSubCategoryMap = map[string]string{
	"SHA": "Shareholder",
	"EMP": "Employee",
	"POL": "Policy Holder",
	"IND": "Individual",
	"OTH": "Others",
	"FI":  "Financial Institutions",
	"FII": "Foreign Institutional Investors",
	"IC":  "IC",
	"NOH": "NOH",
	"CO":  "Corporates",
	"MF":  "Mutual Funds",
}

var ScannersTypeMap = map[string]string{
	"low":                  "price-shocker",
	"high":                 "price-shocker",
	"52high":               "price-shocker",
	"52low":                "price-shocker",
	"uppercircuit":         "price-shocker",
	"lowercircuit":         "price-shocker",
	"1%-uppercircuit":      "price-shocker",
	"1%-lowercircuit":      "price-shocker",
	"priceup-volumeup":     "rising-falling",
	"priceup-volumedown":   "rising-falling",
	"pricedown-volumeup":   "rising-falling",
	"pricedown-volumedown": "rising-falling",
}

var ExchangeToExchangeSegmentMapping = map[string]string{
	"NSE":     "nse_cm",
	"BSE":     "bse_cm",
	"NCDEX":   "ncx_fo",
	"NFO":     "nse_fo",
	"MCX":     "mcx_fo",
	"CDS":     "cde_fo",
	"ICEX":    "icx_fo",
	"NMCE":    "nmc_fo",
	"DGCX":    "dgx_fo",
	"MCXSX":   "mcx_sx",
	"BFO":     "bse_fo",
	"NSEL":    "nsel_sm",
	"MCXSXCM": "mcx_cm",
	"MCXSXFO": "mcx_cmfo",
	"NDM":     "nse_dm",
	"BCD":     "bcs_fo",
	"BSEMF":   "bse_mf",
	"NCO":     "nse_com",
	"BCO":     "bse_com",
	"ANY":     "any",
	"SLBM":    "nse_slb",
}

var ExchangeSegmentToExchangeMapping = map[string]string{
	"nse_cm":   "NSE",
	"bse_cm":   "BSE",
	"ncx_fo":   "NCDEX",
	"nse_fo":   "NFO",
	"mcx_fo":   "MCX",
	"cde_fo":   "CDS",
	"icx_fo":   "ICEX",
	"nmc_fo":   "NMCE",
	"dgx_fo":   "DGCX",
	"mcx_sx":   "MCXSX",
	"bse_fo":   "BFO",
	"nsel_sm":  "NSEL",
	"mcx_cm":   "MCXSXCM",
	"mcx_cmfo": "MCXSXFO",
	"nse_dm":   "NDM",
	"bcs_fo":   "BCD",
	"bse_mf":   "BSEMF",
	"nse_com":  "NCO",
	"bse_com":  "BCO",
	"any":      "ANY",
	"nse_slb":  "SLBM",
	"bse_ipo":  "BSE",
	"nse_ipo":  "NSE",
	"bse_ofs":  "BSE",
	"nse_ofs":  "NSE",
}

var ExchangeToInterOpSegementMapping = map[string]string{
	"NSE":     "CASH",
	"BSE":     "CASH",
	"NCDEX":   "COM",
	"NFO":     "FO",
	"MCX":     "COM",
	"CDS":     "CUR",
	"BFO":     "FO",
	"MCXSXCM": "CASH",
	"BCD":     "CUR",
	"NCO":     "COM",
	"BCO":     "COM",
	"CASH":    "CASH",
	"FO":      "FO",
	"CUR":     "CUR",
	"SLBM":    "SLB",
}

var ExchangeToSegmentIndicatorMapping = map[string]string{
	"NSE":     "EQUITY",
	"BSE":     "EQUITY",
	"NCDEX":   "COMMODITY",
	"NFO":     "FNO",
	"MCX":     "COMMODITY",
	"CDS":     "CURRENCY",
	"BFO":     "FNO",
	"MCXSXCM": "EQUITY",
	"BCD":     "CURRENCY",
	"NCO":     "COMMODITY",
	"BCO":     "COMMODITY",
	"CASH":    "EQUITY",
	"FO":      "FNO",
	"CUR":     "CURRENCY",
	"SLBM":    "SLB",
}

var ExchangeToProductCodeMapping = map[string]string{
	"NSE":   "CNC",
	"BSE":   "CNC",
	"NCDEX": "NRML",
	"NFO":   "NRML",
	"MCX":   "NRML",
	"CDS":   "NRML",
	"BFO":   "NRML",
	"BCD":   "NRML",
	"NCO":   "NRML",
	"BCO":   "NRML",
}

// Nest to BFF Transaction Type Mapping for SLBM
var SLBMTransactionTypeMapping = map[string]string{
	"B":  "Borrow",
	"L":  "Lend",
	"RP": "Repay",
	"RC": "Recall",
}

// Nest to BFF IPO Order Type Mapping
var IPOCategoryResponseMapping = map[string]string{
	"RETAIL":            "Retail",
	"NON-INSTITUTIONAL": "HNI",
}

// Nest to BFF OFS Order Type Mapping
var OFSCategoryResponseMapping = map[string]string{
	"RETAIL":            "Retail",
	"NON-INSTITUTIONAL": "Non-Retail",
	"EMP":               "Employee",
}

// Conversion Type mapping
var ConversionTypeMapping = map[string]string{
	"DAY": "D",
	"CF":  "C",
}

// Common Constants
const (
	ExchangeBSE                  = "BSE"
	ExchangeNSE                  = "NSE"
	GroupRS                      = "RS"
	OFSRetailType                = "Retail"
	OFSNonRetailType             = "Non-Retail"
	ExchangeCashInterOp          = "CASH"
	ExchangeFnoInterOp           = "FNO"
	ExchangeCurrencyInterOp      = "CUR"
	SpreadContractIdentifier     = "SP-"
	DefaultDenominatorFloatValue = 1.0
	DefaultMultiplierValue       = 1
)

var InstrumentFilterOut = map[string]bool{
	"CUR":    true,
	"COMDTY": true,
	"UNDCUR": true,
	"UNDIRD": true,
	"UNDIRC": true,
	"UNDIRT": true,
	"UL":     true,
	"AUCSO":  true,
	"INDEX":  true,
	"UNDCOM": true,
}

var EncryptionKeyMiddlewareMapping = map[string]string{
	MiddlewareEncryptionKey: EncryptionKeyMiddlewareDestination + BffPrivateKey,
	AppEncryptionKey:        EncryptionKeyDestinationPath + ClientPrivateKey,
}

const (
	MiddlewareEncryptionKey = "middleware"
	AppEncryptionKey        = "app"
)

var PreferableExchangeToExchangeSegmentMapping = map[string]string{
	"CASH": "nse_cm",
	"FNO":  "nse_fo",
	"CUR":  "cde_fo",
}

const (
	SegmentIndicatorForSpreadInstrumentType = "SPREAD"
)

// Exchange Names
const (
	NFOExchange = "NFO"
	BFOExchange = "BFO"
	CDSExchange = "CDS"
)

// Alerts Field Type mapping
var AlertFieldTypeMapping = map[string]string{
	"SECURITY LAST TRADE PRICE":          "Last Traded Price",
	"TOTAL TRADED VOLUME FOR THE DAY":    "Total Traded Volume",
	"TOTAL TRADED VALUE FOR THE DAY":     "Total Traded Value",
	"SECURITY VOLUME WEIGHTED AVG PRICE": "Volume Weighted Avg Price",
}

// Exchange Segment Indicator
const (
	EquityExchangeSegmentIndicator    = "EQUITY"
	CommodityExchangeSegmentIndicator = "COMMODITY"
	FNOExchangeSegmentIndicator       = "FNO"
	CurrencyExchangeSegmentIndicator  = "CURRENCY"
)

// Product Code
const (
	BOProductCode   = "BO"
	COProductCode   = "CO"
	NRMLProductCode = "NRML"
	CNCProductCode  = "CNC"
	MISProductCode  = "MIS"
	MTFProductCode  = "MTF"
	T5ProductCode   = "T+5"
)

// Prometheus Metrics
const (
	NestOverallLatency  = "NESTOverallLatency"
	CmotsOverallLatency = "CMOTSOverallLatency"
	APIRequestTime      = "APIRequestTime"
	ServiceNameLabel    = "ServiceName"

	HttpRequestMetricLabel  = "HTTP_Requests"
	NestOverallLatencyLabel = "NEST-Overall"
	NestAPILatencyLabel     = "NEST"
	CMOTSLatencyLabel       = "CMOTS"
	BFFLatencyLabel         = "BFF"
	APILatencyLabel         = "Overall"
)

// NEST API URL Keys
const (
	AccountDetailsKey = "accountdetails"
)

// OLTP HTTP API URL Keys
const (
	OLTPHttpEndpointUrl  = "localhost:4318"
	OLTPHttpTimeoutInSec = 10
)

// Application Environment
const (
	DevEnvironment  = "DEV"
	ProdEnvironment = "PROD"
	QAEnvironment   = "QA"
)

// BFF ENC DEC Middleware Skipper
const (
	StocksIntradayAggrDataEndpoint = "/v1/stocks/intraday-aggr-data"
	ChartSource                    = "CHART"
	AutomationFlag                 = "AUTOMATION"
)

var ExchangeToSegmentMappingCheckMargin = map[string]string{
	"BSE":     "CASH",
	"NSE":     "CASH",
	"MCXSXCM": "CASH",
	"NFO":     "FO",
	"NCDEX":   "COM",
	"BCO":     "COM",
	"NCO":     "COM",
	"CDS":     "CUR",
	"MCX":     "COM",
	"ICEX":    "COM",
	"NMCE":    "COM",
	"DGCX":    "COM",
	"MCXSX":   "CUR",
	"BFO":     "FO",
	"NSEL":    "SPOT",
	"MCXSXFO": "FO",
	"BCD":     "CUR",
	"NDM":     "DEBT",
	"SLBM":    "SLBM",
}

// common NEST API URL Keys
const (
	DefaultLoginKey = "defaultlogin"
)

// Equity SIP
const (
	EquitySIPEnabledAttribute = "NT: EquitySIP"
)

// header constants
const (
	MobSource                   = "MOB"
	WebSource                   = "WEB"
	NestHeaderSourceValue       = "nvantage"
	NestHeaderSourcePrefixValue = "NVNTG"
)

// NEST to BFF mapping for GTT trigger types
var FieldTypeMapping = map[string]string{
	"P": "Price",
	"Q": "Quantity",
	"R": "Percentage",
}

// tag constants
const (
	HideFieldTag = "hide"
)

// protos template id constants
const (
	ResponseLoginConnectionTemplateId = 23
	RequestLoginConnectionTemplateId  = int32(22)
	ResponseLoginTemplateId           = 11
	RequestLoginTemplateId            = int32(10)
	ResponseResendOtpTemplateId       = 43
	RequestResendOtpTemplateId        = int32(42)
)

// web socket channel Constants
const (
	ChannelTplantKey = "TPLANT"
	ChannelReqKey    = "REQ"
	ChannelOrdKey    = "ORD"
)

// admin-login Response Codes and admin user password
const (
	SuccessCode = "0"
	FailedCode  = "1"
)

// CMOTS field mapping
var CMOTSMockFieldsMapping = map[string]string{
	"columnname": "COLUMNNAME",
	"rid":        "RID",
}

// microservice names
const (
	MutualFundsMicroservice = "mutualfunds"
)

var InstrumentToContractTypeMapping = map[string]string{
	"OPTIDX":    "OPTION",
	"OPTSTK":    "OPTION",
	"FUTIDX":    "FUTURE",
	"FUTSTK":    "FUTURE",
	"SP-FUTIDX": "SP-FUTURE",
	"SP-FUTSTK": "SP-FUTURE",
	"FUTCUR":    "FUTURE",
	"FUTIRT":    "FUTURE",
	"OPTCUR":    "OPTION",
	"OPTIRD":    "OPTION",
	"SP-FUTCUR": "SP-FUTURE",
	"SP-FUTIRD": "SP-FUTURE",
	"SP-FUTIRT": "SP-FUTURE",
	"SP-OPTCUR": "OPTION",
	"SP-FUTIRC": "SP-FUTURE",
	"SP-FUTCOM": "SP-FUTURE",
	"SP-IF":     "SP-FUTURE",
	"SP-SF":     "SP-FUTURE",
	"SP-IO":     "OPTION",
	"SP-SO":     "OPTION",
	"FUTCOM":    "FUTURE",
	"OPTFUT":    "OPTION",
	"FUTIRC":    "FUTURE",
	"OPTIRC":    "OPTION",
	"IF":        "FUTURE",
	"SF":        "FUTURE",
	"IO":        "OPTION",
	"SO":        "OPTION",
	"FUTBAS":    "FUTURE",
	"FUTBLN":    "FUTURE",
	"FUTENR":    "FUTURE",
	"FUTIRD":    "FUTURE",
	"OPTBLN":    "OPTION",
	"SP-OPTIRD": "OPTION",
}

// Nest QAPI Call-Back constants
const (
	NestQAPICallBackTimeOutValue          = 5   // 5 seconds
	NestQAPICallBackTimeoutForNilResponse = 300 // 300 milliseconds
	NestQAPICallBackGoRoutineTimeOutValue = 100 // 100 milliseconds
	CustomerFirm                          = "C"
	NestSuccessResponse                   = 1
)

// Multiplier Exchange Mappings and Tags
var MultiplierExchangeMapping = map[string]string{
	"nse_cm":  "lotSize",
	"bse_cm":  "lotSize",
	"ncx_fo":  "lotSize",
	"nse_fo":  "lotSize",
	"mcx_fo":  "lotSize",
	"cde_fo":  "multiplier",
	"mcx_sx":  "lotSize",
	"bse_fo":  "lotSize",
	"mcx_cm":  "lotSize",
	"bcs_fo":  "multiplier",
	"nse_com": "TBD",
	"bse_com": "lotSize",
	"nse_slb": "lotSize",
}

const (
	MultiplierTag = "multiplier"
	LotSizeTag    = "lotSize"
)

const (
	GetGTTOrdersParam = "getGTTOrders"
	GetGTTOrdersTrue  = "true"
	GetGTTOrdersFalse = "false"
)

// List of GTC-GTD Exchanges handled by Place Orders
var Exchanges = []string{"MCX", "NCDEX"}

// DB Query Constants
const (
	UniqueIndexForBasketQuery = `
CREATE UNIQUE INDEX IF NOT EXISTS unique_user_basket_name_ci 
ON baskets (user_id, LOWER(basket_name)) 
WHERE is_deleted = false
`
	UniqueIndexForWatchlistQuery = `
CREATE UNIQUE INDEX IF NOT EXISTS unique_user_watchlist_name_ci 
ON watchlists (user_id, LOWER(watchlist_name)) WHERE watchlist_name IS NOT NULL;
`

	UniqueIndexForWatchlistScripsQuery = `
CREATE UNIQUE INDEX IF NOT EXISTS unique_user_watchlist_id_ci 
ON watchlist_scrips (watchlist_id, scrip_id) WHERE watchlist_id IS NOT NULL;
`
	IndexForBasketOrderOnBasketId = `CREATE INDEX idx_basket_orders_basket_id ON basket_orders(basket_id);`

	IndexOnBasketIdAndIsDeleted = `CREATE INDEX idx_basket_orders_basket_status ON basket_orders (basket_id ASC, is_deleted ASC);`

	IndexOnBasketsIdActive = `CREATE INDEX idx_baskets_id_active ON baskets (id ASC) WHERE is_deleted = false;`
)

//indexes

const (
	IndexForScripmasterOnExchangeAndContractType = `CREATE INDEX idx_scrip_master_covering ON scrip_master (exchange ASC, contract_type ASC, instrument_type ASC, contract_id ASC);`
	IndexForScripmaster                          = `CREATE INDEX idx_scrip_master_exchange_expiry_instrument_symbol ON scrip_master (exchange, instrument_type, symbol_name, expiry_date, option_type, strike_price);`
)

const (
	IndexForWatchlistsOnUserId                         = `CREATE INDEX IF NOT EXISTS idx_watchlists_user_id ON watchlists (user_id);`
	IndexForWatchlistsOnUserIdId                       = `CREATE INDEX IF NOT EXISTS idx_watchlists_user_id_id ON watchlists (user_id, id);`
	IndexForWatchlistsOnScripsWatchlistId              = `CREATE INDEX IF NOT EXISTS idx_watchlist_scrips_watchlist_id ON watchlist_scrips (watchlist_id);`
	IndexForWatchlistScripsOnWatchlistIdScripOrder     = `CREATE UNIQUE INDEX IF NOT EXISTS idx_watchlist_scrips_watchlist_id_scrip_id_order ON watchlist_scrips (watchlist_id, scrip_id, "order");`
	IndexForScripMasterOnExchangeNameToken             = `CREATE INDEX IF NOT EXISTS idx_scrip_master_exchange_scrip_token ON scrip_master (exchange, scrip_token);`
	IndexForScripMasterOnExchangeSegmentToken          = `CREATE INDEX IF NOT EXISTS idx_scrip_master_token_exchange ON scrip_master(exchange_segment, scrip_token);`
	IndexForWatchlistsOnUserIdLastUpdatedAt            = `CREATE INDEX IF NOT EXISTS idx_watchlists_user_id_last_updated_at ON watchlists (user_id, last_updated_at);`
	IndexForScripMasterOnExchangeSegementTradingSymbol = `CREATE INDEX IF NOT EXISTS idx_scrip_master_exchange_trading ON scrip_master (exchange_segment, trading_symbol);`
)

const (
	RemoveSipString = "Sip Order Added for SIP id "
)

// Orders Query constants
const (
	PlaceCriteriaAttribute  = "place"
	ModifyCriteriaAttribute = "modify"
	CancelCriteriaAttribute = "cancel"
)

const (
	IntraDayKey = "INTRADAY"
	DeliveryKey = "DELIVERY"
	GTCKey      = "GTC"
	GTDKey      = "GTD"
	GTDYSKey    = "GTDys"
	GTTKey      = "GTT"
	AuctionKey  = "Auction"
)

// OrderList NEST FilterType Constants
const (
	NestOpenStatus      = "open,open pending,validation pending,put order req received,Trigger pending,trigger pending,modify order req received,modify after market order req received,after market order req received,not cancelled,modify pending,modify validation pending,modified,not modified,cancel pending"
	NestCompleteStatus  = "complete"
	NestCancelledStatus = "cancelled"
	NestRejectedStatus  = "rejected"
	NestProductCO       = "CO"
	NestProductBO       = "BO"
	NestProductNRML     = "NRML"
	NestProductCNC      = "CNC"
	NestProductMIS      = "MIS"
	NestProductT5       = "T5"
	NestProductCodeT5   = "T+5"
	NestProductMTF      = "MTF"
)

// Product Code Map
var SpreadIntraDayDeliveryToProductCodeMap = map[string][]string{
	IntraDayKey: {NestProductMIS},
	DeliveryKey: {NestProductNRML},
}

var NormalIntraDayDeliveryToProductCodeMap = map[string][]string{
	IntraDayKey: {NestProductMIS, NestProductBO, NestProductCO},
	DeliveryKey: {NestProductNRML, NestProductCNC, NestProductMTF, NestProductT5, NestProductCodeT5},
	GTTKey:      {NestProductNRML},
}

var AuctionToProductCodeMap = map[string][]string{
	AuctionKey: {NestProductMIS, NestProductNRML, NestProductCNC, NestProductMTF, NestProductT5},
}

var SpreadProductCodeToIntraDayDeliveryMap = map[string]string{
	NestProductMIS:  IntraDayKey,
	NestProductNRML: DeliveryKey,
}

var NormalProductCodeToIntraDayDeliveryMap = map[string]string{
	NestProductMIS:  IntraDayKey,
	NestProductBO:   IntraDayKey,
	NestProductCO:   IntraDayKey,
	NestProductCNC:  DeliveryKey,
	NestProductNRML: DeliveryKey,
	NestProductMTF:  DeliveryKey,
	NestProductT5:   DeliveryKey,
}

var AuctionProductCodeToIntraDayDeliveryMap = map[string]string{
	NestProductMIS:  AuctionKey,
	NestProductCNC:  AuctionKey,
	NestProductNRML: AuctionKey,
	NestProductMTF:  AuctionKey,
	NestProductT5:   AuctionKey,
}

var (
	CNCExchangeList = []string{"NSE", "BSE"}
)

// OrderStatusTab mapping
var OrderStatusTabMapping = map[string]map[string]bool{
	"open":                                   {"OPEN": true},
	"open pending":                           {"OPEN": false},
	"validation pending":                     {"OPEN": false},
	"put order req received":                 {"OPEN": false},
	"Trigger pending":                        {"OPEN": true},
	"trigger pending":                        {"OPEN": true},
	"modify order req received":              {"OPEN": false},
	"modify after market order req received": {"OPEN": true},
	"after market order req received":        {"OPEN": true},
	"not cancelled":                          {"OPEN": true},
	"modify pending":                         {"OPEN": false},
	"modify validation pending":              {"OPEN": false},
	"modified":                               {"OPEN": true},
	"not modified":                           {"OPEN": true},
	"cancel pending":                         {"OPEN": false},
	"complete":                               {"EXECUTED": false},
	"rejected":                               {"ALL": false},
	"cancelled":                              {"ALL": false},
	"cancelled after market order":           {"ALL": false},
	"Lapsed":                                 {"ALL": false},
}

// nest constants
const (
	NestRejectionStatus = "user/user_target not logged in"
	RMSKey              = "rms"
)

// Scrip Group and Subgroup Separator and FilterType 'or' operator
const (
	OrderStringGroupSeparator    = "#"
	OrderStringSubgroupSeparator = "|"
	OrdersFilterTypeComma        = ","
)

// Retention Types Input Mapping
var GetRetentionTypesInputMapping = map[string]string{
	"NSE":     "NSE",
	"BSE":     "BSE",
	"NCDEX":   "NCDEX",
	"NFO":     "NFO",
	"MCX":     "MCX",
	"CDS":     "CDS",
	"ICEX":    "ICEX",
	"NMCE":    "NMCE",
	"DGCX":    "DGCX",
	"MCXSX":   "MCXSX",
	"BFO":     "BFO",
	"NSEL":    "NSEL",
	"MCXSXCM": "MCXSXCM",
	"MCXSXFO": "MCXSXFO",
	"NDM":     "NDM",
	"BCD":     "BCD",
	"BSEMF":   "BSEMF",
	"NCO":     "NCO",
	"BCO":     "BCO",
	"SLBM":    "SLBM",
	"NCD":     "NCD",
	"LIMIT":   "Limit",
	"MARKET":  "Market",
	"SL":      "SL",
	"SL-M":    "SL-M",
	"SP":      "SP",
	"2L":      "2L",
	"3L":      "3L",
	"4L":      "4L",
}

var InstrumentTypeToContractTypeMapping = map[string]string{
	"OPT":   "OPTION",
	"FUT":   "FUTURE",
	"SPFUT": "SP-FUTURE",
}

// OrderStatus to Constants Mappings
var OrderStatusMapping = map[string]string{
	"accepted": "Y",
	"rejected": "N",
}

var ExcludedExchangeSegmentsForNormalOrderList = map[string]bool{
	"any":     true,
	"nse_slb": true,
	"bse_ipo": true,
	"nse_ipo": true,
	"bse_ofs": true,
	"nse_ofs": true,
}

// Define a slice of exchange names that require the default freeze quantity
var ExchangesWithDefaultFreezeQuantity = []string{
	"CDS", "BCD"}

// Title to subtitle mapping
var TitleToSubtitlesBalanceSheetMapping = map[string][]string{
	"Current Liabilities": {
		"Short-term Borrowings", "Lease Liabilities (Current)", "Trade Payables",
		"Other Current Liabilities", "Provisions",
	},
	"Non-Current Liabilities": {
		"Long-term Borrowings", "Lease Liabilities (Non-Current)", "Other Long-term Liabilities",
		"Long-term Provisions", "Deferred Tax Liabilities",
	},
	"Total Equity": {
		"Shareholders' Funds", "Share Capital", "Other Equity", "Minority Interest",
	},
	"Non-Current Assets": {
		"Fixed Assets", "Right-of-Use Assets", "Intangible Assets",
		"Intangible Assets under Development", "Capital Work in Progress",
		"Non-current Investments", "Long-term Loans and Advances",
		"Other Non-Current Assets", "Deferred Tax Assets",
	},
	"Current Assets": {
		"Inventories", "Biological Assets (other than Bearer Plants)",
		"Current Investments", "Cash and Cash Equivalents", "Trade Receivables",
		"Short-term Loans and Advances", "Other Current Assets",
	},
	"Total Equity & Liabilities": {},
	"Total Assets":               {},
}

var CashFlowCMOTSTitleToBFFTitleMapping = map[string]string{
	"CASH FLOWS FROM OPERATING ACTIVITIES": "Operating Activities",
	"CASH FLOWS FROM INVESTING ACTIVITIES": "Investing Activities",
	"CASH FLOWS FROM FINANCING ACTIVITIES": "Financing Activities",
}

// Cash Flow BFF Constants
const (
	BFFNetCashFlowKey = "Net Cash Flow"
)

// Profile and loss CMOTS Constants
const (
	CMOTSSalesOfProductsKey             = "sale of products"
	CMOTSSalesOfServicesKey             = "sale of services"
	CMOTSFinanceCostsKey                = "finance costs"
	CMOTSDepreciationAndAmortizationKey = "depreciation and amortization"
)

// Profile and loss BFF Constants
const (
	BFFSalesKey                       = "Sales"
	BFFFinanceCostsKey                = "Finance Costs"
	BFFDepreciationAndAmortizationKey = "Depreciation and Amortization"
	BFFExpensesKey                    = "Expenses"
	BFFTotalExpensesKey               = "Total Expenses"
)

var ProfitAndLossCMOTSTitleToBFFTitleMapping = map[string]string{
	"total revenue":                       "Total Revenue",
	"total expenses":                      "Total Expenses",
	"operating profit after depreciation": "Operating Profit",
	"profit before tax":                   "Profit Before Tax",
	"profit after tax":                    "Profit After Tax",
	"earning per share - basic":           "Earning Per Share",
	"dividend per share":                  "Dividend Per Share",
}

var CategoryMap = map[int]string{
	0: "QIB",
	1: "HNI",
	2: "Retail",
	3: "Employee",
}
