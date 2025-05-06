package constants

// Environment Flag Constants
const (
	EnvironmentKey          = "env"
	EnvironmentDefaultValue = ""
	EnvironmentUsage        = "application.yml runtime environment"
)

// Config Flag Constants
const (
	BaseConfigPathKey                  = "base-config-path"
	BaseConfigPathDefaultValue         = "../../configs"
	BaseConfigPathUsage                = "path to folder that stores your configurations"
	MockConfigPathKey                  = "mock-config-path"
	MockConfigPathDefaultValue         = "../../utils/mockResources"
	MockConfigPathUsage                = "path to folder that stores mock resources"
	RootConfigPathKey                  = "root-config-path"
	RootConfigPathDefaultValue         = "src/configs"
	RootConfigPathUsage                = "path to folder that stores root resources"
	EncryptionKeyDestinationPath       = "../../encryptionKeys/"
	EncryptionKeyMiddlewareDestination = "../../middlewareEncryptionKeys/"
	StaticExchangeConfigFilePath       = "resources/static_exchange_config.csv"
)

// Port Flag Constants
const (
	PortKey          = "port"
	PortDefaultValue = 8080
	PortUsage        = "application.yml port"
)
