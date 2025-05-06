package models

type BFFAdPlaylistRequest struct {
	Action     string `json:"action" validate:"required,oneof=add delete"`
	SongIDs    []int  `json:"song_ids" validate:"required"`
	PlaylistID int    `json:"playlist_id" validate:"required"`
	UserID     int    `json:"user_id" validate:"required"`
}

type BFFAdPlaylistResponse struct {
	Message     string   `json:"message"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	SongNames   []string `json:"song_names"`
}
