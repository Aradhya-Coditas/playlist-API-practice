package repositories

import (
	"admin-app/watchlist/commons/constants"
	"context"
	"fmt"
	genericModels "omnenest-backend/src/models"
	"omnenest-backend/src/utils/logger"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AdSongToPlaylistRepository interface {
	ModifyPlaylistSongs(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) error
	GetPlaylistDetails(ctx context.Context, db *gorm.DB, playlistID int) (*genericModels.Playlist, error)
}

type adSongToPlaylistRepository struct{}

func NewAdSongToPlaylistRepository() *adSongToPlaylistRepository {
	return &adSongToPlaylistRepository{}
}

func (r *adSongToPlaylistRepository) ModifyPlaylistSongs(ctx context.Context, db *gorm.DB,conditions map[string]interface{}) error {
	logger := logger.GetLogger(ctx)
	start := time.Now()

	action := conditions["action"].(string)
	playlistID := conditions["playlist_id"].(int)
	songIDs := conditions["song_ids"].([]int)

	switch action {
	case constants.ActionAdd:
		var playlistSongs []genericModels.PlaylistSongs
		for _, songID := range songIDs {
			playlistSongs = append(playlistSongs, genericModels.PlaylistSongs{
				PlaylistID: playlistID,
				SongID:     songID,
			})
		}
		if err := db.WithContext(ctx).Create(&playlistSongs).Error; err != nil {
			return err
		}
	case constants.ActionDelete:
		if err := db.WithContext(ctx).
			Where("playlist_id = ? AND song_id IN ?", playlistID, songIDs).
			Delete(&genericModels.PlaylistSongs{}).Error; err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported action: %s", action)
	}

	logger.Info("ModifyPlaylistSongs completed",
		zap.String("action", action),
		zap.Any("latency", time.Since(start).Milliseconds()))
	return nil
}

func (r *adSongToPlaylistRepository) GetPlaylistDetails(ctx context.Context, db *gorm.DB, playlistID int) (*genericModels.Playlist, error) {
	logger := logger.GetLogger(ctx)
	start := time.Now()

	var playlist genericModels.Playlist
	if err := db.WithContext(ctx).
		Preload("Songs").
		Where("id = ?", playlistID).
		First(&playlist).Error; err != nil {
		return nil, err
	}

	logger.Info("GetPlaylistDetails completed",
		zap.Int("playlistID", playlistID),
		zap.Any("latency", time.Since(start).Milliseconds()))
	return &playlist, nil
}
