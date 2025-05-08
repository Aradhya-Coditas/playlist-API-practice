package business

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"
	"admin-app/watchlist/repositories"
	genericModels "omnenest-backend/src/models"
	"omnenest-backend/src/utils/postgres"
	"omnenest-backend/src/utils/tracer"

	"gorm.io/gorm"
)

type AdSongToPlaylistService struct {
	repo repositories.AdSongToPlaylistRepositories
}

func NewAdSongToPlaylistService(repo repositories.AdSongToPlaylistRepositories) *AdSongToPlaylistService {
	return &AdSongToPlaylistService{
		repo: repo,
	}
}

func (service *AdSongToPlaylistService) AdSongToPlaylist(ctx context.Context, spanCtx context.Context, req models.BFFAdPlaylistRequest) (*models.BFFAdPlaylistResponse, error) {
	childSpanCtx, span := tracer.AddToSpan(ctx, constants.AdSongToPlaylistLog)
	defer span.End()

	db := postgres.GetPostGresClient().GormDb

	action := strings.ToLower(req.Action)
	if action != constants.ActionAdd && action != constants.ActionDelete {
		return nil, fmt.Errorf(constants.InvalidActionChoice)
	}

	var playlistSongs []genericModels.PlaylistSongs
	for _, songID := range req.SongIDs {
		playlistSongs = append(playlistSongs, genericModels.PlaylistSongs{
			PlaylistID: req.PlaylistID,
			SongID:     songID,
		})
	}

	condition := map[string]interface{}{
		constants.Playlist_id: req.PlaylistID,
	}
	switch action {
	case constants.ActionAdd:
		if err := service.repo.AddSongToPlaylist(childSpanCtx, db, condition, playlistSongs); err != nil {
			return nil, service.handleDBError(err)
		}
	case constants.ActionDelete:
		condition[constants.Song_id] = req.SongIDs
		if err := service.repo.DeleteSongToPlaylist(childSpanCtx, db, condition); err != nil {
			return nil, service.handleDBError(err)
		}
	}

	songNames, err := service.repo.GetPlaylistSongNames(childSpanCtx, db, req.PlaylistID)
	if err != nil {
		return nil, err
	}

	response := &models.BFFAdPlaylistResponse{
		Message:    fmt.Sprintf(constants.SuccessfullyAddedDeletedSongsToPlaylist, action),
		PlaylistID: req.PlaylistID,
		SongNames:  songNames,
	}

	return response, nil
}

func (service *AdSongToPlaylistService) handleDBError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf(constants.InvalidPlaylistIdExistError)
	}
	if strings.Contains(err.Error(), constants.DuplicateKeyViolationError) {
		return fmt.Errorf(constants.DuplicatePlaylistError)
	}
	return err
}
