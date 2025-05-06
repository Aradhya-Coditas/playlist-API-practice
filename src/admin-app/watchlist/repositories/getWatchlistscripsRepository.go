package repositories

import (
	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"
	"context"
	"errors"
	"fmt"
	genericModels "omnenest-backend/src/models"
	"omnenest-backend/src/utils/logger"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type getWatchlistScripsRepository struct{}

// GetWatchListScripsRepository defines the interface for retrieving watchlist scripts.
type GetWatchlistScripsRepository interface {
	GetWatchlistScrips(ctx context.Context, db *gorm.DB, columns []string, conditions map[string]interface{}) (*models.BFFGetWatchlistScripsResponse, error)
	CheckWatchlistIdExists(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) (bool, bool, error)
}

// NewGetWatchListScripsRepository returns a new instance of the real repository.
func NewGetWatchlistScripsRepository() *getWatchlistScripsRepository {
	return &getWatchlistScripsRepository{}
}

type mockGetWatchlistScripsRepository struct {
}

func MockGetWatchlistScripsRepository() *mockGetWatchlistScripsRepository {
	return &mockGetWatchlistScripsRepository{}
}

func NewGetWatchiistScripsRepository(useDBMocks bool) GetWatchlistScripsRepository {
	if useDBMocks {
		return MockGetWatchlistScripsRepository()
	}
	return NewGetWatchlistScripsRepository()
}

// CheckWatchlistIdExists verifies if a watchlist ID exists in the database and checks if it contains any scrips.
// Returns (exists, hasScrips, error).
func (gwl *getWatchlistScripsRepository) CheckWatchlistIdExists(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) (bool, bool, error) {
	log := logger.GetLogger(ctx)
	startTime := time.Now()

	var watchlist genericModels.Watchlists

	err := db.WithContext(ctx).Where(conditions[constants.WatchlistId]).First(&watchlist).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, false, nil
		}
		return false, false, fmt.Errorf(constants.FailedToCheckWatchlistIdError, err)
	}

	log.Info(constants.GetWatchlistScripsSuccessMessage,
		zap.Any("result", watchlist),
		zap.Int64("latency", time.Since(startTime).Milliseconds()),
	)
	hasScrips := watchlist.ScripCount > 0

	return true, hasScrips, nil
}

// GetWatchlistScripts executes the raw SQL query to retrieve watchlist scripts based on the provided conditions.
func (gwl *getWatchlistScripsRepository) GetWatchlistScrips(ctx context.Context, db *gorm.DB, columns []string, conditions map[string]interface{}) (*models.BFFGetWatchlistScripsResponse, error) {
	log := logger.GetLogger(ctx)
	startTime := time.Now()
	var watchListScrips []models.WatchlistScrips

	result := db.WithContext(ctx).Model(&genericModels.WatchlistScrips{}).Select(columns).
		Joins(constants.GetWatchlistScripsInnerJoinQuery).Where(conditions).Find(&watchListScrips)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	log.Info(constants.GetWatchlistScripsSuccessMessage,
		zap.Any("result", result),
		zap.Int64("latency", time.Since(startTime).Milliseconds()),
	)

	return &models.BFFGetWatchlistScripsResponse{
		Scrips:     watchListScrips,
		ScripCount: uint8(len(watchListScrips)),
	}, nil
}

func (gwl *mockGetWatchlistScripsRepository) GetWatchlistScrips(ctx context.Context, db *gorm.DB, columns []string, conditions map[string]interface{}) (*models.BFFGetWatchlistScripsResponse, error) {
	return &models.BFFGetWatchlistScripsResponse{}, nil
}

func (gwl *mockGetWatchlistScripsRepository) CheckWatchlistIdExists(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) (bool, bool, error) {
	return true, true, nil
}
