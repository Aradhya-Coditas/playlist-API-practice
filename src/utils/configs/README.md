# Configs Package

This package provides all the necessary configuration pacakages for the application, machine, and the APIs.

## ApplicationConfig

This package provides functions to initialize and retrieve application configurations, allowing easy access to various settings such as server port, token secrets, file paths, and environment-specific overrides, enhancing the flexibility and maintainability of the application's configuration management.

### Initialization

The package initializes the application configurations. 

#### Steps
1. The function retrieves the application configuration from a YAML file using the Get function. 
2. The retrieved configuration is then used to populate the applicationConfig variable. The applicationConfig variable is of type *models.ApplicationConfig and contains various sub-configurations such as SwaggerConfig, Server, AppConfig, Token, and FilePath.
4. Each sub-configuration is populated with values from the applicationViperConfig object. 
5. The values are retrieved using constants defined in the constants package. Some values are also retrieved from environment variables using the getEnvInt and getEnv functions.
6. Finally, the function returns nil if successful, otherwise it returns an error.

### Functions

#### GetApplicationConfig

The `GetApplicationConfig()` function returns the current application configuration. 

#### getEnvInt

The `getEnvInt()` function retrieves the value of an environment variable as an integer. If the environment variable is not set or cannot be parsed as an integer, it returns the provided default value.

#### getEnv

The `getEnv()` function retrieves the value of an environment variable based on the provided key.

#### Usage

```go
func main() {
	err := configs.InitApplicationConfigs(context.Background())
	if err != nil {
		fmt.Println("Error initializing application configurations:", err)
		return
	}

	applicationConfig := configs.GetApplicationConfig()

	fmt.Println("Swagger Host:", applicationConfig.SwaggerConfig.SwaggerHost)
}

```


## Configs

This package provides functions for initializing, accessing, and caching configurations using viper.

### Initialization

The package initializes the configurations by setting the base config paths and creating a new instance of the providers struct.

### Functions

#### Get

The `Get()` function is used to get the instance to the config provider for the configuration name.

#### Usage

```go
func main() {
	configs.Init([]string{"./configs"})

	configName := "example_config"
	configProvider, err := configs.Get(configName)
	if err != nil {
		fmt.Printf("Error getting configuration %s: %v\n", configName, err)
		return
	}

	value := configProvider.GetString("key")
	fmt.Printf("Value from config: %s\n", value)
}
```


## HostNameConfig

This package provides functionality to set and retrieve the hostname of the machine.

### Functions

#### SetHostName

The `SetHostName()` function is used to set the host name of the machine.

#### GetHostName

The `GetHostName()` function returns the host name of the machine.

#### Usage

```go
func main() {
	service := "example"
	configs.SetHostName(service)

	hostName := configs.GetHostName()
	fmt.Println("Host Name:", hostName)
}
```


## NestAPIConfig

The package provides a centralized configuration mechanism to retrieve the API URL for a given endpoint.

### Functions

#### GetNestAPIUrl

The `GetNestAPIUrl()` function returns the API URL for the given endpoint.

#### Usage

```go
func main() {
	endpoint := "endpoint1"
	apiURL := GetNestAPIUrl(endpoint)
	fmt.Println("API URL for", endpoint, ":", apiURL)
}
```