package repositories

import (
	"admin-app/watchlist/commons/constants"
	"context"
	"time"

	genericConstants "omnenest-backend/src/constants"
	genericModels "omnenest-backend/src/models"
	"omnenest-backend/src/utils/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type deleteWatchlistRepository struct{}

type DeleteWatchlistRepository interface {
	DeleteWatchlist(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) error
	DeleteWatchlistScrips(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) error
}

func NewDeleteWatchlistRepository() *deleteWatchlistRepository {
	return &deleteWatchlistRepository{}
}

type mockDeleteWatchlistRepository struct {
}

func MockNewDeleteWatchlistRepository() *mockDeleteWatchlistRepository {
	return &mockDeleteWatchlistRepository{}
}

func GetDeleteWatchlistRepository(useDBMocks bool) DeleteWatchlistRepository {
	if useDBMocks {
		return MockNewDeleteWatchlistRepository()
	}
	return NewDeleteWatchlistRepository()
}

// DeleteWatchlistScrips deletes the scrips entry from the database based on given conditiions.
func (wl *deleteWatchlistRepository) DeleteWatchlistScrips(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) error {
	log := logger.GetLogger(ctx)
	startTime := time.Now()
	result := db.WithContext(ctx).
		Exec(constants.DeleteWatchlistAndScripsQuery,
			conditions[constants.UserId],
			conditions[constants.WatchlistId],
		)

	log.Info(constants.WatchlistAPILog,
		zap.Any(genericConstants.DBRecordsFoundConfig, result.RowsAffected),
		zap.Int64(genericConstants.LatencyKey, time.Since(startTime).Milliseconds()))
	return result.Error
}

// DeleteWatchlist deletes the watchlist entry from the database based on given conditiions.
func (wl *deleteWatchlistRepository) DeleteWatchlist(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) error {
	log := logger.GetLogger(ctx)
	startTime := time.Now()

	result := db.WithContext(ctx).Model(&genericModels.Watchlists{}).Where(conditions).Delete(&genericModels.Watchlists{})
	if result.RowsAffected == genericConstants.ZeroNumericValue {
		return gorm.ErrRecordNotFound
	}

	log.Info(constants.WatchlistAPILog,
		zap.Any(genericConstants.DBRecordsFoundConfig, result.RowsAffected),
		zap.Int64(genericConstants.LatencyKey, time.Since(startTime).Milliseconds()))
	return result.Error
}

// DeleteWatchlistScrips deletes the scrips entry from the database based on given conditiions.
func (wl *mockDeleteWatchlistRepository) DeleteWatchlistScrips(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) error {
	return nil

}

// DeleteWatchlist deletes the watchlist entry from the database based on given conditiions.
func (wl *mockDeleteWatchlistRepository) DeleteWatchlist(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) error {
	return nil
}
