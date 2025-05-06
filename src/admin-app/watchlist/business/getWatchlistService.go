package business

import (
	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"
	"admin-app/watchlist/repositories"
	"context"
	"errors"
	"fmt"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/postgres"
	"omnenest-backend/src/utils/tracer"
	"sync"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GetWatchlistService struct {
	getWatchlistRepository repositories.GetWatchlistRepository
}

func NewGetWatchlistService(getWatchlistRepository repositories.GetWatchlistRepository) *GetWatchlistService {
	return &GetWatchlistService{
		getWatchlistRepository: getWatchlistRepository,
	}
}

func (service *GetWatchlistService) GetWatchlist(ctx context.Context, spanCtx context.Context, request models.BFFGetWatchlistRequest) ([]models.BFFUserdefine, []models.BFFPredefine, error) {
	log := logger.GetLogger(ctx)
	startTime := time.Now()

	childSpanCtx, span := tracer.AddToSpan(spanCtx, "GetWatchlist")
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
		return nil, nil, fmt.Errorf(genericConstants.DatabaseInstanceNilError)
	}

	userIdCondition := map[string]interface{}{
		genericConstants.ID: request.UserID,
	}
	userIdColumns := []string{genericConstants.ID}
	if err := service.getWatchlistRepository.CheckUserIdExists(childSpanCtx, postgresClient, userIdColumns, userIdCondition); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(constants.NoUserIdFoundError, zap.Any("userId", request.UserID))
			return nil, nil, errors.New(constants.NoUserIdFoundError)
		}
		return nil, nil, err
	}

	brokerIdCondition := map[string]interface{}{
		constants.UserIDKey:   request.UserID,
		constants.BrokerIDKey: request.BrokerID,
	}
	brokerIdColumns := []string{genericConstants.ID}
	if err := service.getWatchlistRepository.CheckBrokerIdExists(childSpanCtx, postgresClient, brokerIdColumns, brokerIdCondition); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(constants.NoBrokerIdFoundError, zap.Any("brokerId", request.BrokerID))
			return nil, nil, errors.New(constants.NoBrokerIdFoundError)
		}
		return nil, nil, err
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var (
		userWatchlists        []models.BFFUserdefine
		predefineWatchlists   []models.BFFPredefine
		userErr, predefineErr error
	)

	columns := []string{
		genericConstants.ID,
		genericConstants.WatchlistName,
	}

	userConditions := map[string]interface{}{
		constants.UserIDKey: request.UserID,
	}

	predefineConditions := map[string]interface{}{
		constants.BrokerIDKey: request.BrokerID,
	}

	go func() {
		defer wg.Done()
		userWatchlistsResult, err := service.getWatchlistRepository.GetUserWatchlist(childSpanCtx, postgresClient, columns, userConditions)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			userErr = fmt.Errorf(constants.UserIdFetchingError, err)
			return
		}
		var result []models.BFFUserdefine
		for _, w := range userWatchlistsResult {
			result = append(result, models.BFFUserdefine{
				Id:            w.Id,
				WatchlistName: w.WatchlistName,
			})
		}
		userWatchlists = result
	}()

	go func() {
		defer wg.Done()
		predefineWatchlistsResult, err := service.getWatchlistRepository.GetBrokerWatchlist(childSpanCtx, postgresClient, columns, predefineConditions)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			predefineErr = fmt.Errorf(constants.BrokerIdFetchingError, err)
			return
		}
		var result []models.BFFPredefine
		for _, w := range predefineWatchlistsResult {
			result = append(result, models.BFFPredefine{
				Id:            w.Id,
				WatchlistName: w.WatchlistName,
			})
		}
		predefineWatchlists = result
	}()

	wg.Wait()

	if userErr != nil {
		return nil, nil, userErr
	}
	if predefineErr != nil {
		return nil, nil, predefineErr
	}

	if len(userWatchlists) == 0 && len(predefineWatchlists) == 0 {
		log.Info(constants.BothWatchlistError)
		return nil, nil, errors.New(constants.BothWatchlistError)
	}

	log.Info(constants.WatchlistAPILog,
		zap.String("message", constants.SuccessfullyFetchedWatchlists),
		zap.Int64("latency", time.Since(startTime).Milliseconds()))

	return userWatchlists, predefineWatchlists, nil
}
