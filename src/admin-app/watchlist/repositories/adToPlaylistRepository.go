package repositories

import (
	"admin-app/watchlist/commons/constants"
	"context"
	"fmt"
	genericModels "omnenest-backend/src/models"

	"gorm.io/gorm"
)

type AdSongToPlaylistRepositories interface {
	AddSongToPlaylist(ctx context.Context, db *gorm.DB, condition map[string]interface{}, playlistSongs []genericModels.PlaylistSongs) error
	DeleteSongToPlaylist(ctx context.Context, db *gorm.DB, condition map[string]interface{}) error
	GetPlaylistSongNames(ctx context.Context, db *gorm.DB, playlistID int) ([]string, error)
}

type adSongToPlaylistRepository struct{}

func NewAdSongToPlaylistRepositories(useDBMocks bool) AdSongToPlaylistRepositories {
	return &adSongToPlaylistRepository{}
}

func (repo *adSongToPlaylistRepository) AddSongToPlaylist(ctx context.Context, db *gorm.DB, condition map[string]interface{}, playlistSongs []genericModels.PlaylistSongs) error {
	return db.WithContext(ctx).Model(&genericModels.PlaylistSongs{}).Create(&playlistSongs).Error
}

func (repo *adSongToPlaylistRepository) DeleteSongToPlaylist(ctx context.Context, db *gorm.DB, condition map[string]interface{}) error {
	return db.WithContext(ctx).Model(&genericModels.PlaylistSongs{}).Where(condition).Delete(&genericModels.PlaylistSongs{}).Error
}

func (repo *adSongToPlaylistRepository) GetPlaylistSongNames(ctx context.Context, db *gorm.DB, playlistID int) ([]string, error) {
	var songNames []string
	query := constants.GetPlaylistSongNamesQuery

	rows, err := db.WithContext(ctx).Raw(query, playlistID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var songName string
		if err := rows.Scan(&songName); err != nil {
			return nil, err
		}
		songNames = append(songNames, songName)
	}

	if len(songNames) == 0 {
		return nil, fmt.Errorf(constants.NoSongsFoundError, playlistID)
	}

	return songNames, nil
}
