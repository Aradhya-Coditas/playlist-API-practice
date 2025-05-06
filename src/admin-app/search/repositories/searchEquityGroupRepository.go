package repositories

import (
	"context"
	genericModels "omnenest-backend/src/models"

	"gorm.io/gorm"
)

type SearchEquityGroupRepository interface {
	ReadRecordsWithConditions(ctx context.Context, db *gorm.DB, columns []string, conditions map[string]interface{}) ([]string, error)
}

type getSearchEquityGroupRepository struct{}

func NewSearchEquityGroupRepository() *getSearchEquityGroupRepository {
	return &getSearchEquityGroupRepository{}
}

func MockNewSearchEquityGroupRepository() *getSearchEquityGroupRepository {
	return &getSearchEquityGroupRepository{}
}

func GetSearchEquityGroupRepository(useDBMocks bool) SearchEquityGroupRepository {
	if useDBMocks {
		return MockNewSearchEquityGroupRepository()
	}
	return NewSearchEquityGroupRepository()
}

func (repository *getSearchEquityGroupRepository) ReadRecordsWithConditions(ctx context.Context, db *gorm.DB, columns []string, conditions map[string]interface{}) ([]string, error) {
	var equityGroups []string
	result := db.WithContext(ctx).
		Model(&genericModels.ScripMaster{}).
		Distinct().
		Select(columns).
		Where(conditions).
		Find(&equityGroups)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return equityGroups, nil
}
