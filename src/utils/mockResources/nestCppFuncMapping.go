package mockResources

// admin-app authentication
const (
	ChangePasswordAPIName = "ChangePassword"
)

// admin-app orders
const (
	PlaceEquitySIPOrderAPIName          = "PlaceEquitySIPOrder"
	GetOrdersAPIName                    = "GetOrders"
	GetSpreadOrderDetailsAPIName        = "GetSpreadOrderDetails"
	GetHoldingsAPIName                  = "GetHoldings"
	ReplaceGTTOrderAPIName              = "ReplaceGTTOrder"
	PlaceMultiLegOrderAPIName           = "PlaceMultiLegOrder"
	PlaceOrderAPIName                   = "PlaceOrder"
	CancelGTCGTDOrderAPIName            = "CancelGTCGTDOrder"
	ConfirmOrderAPIName                 = "ConfirmOrder"
	ReplaceGTCGTDOrderAPIName           = "ReplaceGTCGTDOrder"
	PlaceGTTOrderAPIName                = "PlaceGTTOrder"
	CancelGTTOrderAPIName               = "CancelGTTOrder"
	GetCoverOrderLimitsAPIName          = "GetCoverOrderLimits"
	CallBackInternalServerError         = "CallBackInternalServerError"
	UnexpectedCppResponse               = "UnexpectedCppResponse"
	TriggerIdNotFound                   = "TriggerIdNotFound"
	GetTriggerOrdersAPIName             = "GetTriggerOrders"
	GetTradesAPIName                    = "GetTrades"
	NoDataFoundResponse                 = "NoDataFoundResponse"
	ModifyGTTOrderAPIName               = "ModifyGTTOrder"
	ModifyGTTOrderRejectedError         = "ModifyGTTOrderRejectedError"
	ModifyOrderAPIName                  = "ModifyOrder"
	CancelOrderAPIName                  = "CancelOrder"
	MultiCancelOrdersAPIName            = "MultiCancelOrders"
	MultiCancelOrdersWithErrorText      = "MultiCancelOrdersWithErrorText"
	ExitCoverOrderAPIName               = "ExitCoverOrder"
	PlaceGTCGTDOrderAPINAme             = "PlaceGTCGTDOrder"
	ExitCoverOrderRejected              = "ExitCoverOrderRejected"
	MultiExitCoverOrdersAPIName         = "MultiExitCoverOrders"
	MultiExitCoverOrdersWithErrorText   = "MultiExitCoverOrdersWithErrorText"
	ExitBracketOrderAPIName             = "ExitBracketOrder"
	ExitBracketOrderRejected            = "ExitBracketOrderRejected"
	MultiExitBracketOrdersAPIName       = "MultiExitBracketOrders"
	MultiExitBracketOrdersWithErrorText = "MultiExitBracketOrdersWithErrorText"
	CancelBracketOrderAPIName           = "CancelBracketOrder"
	GetGTCGTDOrderDetailsAPIName        = "GetGTCGTDOrderDetails"
	GetGTCGTDOrderDetailsEmptyResponse  = "GetGTCGTDOrderDetailsEmptyResponse"
	GetSpreadOrders                     = "GetSpreadOrders"
	GetSpreadOrdersEmptyResponse        = "GetSpreadOrdersEmptyResponse"
	GetSpreadOrdersSP                   = "GetSpreadOrdersSP"
	GetOrderDetailsAPIName              = "GetOrderDetails"
	ModifySpreadOrderAPIName            = "ModifySpreadOrder"
	PlaceGTCGTDOrderAPIName             = "PlaceGTCGTDOrder"
)

// admin-app watchlist
const (
	GetPredefinedWatchlistData      = "GetPredefinedWatchlistData"
	GetPredefinedWatchlistEmptyData = "GetPredefinedWatchlistEmptyData"
)

// admin-app stocks
const (
	GetStockDetailsAPIName            = "GetStockDetails"
	GetMarketStatusAPIName            = "GetMarketStatus"
	GetMarketStatusWithWarningAPIName = "GetMarketStatusWithWarning"
)

// admin-app profile
const (
	CheckMarginAPIName = "CheckMargin"
)
