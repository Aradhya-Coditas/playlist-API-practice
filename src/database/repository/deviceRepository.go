package repositories

import (
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/models"

	"gorm.io/gorm"
)

type DeviceRepository interface {
	GetDeviceByDeviceID(db *gorm.DB, deviceId string, selectedColumns ...string) (*models.Devices, bool, error)
}

type deviceRepository struct{}

func NewDeviceRepository() *deviceRepository {
	return &deviceRepository{}
}

// GetDeviceByDeviceID retrieves selected columns for a device with the given deviceId.
func (dr *deviceRepository) GetDeviceByDeviceID(db *gorm.DB, deviceId string, selectedColumns ...string) (*models.Devices, bool, error) {
	var device models.Devices
	result := db.Select(selectedColumns).Where(genericConstants.DeviceIdSQLEntry, deviceId).First(&device)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, false, nil
		}
		return nil, false, result.Error
	}
	return &device, true, nil
}
