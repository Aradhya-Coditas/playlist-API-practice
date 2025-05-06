package configs

import (
	"fmt"
	"sync"

	genericConstants "omnenest-backend/src/constants"

	"github.com/spf13/viper"
)

type providers struct {
	providers map[string]*viper.Viper
	mu        sync.Mutex
}

var baseConfigPaths []string
var p *providers

// Init initializes the configurations.
func Init(paths []string) {
	baseConfigPaths = paths
	p = &providers{
		providers: make(map[string]*viper.Viper),
	}
}

// Get is used to get the instance to the config provider for the configuration name
func Get(name string) (*viper.Viper, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// see for an existing provider
	if provider, ok := p.providers[name]; ok {
		// provider already exists
		// use it
		return provider, nil
	}

	// try to get the provider
	provider := viper.New()
	provider.SetConfigName(name)
	for _, baseConfigPath := range baseConfigPaths {
		provider.AddConfigPath(baseConfigPath)
	}

	err := provider.ReadInConfig()
	if err != nil {
		// config not found or some other parsing errors
		return nil, fmt.Errorf(genericConstants.ConfigParsingError, name, err)
	}

	// added on file change listener uncomment only if needed
	// provider.OnConfigChange(func(e fsnotify.Event) {
	// 	p.providers[name] = provider
	// })

	// add a watcher for this provider uncomment only if needed
	// provider.WatchConfig()

	// successfully found config, store it for future use
	p.providers[name] = provider

	return provider, nil
}
