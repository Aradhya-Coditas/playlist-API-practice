package models

type BFFGetWatchlistScripsRequest struct {
	WatchlistId uint64 `json:"watchlistId" validate:"required" example:"30"`
}

type BFFGetWatchlistScripsResponse struct {
	Scrips     []WatchlistScrips `json:"scrips" example:"result"`
	ScripCount uint8             `json:"scripsCount" example:"0"`
}

type WatchlistScrips struct {
	DecimalPrecision uint16  `json:"decimalPrecision" example:"5"`
	Exchange         string  `json:"exchange" example:"NFO"`
	ExchangeSegment  string  `json:"exchangeSegment" example:"nse_fo"`
	ExpiryDate       string  `json:"expiryDate" example:"28/12/2023 00:00:00"`
	ScripToken       string  `json:"scripToken" example:"61627"`
	StrikePrice      float32 `json:"strikePrice" example:"0"`
	SymbolName       string  `json:"symbolName" example:"BANKNIFTY"`
	TradingSymbol    string  `json:"tradingSymbol" example:"BANKNIFTY23DECFUT"`
	UniqueKey        string  `json:"uniqueKey" example:"BANKNIFTY23DECFUT"`
}
