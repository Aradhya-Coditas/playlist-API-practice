package business

import (
	"admin-app/search/commons/constants"
	"admin-app/search/models"
	"admin-app/search/repositories"
	"context"
	"errors"
	"fmt"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/postgres"
	"omnenest-backend/src/utils/tracer"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// SearchEquityGroupService is a business service for searching equity group.
type SearchEquityGroupService struct {
	repositories repositories.SearchEquityGroupRepository
}

// NewSearchEquityGroupService creates a new instance of SearchEquityGroupService.
func NewSearchEquityGroupService(repositories repositories.SearchEquityGroupRepository) *SearchEquityGroupService {
	return &SearchEquityGroupService{
		repositories: repositories,
	}
}

// SearchEquityGroup searches equity group by exchange segment.
func (service *SearchEquityGroupService) SearchEquityGroup(ctx context.Context, spanCtx context.Context, Exchange string) (*models.BFFSearchEquityGroupResponse, error) {
	childSpanCtx, span := tracer.AddToSpan(spanCtx, "SearchEquityGroup")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	client := postgres.GetPostGresClient()
	log := logger.GetLogger(ctx)

	columns := []string{
		constants.Group,
	}
	conditions := map[string]interface{}{
		constants.Exchange: Exchange,
	}

	startTime := time.Now()
	equityGroups, err := service.repositories.ReadRecordsWithConditions(childSpanCtx, client.GormDb, columns, conditions)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New(genericConstants.NoDataFoundError)
		}
		return nil, fmt.Errorf(genericConstants.DatabaseQueryError, err)
	}

	log.Info(genericConstants.DBQueryConfig,
		zap.Any(genericConstants.DBConditionsConfig, len(equityGroups)),
		zap.Int64(genericConstants.LatencyKey, time.Since(startTime).Milliseconds()))

	var nonEmptyGroups []string
	for _, group := range equityGroups {
		trimmedGroup := strings.TrimSpace(group)
		if trimmedGroup != genericConstants.EmptySpace {
			nonEmptyGroups = append(nonEmptyGroups, trimmedGroup)
		}
	}

	if len(nonEmptyGroups) == 0 {
		return nil, errors.New(genericConstants.NoDataFoundError)
	}

	return &models.BFFSearchEquityGroupResponse{
		Groups: nonEmptyGroups,
	}, nil
}
