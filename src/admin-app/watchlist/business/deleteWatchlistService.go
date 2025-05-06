package business

import (
	"context"
	"errors"
	"fmt"

	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"
	"admin-app/watchlist/repositories"
	genericConstants "omnenest-backend/src/constants"

	"omnenest-backend/src/utils/postgres"
	"omnenest-backend/src/utils/tracer"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeleteWatchlistService represents a service to delete a watchlist.
type DeleteWatchlistService struct {
	deleteWatchlistRepository repositories.DeleteWatchlistRepository
}

// NewDeleteWatchlistService creates a new instance of DeleteWatchlistService.
func NewDeleteWatchlistService(deleteWatchlistRepository repositories.DeleteWatchlistRepository) *DeleteWatchlistService {
	return &DeleteWatchlistService{
		deleteWatchlistRepository: deleteWatchlistRepository,
	}
}

// DeleteWatchlist is a method to delete a watchlist.
func (service *DeleteWatchlistService) DeleteWatchlist(ctx *gin.Context, spanCtx context.Context, bffDeleteWatchlistRequest models.BFFDeleteWatchlistRequest) error {
	childSpanCtx, span := tracer.AddToSpan(spanCtx, "DeleteWatchlist")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	client := postgres.GetPostGresClient()
	userId := ctx.Value(genericConstants.UserID).(uint64)
	// Begin the transaction for deleting the basket and its orders
	tx := client.GormDb.Begin()
	if tx.Error != nil {
		return fmt.Errorf(genericConstants.DatabaseTransactionBeginError, tx.Error)
	}

	conditions := map[string]interface{}{
		constants.WatchlistId: bffDeleteWatchlistRequest.WatchlistId,
		constants.UserId:      userId,
	}

	// Delete associated scrips first to satisfy foreign key constraint
	if err := service.deleteWatchlistRepository.DeleteWatchlistScrips(childSpanCtx, tx, conditions); err != nil {
		tx.Rollback()
		return fmt.Errorf(genericConstants.DatabaseQueryError, err)
	}

	conditions = map[string]interface{}{
		genericConstants.Id: bffDeleteWatchlistRequest.WatchlistId,
		constants.UserId:    userId,
	}

	if err := service.deleteWatchlistRepository.DeleteWatchlist(childSpanCtx, tx, conditions); err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(genericConstants.NoDataFoundError)

		}
		return fmt.Errorf(genericConstants.DatabaseQueryError, err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf(genericConstants.DatabaseTransactionCommitError, err)
	}

	return nil
}
