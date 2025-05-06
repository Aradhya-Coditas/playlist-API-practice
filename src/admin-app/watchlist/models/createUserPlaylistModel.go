package models

type BFFPlaylistRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	SongIDs     []int  `json:"song_ids"`
	UserID      int    `json:"user_id" validate:"required"`
}

type BFFPlaylistResponse struct {
	SuccessMessage string `json:"message"`
}