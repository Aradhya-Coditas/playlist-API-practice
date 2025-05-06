# Flags Package

This package provides setting all the flags

## Initialization

The package defines and initializes several variables for storing flag values, such as `env`, `port`, `baseConfigPath`, `mockConfigPath`, and `rootConfigPath`. The `init` function is called to parse the flags.

## Functions

### Env

The `Env()` function returns the application.yml runtime environment.

### Port

The `Port()` function returns the application.yml port number where the process will be started.

### BaseConfigPath

The `BaseConfigPath()` function returns the path that holds the configuration files.

### MockConfigPath

The `MockConfigPath()` function returns the path that holds the mock configuration files.

### RootConfigPath

The `RootConfigPath()` function returns the path that holds the root configuration files.

## Usage

To incorporate the flags service into your application, adhere to the following steps:

1. Import the `flags` package into your application.
2. Utilize the `init()` function to acquire the singleton instance of the flags service.
3. Invoke the necessary functions to fetch the respective paths as per your requirements.

Example:

```go
package main

import (
	"context"
	"fmt"
)

type flags struct{}

func (f flags) RootConfigPath() string {
	return "/path/to/root/config"
}

type configs struct{}

func (c configs) Init(paths []string) {
	fmt.Println("Initializing configurations with paths:")
	for _, path := range paths {
		fmt.Println(path)
	}
}

func main() {
	flags := flags{}

	rootConfigPath := flags.RootConfigPath()

	configs := configs{}
	configs.Init([]string{rootConfigPath})
}
```