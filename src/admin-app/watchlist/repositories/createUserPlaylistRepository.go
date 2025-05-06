package repositories

import (
	"context"
	"errors"
	"strings"

	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"
	genericConstants "omnenest-backend/src/constants"

	"gorm.io/gorm"
)

type CreateUserPlaylistRepository interface {
	CreatePlaylistWithSongs(ctx context.Context, db *gorm.DB, condition map[string]interface{}) (models.BFFPlaylistResponse, error)
}

type createUserPlaylistRepository struct{}

func NewCreatePlaylistRepository() CreateUserPlaylistRepository {
	return &createUserPlaylistRepository{}
}

func GetCreateUserPlaylistRepository(useDBMocks bool) CreateUserPlaylistRepository {
	if useDBMocks {
		return &createUserPlaylistRepository{}
	}
	return &createUserPlaylistRepository{}
}

func (repository *createUserPlaylistRepository) CreatePlaylistWithSongs(ctx context.Context, db *gorm.DB, condition map[string]interface{}) (models.BFFPlaylistResponse, error) {

	if db == nil {
		return models.BFFPlaylistResponse{}, errors.New(genericConstants.DatabaseInstanceNilError)
	}

	userID, ok := condition[constants.User_id].(int)
	if !ok || userID <= 0 {
		return models.BFFPlaylistResponse{}, errors.New(constants.UserIDRequiredError)
	}
	songCondition, ok := condition[constants.SongCondition].(map[string]interface{})
	if !ok && len(condition[constants.Song_ids].([]int)) > 0 {
		return models.BFFPlaylistResponse{}, errors.New(constants.SongConditionRequiredError)
	}
	playlistData, ok := condition[constants.PlaylistData].(map[string]interface{})
	if !ok {
		return models.BFFPlaylistResponse{}, errors.New(constants.PlaylistDataRequiredError)
	}
	playlistSongs, _ := condition[constants.PlaylistSongs].([]map[string]interface{})
	name, ok := playlistData[constants.Name].(string)
	if !ok || name == "" {
		return models.BFFPlaylistResponse{}, errors.New(constants.PlaylistNameRequiredError)
	}

	userCondition := map[string]interface{}{
		genericConstants.ID: userID,
	}
	userColumns := []string{genericConstants.ID}
	var user struct {
		ID int
	}
	result := db.WithContext(ctx).Table(constants.UsersTable).Select(userColumns).Where(userCondition).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.BFFPlaylistResponse{}, errors.New(constants.NoUserIdFoundError)
		}
		return models.BFFPlaylistResponse{}, result.Error
	}

	var playlist struct {
		ID int
	}
	if err := db.WithContext(ctx).Table(constants.PlaylistsTable).Create(playlistData).Scan(&playlist).Error; err != nil {
		if strings.Contains(err.Error(), constants.DuplicateKeyViolationError) {
			return models.BFFPlaylistResponse{}, errors.New(constants.DuplicatePlaylistError)
		}
		return models.BFFPlaylistResponse{}, err
	}

	if len(playlistSongs) > 0 {
		var songs []struct {
			ID int
		}
		query := db.WithContext(ctx).Table(constants.SongsTable)
		for key, value := range songCondition {
			query = query.Where(key+" ?", value)
		}
		if err := query.Find(&songs).Error; err != nil {
			return models.BFFPlaylistResponse{}, err
		}
		if len(songs) != len(playlistSongs) {
			return models.BFFPlaylistResponse{}, errors.New(constants.InvalidSongIDsError)
		}

		for _, song := range playlistSongs {
			song[constants.Playlist_id] = playlist.ID
			if err := db.WithContext(ctx).Table(constants.PlaylistSongs).Create(song).Error; err != nil {
				if strings.Contains(err.Error(), constants.DuplicateKeyViolationError) {
					continue
				}
				return models.BFFPlaylistResponse{}, err
			}
		}
	}

	return models.BFFPlaylistResponse{
		SuccessMessage: constants.SuccessfullyCreatedPlaylist,
	}, nil
}
