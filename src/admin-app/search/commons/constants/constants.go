package constants

// search NEST API URL Keys
const (
	ServiceName      = "search"
	PortDefaultValue = 9099
)

// database columns constants
const (
	ScripId            = "id"
	SpreadOrderType    = "SP"
	InstrumentType     = "instrument_type"
	ContractType       = "contract_type"
	ContractId         = "contract_id"
	Exchange           = "exchange"
	ExchangeSegment    = "exchange_segment"
	SymbolName         = "symbol_name"
	ExpiryDate         = "expiry_date"
	OptionType         = "option_type"
	Group              = "group"
	ScripName          = "description"
	TradingSymbol      = "trading_symbol"
	ScripToken         = "scrip_token"
	ISINValue          = "isin"
	Multiplier         = "multiplier"
	DecimalPrecision   = "decimal_precision"
	TickSize           = "tick_size"
	LotSize            = "lot_size"
	UniqueKey          = "unique_key"
	StrikePrice        = "strike_price"
	CombinedScripToken = "combined_scrip_token"
	SegmentIndicator   = "segment_indicator"
	DisplayStrikePrice = "display_strike_price"
	Description        = "description"
	DisplayExpiryDate  = "display_expiry_date"
)

// contract type constants
const (
	SpreadFuture = "SP-FUTURE"
	Future       = "FUTURE"
	Option       = "OPTION"
)

// query constants
const (
	SearchTextQuery                  = "(\"group\" = ? AND (description ILIKE ? OR trading_symbol ILIKE ?))"
	GetDistinctGroupsByExchangeQuery = `SELECT DISTINCT "group" FROM "scrip_master" WHERE "exchange" = $1`
)

// DB queries
const (
	GetDerivativesInstrumentSelectQuery       = `SELECT DISTINCT "instrument_type","contract_type","contract_id" FROM "scrip_master" WHERE exchange = $1 AND contract_type = $2`
	GetDerivativesScripInformationSelectQuery = `SELECT "id","scrip_token","trading_symbol","lot_size","expiry_date","symbol_name","tick_size","multiplier","decimal_precision","strike_price","combined_scrip_token","unique_key","exchange","instrument_type","option_type","exchange_segment" FROM "scrip_master" WHERE exchange = $1 AND instrument_type = $2 AND symbol_name = $3 AND expiry_date = $4`
	GetDerivativesScripSelectQuery            = `SELECT DISTINCT "symbol_name" FROM "scrip_master" WHERE exchange = $1 AND instrument_type = $2`
	GetDerivativesOptionTypesSelectQuery      = `SELECT DISTINCT "option_type" FROM "scrip_master" WHERE exchange = $1 AND instrument_type = $2 AND symbol_name = $3 AND expiry_date = $4`
	GetDerivativesStrikePriceSelectQuery      = `SELECT DISTINCT "strike_price","display_strike_price" FROM "scrip_master" WHERE exchange = $1 AND instrument_type = $2 AND symbol_name = $3 AND expiry_date = $4 AND option_type = $5`
	SearchEquityScripSelectQuery              = `SELECT "id","exchange","description","trading_symbol","exchange_segment","scrip_token","isin","tick_size","lot_size","unique_key","multiplier","decimal_precision" FROM "scrip_master" WHERE ("group" = $1 AND (description ILIKE $2 OR trading_symbol ILIKE $3))`
	GetDerivativesExpiryDateSelectQuery       = `SELECT DISTINCT "expiry_date","display_expiry_date" FROM "scrip_master" WHERE exchange = $1 AND instrument_type = $2 AND symbol_name = $3 ORDER BY expiry_date ASC`
	GetCircuitLimitsSelectQuery               = `SELECT "multiplier","decimal_precision","expiry_date","combined_scrip_token" FROM "scrip_master" WHERE "exchange_segment" = $1 AND "scrip_token" = $2`
)

const (
	GetDerivativesStrikePriceQueryCondition      = "exchange = ? AND instrument_type = ? AND symbol_name = ? AND expiry_date = ? AND option_type = ?"
	GetDerivativesScripInformationQueryCondition = "exchange = ? AND instrument_type = ? AND symbol_name = ? AND expiry_date = ?"
	GetDerivativesExpiryDateQueryCondition       = "exchange = ? AND instrument_type = ? AND symbol_name = ?"
	GetDerivativesInstrumentQueryCondition       = "exchange = ? AND contract_type = ?"
	GetDerivativesScripQueryCondition            = "exchange = ? AND instrument_type = ?"
	GetDeivativeOptionTypesQueryCondition        = "exchange = ? AND instrument_type = ? AND symbol_name = ? AND expiry_date = ?"
)
