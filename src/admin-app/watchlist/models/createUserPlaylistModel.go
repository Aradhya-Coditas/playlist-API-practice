package models

type BFFPlaylistRequest struct {
	Name        string `json:"name" validate:"required" example:"Playlist Name"`
	Description string `json:"description" example:"Playlist Description"`
	SongIDs     []int  `json:"songIds" example:"[1,2,3]"`
	UserID      int    `json:"userId" validate:"required" example:"1"`
}

type BFFPlaylistResponse struct {
	Message string `json:"message" example:"Successfully created playlist"`
}