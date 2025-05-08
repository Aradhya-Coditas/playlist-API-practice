package repositories

import (
	"context"
	"fmt"
	genericModels "omnenest-backend/src/models"
	"strings"

	"admin-app/watchlist/commons/constants"

	"gorm.io/gorm"
)

type createUserPlaylistRepository struct{}

type CreateUserPlaylistRepository interface {
	CreatePlaylist(ctx context.Context, db *gorm.DB, playlist genericModels.Playlist) (*int, error)
	CreatePlaylistSongs(ctx context.Context, db *gorm.DB, playlistSongs []genericModels.PlaylistSongs) error
}

func NewCreateUserPlaylistRepository(useDBMocks bool) CreateUserPlaylistRepository {
	if useDBMocks {
		return &createUserPlaylistRepository{}
	}
	return &createUserPlaylistRepository{}
}

func (r *createUserPlaylistRepository) CreatePlaylist(ctx context.Context, db *gorm.DB, playlist genericModels.Playlist) (*int, error) {
	result := db.WithContext(ctx).Table(constants.PlaylistsTable).Create(&playlist)

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf(constants.NoRowsAffectedError)
	}
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), constants.DuplicateKeyViolationError) {
			return nil, fmt.Errorf(constants.DuplicatePlaylistError)
		}
		return nil, result.Error
	}

	return &playlist.ID, nil
}

func (r *createUserPlaylistRepository) CreatePlaylistSongs(ctx context.Context, db *gorm.DB, playlistSongs []genericModels.PlaylistSongs) error {
	if len(playlistSongs) == 0 {
		return nil
	}
	result := db.WithContext(ctx).Table(constants.PlaylistSongs).Create(&playlistSongs)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), constants.DuplicateKeyViolationError) {
			return fmt.Errorf(constants.DuplicatePlaylistError)
		}
		return result.Error
	}
	return nil
}
