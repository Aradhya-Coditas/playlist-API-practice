package business

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"
	"admin-app/watchlist/repositories"

	genericConstants "omnenest-backend/src/constants"
	genericModels "omnenest-backend/src/models"
	"omnenest-backend/src/utils/postgres"
	"omnenest-backend/src/utils/tracer"

	"gorm.io/gorm"
)

type CreateSongPlaylistService struct {
	repo repositories.CreateUserPlaylistRepository
}

func NewCreateSongPlaylistService(repo repositories.CreateUserPlaylistRepository) *CreateSongPlaylistService {
	return &CreateSongPlaylistService{
		repo: repo,
	}
}

func (s *CreateSongPlaylistService) CreatePlaylist(ctx context.Context, spanCtx context.Context, req models.BFFPlaylistRequest) (*int, error) {
	childCtx, span := tracer.AddToSpan(spanCtx, "CreatePlaylistService")
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	db := postgres.GetPostGresClient().GormDb
	if db == nil {
		return nil, fmt.Errorf(genericConstants.DatabaseInstanceNilError)
	}

	playlist := genericModels.Playlist{
		UserID:      req.UserID,
		Name:        req.Name,
		Description: req.Description,
	}

	playlistID, err := s.repo.CreatePlaylist(childCtx, db, playlist)
	if err != nil {
		if strings.Contains(err.Error(), constants.DuplicatePlaylistError) {
			return nil, fmt.Errorf(constants.DuplicatePlaylistError)
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	if len(req.SongIDs) > 0 {
		var playlistSongs []genericModels.PlaylistSongs
		for _, songID := range req.SongIDs {
			playlistSongs = append(playlistSongs, genericModels.PlaylistSongs{
				PlaylistID: *playlistID,
				SongID:     songID,
			})
		}

		if err := s.repo.CreatePlaylistSongs(childCtx, db, playlistSongs); err != nil {
			return nil, err
		}
	}

	return playlistID, nil
}
