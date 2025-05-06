package repositories

import (
	"context"
	"omnenest-backend/src/models"

	"gorm.io/gorm"
)

type ScripMasterRepository interface {
	GetScripMasterByCondition(ctx context.Context, db *gorm.DB, conditions map[string]interface{}, parameters []string) (*models.ScripMaster, error)
}

type scripMasterRepository struct{}

func NewGetScripMasterRepository() *scripMasterRepository {
	return &scripMasterRepository{}
}

type mockScripMasterRepository struct{}

func MockGetScripMasterRepository() *mockScripMasterRepository {
	return &mockScripMasterRepository{}
}

func GetScripMasterRepository(useDBMocks bool) ScripMasterRepository {
	if useDBMocks {
		return MockGetScripMasterRepository()
	}
	return NewGetScripMasterRepository()
}

func (scrip *scripMasterRepository) GetScripMasterByCondition(ctx context.Context, db *gorm.DB, conditions map[string]interface{}, parameters []string) (*models.ScripMaster, error) {
	var scripMaster models.ScripMaster
	result := db.WithContext(ctx).Model(scripMaster).Select(parameters).Where(conditions).Find(&scripMaster)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &scripMaster, nil
}

func (scrip *mockScripMasterRepository) GetScripMasterByCondition(ctx context.Context, db *gorm.DB, conditions map[string]interface{}, parameters []string) (*models.ScripMaster, error) {
	return &models.ScripMaster{}, nil
}
