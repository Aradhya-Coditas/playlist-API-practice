package configs

import (
	"context"
	"omnenest-backend/src/constants"
	"omnenest-backend/src/models"
)

var multiCastConfig *models.MultiCastConfig

func GetMultiCastConfig() *models.MultiCastConfig {
	return multiCastConfig
}

func InitMultiCastConfig(ctx context.Context) error {
	multiCastViperConfig, err := Get(constants.MultiCastConfig)
	if err != nil {
		return err
	}

	multiCastConfig = &models.MultiCastConfig{
		Host:            getEnv(constants.MulticastIPEnvKey, multiCastViperConfig.GetString(constants.MulticastIPKey)),
		Port:            getEnvInt(constants.MulticastPortEnvKey, multiCastViperConfig.GetInt(constants.MulticastPortKey)),
		UdpVersion:      multiCastViperConfig.GetString(constants.MulticastUdpVersion),
		MaxDatagramSize: multiCastViperConfig.GetInt(constants.MaxDatagramSizeKey),
	}
	return nil
}
