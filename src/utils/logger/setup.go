package logger

import (
	"context"
)

// StartLogger is a function that initializes the logging setup for the application.
func StartLogger(ctx context.Context, logLevel string) {
	SetupLogging(logLevel)
}
