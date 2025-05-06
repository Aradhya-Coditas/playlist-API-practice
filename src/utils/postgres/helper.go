package postgres

import (
	"omnenest-backend/src/models"

	"gorm.io/gorm"
)

// DBClient wraps the DBConnectionClient to allow adding methods
type DBClient struct {
	*models.DBConnectionClient
}

// NewDBClient creates a new DBClient wrapper
func NewDBClient(client *models.DBConnectionClient) *DBClient {
	return &DBClient{client}
}

// CreateRecord creates a new record in the PostgreSQL database.
func (client *DBClient) CreateRecord(record interface{}) *gorm.DB {
	return client.GormDb.Create(record)
}

// ReadRecordByID retrieves a record from the PostgreSQL database by its ID.
func (client *DBClient) ReadRecordByID(record interface{}, id uint) *gorm.DB {
	return client.GormDb.Model(record).First(record, id)
}

// Retrieves all record from the PostgreSQL database.
func (client *DBClient) ReadAllRecords(record interface{}) *gorm.DB {
	return client.GormDb.Model(record).Find(record)
}

// UpdateRecord updates an existing record in the PostgreSQL database with condition.
func (client *DBClient) UpdateRecordWithCondition(record interface{}, conditions map[string]interface{}, updates map[string]interface{}) *gorm.DB {
	return client.GormDb.Model(record).Where(conditions).Updates(updates)
}

// DeleteRecordByID deletes a record from the PostgreSQL database by its ID.
func (client *DBClient) DeleteRecordByID(record interface{}, id uint) *gorm.DB {
	return client.GormDb.Model(record).Delete(record, id)
}

// Retrieves all the record from the PostgreSQL database with condition.
func (client *DBClient) ReadRecordsWithConditions(record interface{}, parameters []string, conditions map[string]interface{}) *gorm.DB {
	return client.GormDb.Model(record).Select(parameters).Where(conditions).Find(record)
}

// Deletes the records with conditions
func (client *DBClient) DeleteRecordsWithConditions(record interface{}, conditions map[string]interface{}) *gorm.DB {
	return client.GormDb.Model(record).Where(conditions).Delete(record)
}

// Counts the records with conditions
func (client *DBClient) CountRecordsWithConditions(record interface{}, count *int64, conditions map[string]interface{}) *gorm.DB {
	return client.GormDb.Model(record).First(record, conditions).Count(count)
}
