package configs

import (
	"context"
	"omnenest-backend/src/constants"
	"omnenest-backend/src/models"
)

var nestIniConfig *models.NestIniConfig

func InitNestIniConfigs(ctx context.Context) error {
	// Get a Apis config from yaml
	applicationViperConfig, err := Get(constants.NestIniConfig)
	if err != nil {
		return err
	}

	nestIniConfig = &models.NestIniConfig{
		NestEnvSettings: models.NestEnvSettings{
			MmlLocBrokAddr: getEnv(constants.MmlLocBrokAddrEnvKey, applicationViperConfig.GetString(constants.MmlLocBrokAddrKey)),
			MmlDmnSrvrAddr: getEnv(constants.MmlDmnSrvrAddrEnvKey, applicationViperConfig.GetString(constants.MmlDmnSrvrAddrKey)),
			MmlDsFoAddr:    getEnv(constants.MmlDsFoAddrEnvKey, applicationViperConfig.GetString(constants.MmlDsFoAddrKey)),
			MmlLicSrvrAddr: getEnv(constants.MmlLicSrvrAddrEnvKey, applicationViperConfig.GetString(constants.MmlLicSrvrAddrKey)),
		},
		AdminName: models.AdminName{
			AdminName: getEnv(constants.AdminNameEnvKey, applicationViperConfig.GetString(constants.AdminNameKey)),
		},
		IntDDName: models.IntDDName{
			DDName: getEnv(constants.IntDDNameEnvKey, (applicationViperConfig.GetString(constants.IntDDNameKey))),
		},
		BcastDDName: models.BcastDDName{
			DDName: getEnv(constants.BcastDDNameEnvKey, applicationViperConfig.GetString(constants.BcastDDNameKey)),
		},
		IntReqDDName: models.IntReqDDName{
			DDName: getEnv(constants.IntReqDDNameEnvKey, applicationViperConfig.GetString(constants.IntReqDDNameKey)),
		},
		RmsGetPrsntDDName: models.RmsGetPrsntDDName{
			DDName: getEnv(constants.RmsGetPrsntDDNameEnvKey, applicationViperConfig.GetString(constants.RmsGetPrsntDDNameKey)),
		},
		TouchlineDDName: models.TouchlineDDName{
			DDName: getEnv(constants.TouchlineDDNameEnvKey, applicationViperConfig.GetString(constants.TouchlineDDNameKey)),
		},
		RmsDDName: models.RmsDDName{
			DDName: getEnv(constants.RmsDDNameEnvKey, applicationViperConfig.GetString(constants.RmsDDNameKey)),
		},
	}
	return nil
}

func GetNestIniConfig() *models.NestIniConfig {
	return nestIniConfig
}
