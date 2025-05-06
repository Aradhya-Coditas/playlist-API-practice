package models

import "time"

// Only used for Swagger generation
type ErrorMessage struct {
	Key          string `json:"key,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type ErrorAPIResponse struct {
	Message      []ErrorMessage `json:"errors,omitempty"`
	ErrorMessage string         `json:"error,omitempty"`
}

type BFFScrip struct {
	ScripToken   string `json:"scripToken" example:"8075"`
	ExchangeName string `json:"exchange" example:"NSE"`
}

type WatchlistWithId struct {
	WatchlistName string    `json:"watchlistName" example:"test1"`
	WatchlistId   uint64    `json:"watchlistId" example:"30"`
	LastUpdatedAt time.Time `json:"-" example:"15/11/2023 11:14:48"`
}
