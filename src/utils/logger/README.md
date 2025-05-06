# Logger Service

The logger package provides a flexible and efficient logging mechanism for the application. It is built using Zap which is a logging library developed by Uber.

Below is the short description for the logger function that are being used.
* `Info(string, ...zap.Field)`: Logs an informational message with optional fields.
* `Warn(string, ...zap.Field)`: Logs a warning message with optional fields.
* `Error(string, ...zap.Field)`: Logs an error message with optional fields.
* `Debug(string, ...zap.Field)`: Logs a debug message with optional fields.
* `Fatal(string, ...zap.Field)`: Logs a fatal error message with optional fields and exits the application.
* `Panic(string, ...zap.Field)`: Logs a panic message with optional fields and panics.
* `With(args ...zap.Field) Logger`: Returns a new logger with additional context fields.
* `Sync()`: Flushes any buffered log entries.

## Logger Functions

### GetLogger

The `GetLogger` function returns a logger with a request ID field if available in the context; otherwise, returns the default logger.

### GetLoggerWithoutContext

The `GetLoggerWithoutContext` function returns the default logger instance without any additional context. It simply returns the global log variable, which holds the logger instance.

### LogLatency

The `LogLatency` function that logs the latency (time elapsed) since a given start time using the provided logger. The logged message contains the latency value in milliseconds along with a custom message. `once sync.Once` is a synchronization mechanism ensuring loc initialization happens only once.

### LogTimeEncoder

`LogTimeEncoder` function formats a given time value using a specific time zone and time format, and appends it to an encoder.

### LevelEncoder

The `LevelEncoder`takes a log level and an encoder as input parameters. It appends the capital string representation of the log level to the encoder.

## Setup Function

### StartLogger

`StartLogger` function initializes the logging setup for the application by calling SetupLogging with the specified log level.

## ZapLogger Struct
    This struct defines a struct type that contains a single field which is a pointer to a zap.Logger instance.

## Zap Functions

### SetupLogging

The `SetupLogging` function initializes the logging system based on the provided log level using the zap package. It defaults to InfoLevel if the specified level is not found. Configuration includes settings for message, level, time, and function keys. Finally, it creates a new logger with caller information added.

## Usage

```go
    package main

    import (
        "context"
        "time"

        "github.com/your-username/your-repo/logger"
    )

    func main() {
        // Setup logging with a specified log level
        logger.SetupLogging("debug")

        // Get logger instance
        log := logger.GetLoggerWithoutContext()

        // Log some messages
        log.Info("This is an informational message.")
        log.Warn("This is a warning message.")
        log.Debug("This is a debug message.")
        log.Error("This is an error message.")

        // Log latency
        start := time.Now()
        // Perform some operation
        time.Sleep(100 * time.Millisecond)
        elapsed := time.Since(start)
        logger.LogLatency(log, start, "Operation completed in")

        // Flush log entries
        log.Sync()
    }
```

## Reference
For more information on the `zap` package, you can refer to the documentation on [zap](https://pkg.go.dev/go.uber.org/zap).

