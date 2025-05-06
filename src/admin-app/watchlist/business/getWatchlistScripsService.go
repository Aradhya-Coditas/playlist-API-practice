package business

import (
	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"
	"admin-app/watchlist/repositories"
	"context"
	"errors"
	"fmt"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/utils/postgres"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetWatchlistService handles business logic for fetching watchlist scrips.
type GetWatchlistScripsService struct {
	getWatchlistScripsRepository repositories.GetWatchlistScripsRepository
}

// GetWatchListService initializes a new service with the provided repository.
func NewGetWatchlistScripsService(getWatchlistScripsRepository repositories.GetWatchlistScripsRepository) *GetWatchlistScripsService {
	return &GetWatchlistScripsService{
		getWatchlistScripsRepository: getWatchlistScripsRepository,
	}
}

// GetWatchListService initializes a new service with the provided repository
func (service *GetWatchlistScripsService) GetWatchlistScrips(ctx *gin.Context, spanCtx context.Context, bffGetWatchlistScripsRequest models.BFFGetWatchlistScripsRequest) (*models.BFFGetWatchlistScripsResponse, error) {
	client := postgres.GetPostGresClient()
	tx := client.GormDb.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf(genericConstants.DatabaseTransactionBeginError, tx.Error)
	}

	conditions := map[string]interface{}{
		constants.WatchlistId: bffGetWatchlistScripsRequest.WatchlistId,
	}

	exists, hasScrips, err := service.getWatchlistScripsRepository.CheckWatchlistIdExists(spanCtx, tx, conditions)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf(genericConstants.DatabaseQueryError, err)
	}

	if !exists {
		tx.Rollback()
		return nil, fmt.Errorf(genericConstants.InvalidWatchlistIdError, bffGetWatchlistScripsRequest.WatchlistId)
	}

	if !hasScrips {
		tx.Rollback()
		return nil, errors.New(genericConstants.NoWatchlistScripsFoundError)
	}

	columns := []string{constants.DecimalPrecision, constants.Exchange, constants.ExchangeSegement, constants.ExpiryDate,
		constants.ScripToken, constants.StrikePrice, constants.SymbolName, constants.TradingSymbol, constants.UniqueKey,
	}

	scrips, err := service.getWatchlistScripsRepository.GetWatchlistScrips(spanCtx, tx, columns, conditions)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(genericConstants.NoDataFoundError)
		}
		return nil, fmt.Errorf(genericConstants.DatabaseQueryError, err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf(genericConstants.DatabaseTransactionCommitError, err)
	}

	return scrips, nil
}
