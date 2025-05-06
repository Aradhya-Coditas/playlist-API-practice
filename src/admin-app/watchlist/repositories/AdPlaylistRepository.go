package repositories

import (
	"admin-app/watchlist/commons/constants"
	"context"
	"errors"
	genericModels "omnenest-backend/src/models"

	"gorm.io/gorm"
)

type AdPlaylistRepository interface {
	GetPlaylist(ctx context.Context, db *gorm.DB, condition map[string]interface{}) (*genericModels.Playlist, error)
	GetSongs(ctx context.Context, db *gorm.DB, condition map[string]interface{}) ([]genericModels.Song, error)
	AddSongsToPlaylist(ctx context.Context, db *gorm.DB, playlistID int, songIDs []int) error
	DeleteSongsFromPlaylist(ctx context.Context, db *gorm.DB, playlistID int, songIDs []int) error
}

type adPlaylistRepository struct{}

func NewAdPlaylistRepository(useDBMocks bool) AdPlaylistRepository {
	if useDBMocks {
		return &adPlaylistRepository{}
	}
	return &adPlaylistRepository{}
}

func (r *adPlaylistRepository) GetPlaylist(ctx context.Context, db *gorm.DB, condition map[string]interface{}) (*genericModels.Playlist, error) {
	var playlist genericModels.Playlist
	err := db.WithContext(ctx).Where(condition).First(&playlist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.PlaylistNotFoundError)
		}
		return nil, err
	}
	return &playlist, nil
}

func (r *adPlaylistRepository) GetSongs(ctx context.Context, db *gorm.DB, condition map[string]interface{}) ([]genericModels.Song, error) {
	var songs []genericModels.Song
	err := db.WithContext(ctx).Where(condition).Find(&songs).Error
	return songs, err
}

func (r *adPlaylistRepository) AddSongsToPlaylist(ctx context.Context, db *gorm.DB, playlistID int, songIDs []int) error {
	var playlistSongs []genericModels.PlaylistSongs
	for _, songID := range songIDs {
		playlistSongs = append(playlistSongs, genericModels.PlaylistSongs{
			PlaylistID: playlistID,
			SongID:     songID,
		})
	}
	return db.WithContext(ctx).Create(&playlistSongs).Error
}

func (r *adPlaylistRepository) DeleteSongsFromPlaylist(ctx context.Context, db *gorm.DB, playlistID int, songIDs []int) error {
	return db.WithContext(ctx).
		Where("playlist_id = ? AND song_id IN ?", playlistID, songIDs).
		Delete(&genericModels.PlaylistSongs{}).Error
}
