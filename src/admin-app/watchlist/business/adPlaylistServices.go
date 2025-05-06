package business

import (
	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"
	"admin-app/watchlist/repositories"
	"context"
	"errors"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/utils/postgres"
	"omnenest-backend/src/utils/tracer"
)

type AdPlaylistService struct {
	repo repositories.AdPlaylistRepository
}

func NewAdPlaylistService(repo repositories.AdPlaylistRepository) *AdPlaylistService {
	return &AdPlaylistService{repo: repo}
}

func (s *AdPlaylistService) ModifyPlaylistSongs(ctx context.Context, spanCtx context.Context, req models.BFFAdPlaylistRequest) (*models.BFFAdPlaylistResponse, error) {
	childSpanCtx, span := tracer.AddToSpan(spanCtx, "ModifyPlaylistSongs")
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	db := postgres.GetPostGresClient().GormDb
	if db == nil {
		return nil, errors.New(genericConstants.DatabaseInstanceNilError)
	}

	playlistCondition := map[string]interface{}{
		"id":      req.PlaylistID,
		"user_id": req.UserID,
	}
	playlist, err := s.repo.GetPlaylist(childSpanCtx, db, playlistCondition)
	if err != nil {
		return nil, err
	}

	songCondition := map[string]interface{}{
		"id": req.SongIDs,
	}
	songs, err := s.repo.GetSongs(childSpanCtx, db, songCondition)
	if err != nil {
		return nil, err
	}
	if len(songs) != len(req.SongIDs) {
		return nil, errors.New(constants.InvalidSongIDsError)
	}

	var responseMessage string
	switch req.Action {
	case "add":
		err = s.repo.AddSongsToPlaylist(childSpanCtx, db, req.PlaylistID, req.SongIDs)
		responseMessage = constants.SuccessfullyAddedSongInPlaylist
	case "delete":
		err = s.repo.DeleteSongsFromPlaylist(childSpanCtx, db, req.PlaylistID, req.SongIDs)
		responseMessage = constants.SuccessfullyDeletedSongFromPlaylist
	default:
		err = errors.New(constants.InvalidActionChoice)
	}

	if err != nil {
		return nil, err
	}

	var songNames []string
	for _, song := range songs {
		songNames = append(songNames, song.Title)
	}

	return &models.BFFAdPlaylistResponse{
		Message:     responseMessage,
		Name:        playlist.Name,
		Description: playlist.Description,
		SongNames:   songNames,
	}, nil
}
