package repositories

import (
	"admin-app/watchlist/commons/constants"
	"context"
	genericModels "omnenest-backend/src/models"
	"omnenest-backend/src/utils/logger"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type getWatchlistRepository struct{}

type GetWatchlistRepository interface {
	CheckUserIdExists(ctx context.Context, db *gorm.DB, columns []string, condition map[string]interface{}) error
	CheckBrokerIdExists(ctx context.Context, db *gorm.DB, columns []string, condition map[string]interface{}) error
	GetUserWatchlist(ctx context.Context, db *gorm.DB, columns []string, condition map[string]interface{}) ([]genericModels.Watchlists, error)
	GetBrokerWatchlist(ctx context.Context, db *gorm.DB, columns []string, condition map[string]interface{}) ([]genericModels.BrokerWatchlists, error)
}

func NewGetWatchlistRepository() *getWatchlistRepository {
	return &getWatchlistRepository{}
}

type mockGetWatchlistRepository struct{}

func MockNewGetWatchlistRepository() *mockGetWatchlistRepository {
	return &mockGetWatchlistRepository{}
}

func GetGetWatchlistRepository(useDBMocks bool) GetWatchlistRepository {
	if useDBMocks {
		return MockNewGetWatchlistRepository()
	}
	return NewGetWatchlistRepository()
}

func (wl *getWatchlistRepository) CheckUserIdExists(ctx context.Context, db *gorm.DB, columns []string, condition map[string]interface{}) error {
	log := logger.GetLogger(ctx)
	startTime := time.Now()

	var user genericModels.UsersInfo
	result := db.WithContext(ctx).Model(&user).Select(columns).Debug().Where(condition).First(&user)

	log.Info(constants.CheckUserIdAPILog,
		zap.Any("Result", result),
		zap.Int64("LatencyKey", time.Since(startTime).Milliseconds()))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (wl *getWatchlistRepository) CheckBrokerIdExists(ctx context.Context, db *gorm.DB, columns []string, condition map[string]interface{}) error {
	log := logger.GetLogger(ctx)
	startTime := time.Now()

	var count int64
	result := db.WithContext(ctx).Model(&genericModels.UsersInfo{}).Joins(constants.BrokerIdJoinQuery).Debug().Where(constants.BrokerIdWhereQuery, condition[constants.UserIDKey], condition[constants.BrokerIDKey]).Count(&count)

	if result.Error != nil {
		return result.Error
	}

	if count == 0 {
		return gorm.ErrRecordNotFound
	}

	log.Info(constants.CheckBrokerIdAPILog,
		zap.Any("Result", result),
		zap.Any("LatencyKey", time.Since(startTime).Milliseconds()),
	)
	return nil
}

func (wl *getWatchlistRepository) GetUserWatchlist(ctx context.Context, db *gorm.DB, columns []string, condition map[string]interface{}) ([]genericModels.Watchlists, error) {
	log := logger.GetLogger(ctx)
	startTime := time.Now()

	var watchlists []genericModels.Watchlists
	result := db.WithContext(ctx).Model(&genericModels.Watchlists{}).Select(columns).Debug().Where(condition).Find(&watchlists)

	log.Info(constants.UserWatchlistAPILog,
		zap.Any("Result", result),
		zap.Int64("LatencyKey", time.Since(startTime).Milliseconds()))

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return watchlists, nil
}

func (wl *getWatchlistRepository) GetBrokerWatchlist(ctx context.Context, db *gorm.DB, columns []string, condition map[string]interface{}) ([]genericModels.BrokerWatchlists, error) {
	log := logger.GetLogger(ctx)
	startTime := time.Now()

	var watchlists []genericModels.BrokerWatchlists
	result := db.WithContext(ctx).Model(&genericModels.BrokerWatchlists{}).Select(columns).Debug().Where(condition).Find(&watchlists)

	if result.Error != nil {
		log.Error(constants.BrokerIdFetchingError, zap.Error(result.Error))
		return nil, result.Error
	}

	log.Info(constants.BrokerWatchlistAPILog,
		zap.Any("Result", result),
		zap.Int64("LatencyKey", time.Since(startTime).Milliseconds()),
		zap.Int64("RecordsFound", result.RowsAffected))

	return watchlists, nil
}

func (wl *mockGetWatchlistRepository) CheckUserIdExists(ctx context.Context, db *gorm.DB, columns []string, condition map[string]interface{}) error {
	return nil
}

func (wl *mockGetWatchlistRepository) CheckBrokerIdExists(ctx context.Context, db *gorm.DB, columns []string, condition map[string]interface{}) error {
	return nil
}

func (wl *mockGetWatchlistRepository) GetUserWatchlist(ctx context.Context, db *gorm.DB, columns []string, condition map[string]interface{}) ([]genericModels.Watchlists, error) {
	return nil, nil
}

func (wl *mockGetWatchlistRepository) GetBrokerWatchlist(ctx context.Context, db *gorm.DB, columns []string, condition map[string]interface{}) ([]genericModels.BrokerWatchlists, error) {
	return nil, nil
}
