package models

type BFFGetWatchlistRequest struct {
	UserID   *uint64 `json:"userId" validate:"required" example:"5"`
	BrokerID *uint64 `json:"brokerId" validate:"required" example:"1"`
}

type BFFWatchlistResponse struct {
	Userdefine []BFFUserdefine `json:"userdefine"`
	Predefine  []BFFPredefine  `json:"predefine"`
}

// BFFUserdefine represents the user-defined watchlist details
type BFFUserdefine struct {
	Id uint64 `json:"id"`
	WatchlistName string `json:"watchlistName"`
}

// BFFPredefine represents the predefined watchlist details
type BFFPredefine struct {
	Id uint64 `json:"id"`
	WatchlistName string `json:"watchlistName"`
}
