package flags

import (
	genericConstants "omnenest-backend/src/constants"

	flag "github.com/spf13/pflag"
)

var (
	env            = flag.String(genericConstants.EnvironmentKey, genericConstants.EnvironmentDefaultValue, genericConstants.EnvironmentUsage)
	port           = flag.Int(genericConstants.PortKey, genericConstants.PortDefaultValue, genericConstants.PortUsage)
	baseConfigPath = flag.String(genericConstants.BaseConfigPathKey, genericConstants.BaseConfigPathDefaultValue, genericConstants.BaseConfigPathUsage)
	mockConfigPath = flag.String(genericConstants.MockConfigPathKey, genericConstants.MockConfigPathDefaultValue, genericConstants.MockConfigPathUsage)
	rootConfigPath = flag.String(genericConstants.RootConfigPathKey, genericConstants.RootConfigPathDefaultValue, genericConstants.RootConfigPathUsage)
)

func init() {
	flag.Parse()
}

// Env returns the application.yml runtime environment.
func Env() string {
	return *env
}

// Port returns the application.yml port number where the process will be started.
func Port() int {
	return *port
}

// BaseConfigPath returns the path that holds the configuration files.
func BaseConfigPath() string {
	return *baseConfigPath
}

// MockConfigPath returns the path that holds the mock configuration files.
func MockConfigPath() string {
	return *mockConfigPath
}

// RootConfigPath returns the path that holds the root configuration files.
func RootConfigPath() string {
	return *rootConfigPath
}
