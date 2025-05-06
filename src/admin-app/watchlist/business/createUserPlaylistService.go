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

type CreateUserPlaylistService struct {
	createUserPlaylistRepository repositories.CreateUserPlaylistRepository
}

func NewCreateUserPlaylistService(createPlaylistRepository repositories.CreateUserPlaylistRepository) *CreateUserPlaylistService {
	if createPlaylistRepository == nil {
		panic(constants.CreatePlaylistRepositoryNilError)
	}
	return &CreateUserPlaylistService{
		createUserPlaylistRepository: createPlaylistRepository,
	}
}

func (service *CreateUserPlaylistService) CreateUserPlaylist(ctx context.Context, spanCtx context.Context, request models.BFFPlaylistRequest) (models.BFFPlaylistResponse, error) {
	log := logger.GetLogger(ctx)
	startTime := time.Now()

	childSpanCtx, span := tracer.AddToSpan(spanCtx, "CreatePlaylist")
	if span == nil {
		childSpanCtx = spanCtx
	}
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	postgresClient := postgres.GetPostGresClient().GormDb
	if postgresClient == nil {
		log.Error(genericConstants.DatabaseInstanceNilError)
		return models.BFFPlaylistResponse{}, fmt.Errorf(genericConstants.DatabaseInstanceNilError)
	}

	playlistData := map[string]interface{}{
		constants.Name:        request.Name,
		constants.Description: request.Description,
		constants.User_id:     request.UserID,
	}

	var playlistSongs []map[string]interface{}
	for _, songID := range request.SongIDs {
		playlistSong := map[string]interface{}{
			constants.Song_id: songID,
		}
		playlistSongs = append(playlistSongs, playlistSong)
	}

	songCondition := map[string]interface{}{
		constants.ID_IN: request.SongIDs,
	}

	condition := map[string]interface{}{
		constants.User_id:       request.UserID,
		constants.SongCondition: songCondition,
		constants.PlaylistData:  playlistData,
		constants.PlaylistSongs: playlistSongs,
	}

	response, err := service.createUserPlaylistRepository.CreatePlaylistWithSongs(childSpanCtx, postgresClient, condition)
	if err != nil {
		log.Error(constants.PlaylistCreationError, zap.Error(err), zap.Any("request", request))
		return models.BFFPlaylistResponse{}, err
	}

	log.Info(constants.PlaylistAPILog,
		zap.String("message", constants.SuccessfullyCreatedPlaylist),
		zap.Int64("latency", time.Since(startTime).Milliseconds()))

	return response, nil
}
