package models

type BFFDeleteWatchlistRequest struct {
	WatchlistId *uint64 `json:"watchlistId" validate:"required" example:"30"`
}

type BFFDeleteWatchlistResponse struct {
	Message string `json:"message" example:"success"`
}
