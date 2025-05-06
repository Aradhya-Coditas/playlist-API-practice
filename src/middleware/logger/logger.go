package logger

import (
	"time"

	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/utils/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logger is the middleware to be used for logging the request and response information
// This should be the first middleware to be utils added, in case the recovery middleware is not being used.
// Otherwise, it should be the second one, just after the recovery middleware.
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()
		log := logger.GetLogger(ctx)

		// Log the initial request
		log.With(
			zap.Time(genericConstants.StartTimeLogParam, start),
			zap.String(genericConstants.RequestPath, ctx.Request.URL.Path),
			zap.String(genericConstants.RequestMethod, ctx.Request.Method),
		).Info(genericConstants.StartTransaction)
		// Process request
		ctx.Next()

		zapRequestId := zap.String(genericConstants.RequestIDHeader, "")
		if ctx.Value(genericConstants.RequestIDHeader) != nil {
			zapRequestId = zap.String(genericConstants.RequestIDHeader, ctx.Value(genericConstants.RequestIDHeader).(string))
		}
		log.Info(genericConstants.EndTransaction,
			zap.Int(genericConstants.StatusCodeKey, ctx.Writer.Status()),
			zap.Any(genericConstants.RequestBody, ctx.Value(genericConstants.RequestBody)),
			zap.Any(genericConstants.ResponseBody, ctx.Value(genericConstants.ResponseBody)),
			zapRequestId,
			zap.String(genericConstants.RequestPath, ctx.Request.URL.Path),
			zap.String(genericConstants.RequestMethod, ctx.Request.Method),
			zap.Int64(genericConstants.LatencyKey, time.Since(start).Milliseconds()),
			zap.String(genericConstants.ClientIPLogParam, ctx.ClientIP()),
			zap.String(genericConstants.ErrorLogParam, ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
	}
}
