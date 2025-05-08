package models

type BFFAdPlaylistRequest struct {
	Action     string `json:"action" validate:"required,oneof=add delete" example:"add"`
	SongIDs    []int  `json:"songIds" validate:"required" example:"[1,2,3]"`
	PlaylistID int    `json:"playlistId" validate:"required" example:"1"`
	UserID     int    `json:"userId" validate:"required" example:"1"`
}

type BFFAdPlaylistResponse struct {
	Message      string   `json:"message" example:"Successfully added songs to playlist"`
	PlaylistName string   `json:"playlistName" example:"Favourite Songs"`
	SongNames    []string `json:"songNames" example:"[Song Name 1, Song Name 2]"`
	PlaylistID   int      `json:"playlistId" validate:"required" example:"1"`
}
