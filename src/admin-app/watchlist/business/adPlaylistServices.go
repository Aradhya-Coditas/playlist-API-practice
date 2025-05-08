package business

import (
	"context"
	"fmt"
	"time"

	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"
	"admin-app/watchlist/repositories"

	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/postgres"
	"omnenest-backend/src/utils/tracer"

	"go.uber.org/zap"
)

type AdSongToPlaylistService struct {
	repo repositories.AdSongToPlaylistRepository
}

func NewAdSongToPlaylistService(repo repositories.AdSongToPlaylistRepository) *AdSongToPlaylistService {
	return &AdSongToPlaylistService{
		repo: repo,
	}
}

func (s *AdSongToPlaylistService) ModifyPlaylistSongs(ctx context.Context, spanCtx context.Context, req models.BFFAdPlaylistRequest) (*models.BFFAdPlaylistResponse, error) {
	childCtx, span := tracer.AddToSpan(spanCtx, "ModifyPlaylistSongsService")
	defer span.End()

	db := postgres.GetPostGresClient().GormDb
	if db == nil {
		return nil, fmt.Errorf(genericConstants.DatabaseInstanceNilError)
	}

	logger := logger.GetLogger(ctx)
	start := time.Now()

	var conditions map[string]interface{}
	switch req.Action {
	case constants.ActionAdd:
		conditions = map[string]interface{}{
			"action":      constants.ActionAdd,
			"playlist_id": req.PlaylistID,
			"song_ids":    req.SongIDs,
		}
	case constants.ActionDelete:
		conditions = map[string]interface{}{
			"action":      constants.ActionDelete,
			"playlist_id": req.PlaylistID,
			"song_ids":    req.SongIDs,
		}
	default:
		return nil, fmt.Errorf("invalid action: %s", req.Action)
	}

	if err := s.repo.ModifyPlaylistSongs(childCtx, db, conditions); err != nil {
		logger.Error("Error modifying playlist songs", zap.Error(err))
		return nil, fmt.Errorf("failed to modify playlist songs: %w", err)
	}

	playlist, err := s.repo.GetPlaylistDetails(childCtx, db, req.PlaylistID)
	if err != nil {
		logger.Error("Error fetching playlist details", zap.Error(err))
		return nil, fmt.Errorf("failed to fetch playlist details: %w", err)
	}

	var songNames []string
	for _, song := range playlist.Songs {
		songNames = append(songNames, song.Title)
	}

	response := &models.BFFAdPlaylistResponse{
		PlaylistID:   playlist.ID,
		PlaylistName: playlist.Name,
		SongNames:    songNames,
	}

	logger.Info("Successfully modified playlist songs",
		zap.Int("playlistID", req.PlaylistID),
		zap.Any("latency", time.Since(start).Milliseconds()))

	return response, nil
}
