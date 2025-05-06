package headerCheck

import (
	"fmt"
	"net/http"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/responseUtils"
	"omnenest-backend/src/utils/tracer"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// HeaderCheck is the middleware to be used for validating x-request-id, DeviceId,
func HeaderCheck(serviceName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		parentSpan := trace.SpanFromContext(ctx.Request.Context())

		_, span := tracer.AddToSpan(ctx.Request.Context(), "Middleware-HeaderCheck")
		defer func() {
			if span != nil {
				span.End()
			}
		}()
		// add serviceName Constant to context
		ctx.Set(genericConstants.ServiceNameLabel, serviceName)
		log := logger.GetLoggerWithoutContext()
		requestId := ctx.GetHeader(genericConstants.RequestIDHeader)
		// check if requestId is nil or empty
		if requestId == "" {
			log.With(
				zap.String(genericConstants.RequestPath, ctx.Request.URL.Path),
				zap.String(genericConstants.RequestMethod, ctx.Request.Method),
				zap.String(genericConstants.ClientIPLogParam, ctx.ClientIP()),
				zap.String(genericConstants.RequestIDHeader, ctx.GetHeader(genericConstants.RequestIDHeader)),
				zap.Error(fmt.Errorf(genericConstants.XRequestIdKeyMissingError))).Error(genericConstants.XRequestIdKeyMissingError)
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.XRequestIdKeyMissingError))
			return
		}
		ctx.Set(genericConstants.RequestIDHeader, requestId)
		parentSpan.SetAttributes(attribute.String(genericConstants.RequestIDHeader, requestId))
		deviceId := ctx.GetHeader(genericConstants.DeviceIdHeader)
		// check if deviceId is nil or empty
		if deviceId == "" {
			log.With(
				zap.String(genericConstants.RequestPath, ctx.Request.URL.Path),
				zap.String(genericConstants.RequestMethod, ctx.Request.Method),
				zap.String(genericConstants.ClientIPLogParam, ctx.ClientIP()),
				zap.String(genericConstants.DeviceIdHeader, ctx.GetHeader(genericConstants.DeviceIdHeader)),
				zap.Error(fmt.Errorf(genericConstants.DeviceIdKeyMissingError))).Error(genericConstants.DeviceIdKeyMissingError)
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.DeviceIdKeyMissingError))
			return
		}
		ctx.Set(genericConstants.DeviceIdHeader, deviceId)
		parentSpan.SetAttributes(attribute.String(genericConstants.DeviceIdHeader, deviceId))
		appName := ctx.GetHeader(genericConstants.AppName)
		// check if appName is nil or empty
		if appName != "" {
			ctx.Set(genericConstants.AppName, appName)
		}

		buildNumber := ctx.GetHeader(genericConstants.BuildNumber)
		// check if buildNumber is nil or empty
		if buildNumber != "" {
			ctx.Set(genericConstants.BuildNumber, buildNumber)
		}

		packageName := ctx.GetHeader(genericConstants.PackageName)
		// check if PackageName is nil or empty
		if packageName != "" {
			ctx.Set(genericConstants.PackageName, packageName)
		}

		appVersion := ctx.GetHeader(genericConstants.AppVersion)
		// check if appVersion is nil or empty
		if appVersion == "" {
			log.With(
				zap.String(genericConstants.RequestPath, ctx.Request.URL.Path),
				zap.String(genericConstants.RequestMethod, ctx.Request.Method),
				zap.String(genericConstants.ClientIPLogParam, ctx.ClientIP()),
				zap.String(genericConstants.AppVersion, ctx.GetHeader(genericConstants.AppVersion)),
				zap.Error(fmt.Errorf(genericConstants.AppVersionKeyMissingError))).Error(genericConstants.AppVersionKeyMissingError)
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.AppVersionKeyMissingError))
			return
		}
		ctx.Set(genericConstants.AppVersion, appVersion)

		source := ctx.GetHeader(genericConstants.Source)
		// check if source is nil or empty
		if source == "" {
			log.With(
				zap.String(genericConstants.RequestPath, ctx.Request.URL.Path),
				zap.String(genericConstants.RequestMethod, ctx.Request.Method),
				zap.String(genericConstants.ClientIPLogParam, ctx.ClientIP()),
				zap.String(genericConstants.Source, ctx.GetHeader(genericConstants.Source)),
				zap.Error(fmt.Errorf(genericConstants.SourceKeyMissingError))).Error(genericConstants.SourceKeyMissingError)
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.SourceKeyMissingError))
			return

		} else if (strings.ToUpper(source) != genericConstants.MobSource) && (strings.ToUpper(source) != genericConstants.WebSource) {
			log.With(
				zap.String(genericConstants.RequestPath, ctx.Request.URL.Path),
				zap.String(genericConstants.RequestMethod, ctx.Request.Method),
				zap.String(genericConstants.ClientIPLogParam, ctx.ClientIP()),
				zap.String(genericConstants.Source, ctx.GetHeader(genericConstants.Source)),
				zap.Error(fmt.Errorf(genericConstants.InvalidSourceKeyError))).Error(genericConstants.InvalidSourceKeyError)
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.InvalidSourceKeyError))
			return
		}
		ctx.Set(genericConstants.Source, source)

		os := ctx.GetHeader(genericConstants.OS)
		// check if os is nil or empty
		if os != "" {
			ctx.Set(genericConstants.OS, os)
		}

		appInstallId := ctx.GetHeader(genericConstants.AppInstallId)
		// check if appInstallId is nil or empty
		if appInstallId == "" {
			log.With(
				zap.String(genericConstants.RequestPath, ctx.Request.URL.Path),
				zap.String(genericConstants.RequestMethod, ctx.Request.Method),
				zap.String(genericConstants.ClientIPLogParam, ctx.ClientIP()),
				zap.String(genericConstants.AppInstallId, ctx.GetHeader(genericConstants.AppInstallId)),
				zap.Error(fmt.Errorf(genericConstants.AppInstallIdKeyMissingError))).Error(genericConstants.AppInstallIdKeyMissingError)
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.AppInstallIdKeyMissingError))
			return
		}
		ctx.Set(genericConstants.AppInstallId, appInstallId)

		userAgent := ctx.GetHeader(genericConstants.UserAgent)
		// check if userAgent is nil or empty
		if userAgent == "" {
			log.With(
				zap.String(genericConstants.RequestPath, ctx.Request.URL.Path),
				zap.String(genericConstants.RequestMethod, ctx.Request.Method),
				zap.String(genericConstants.ClientIPLogParam, ctx.ClientIP()),
				zap.String(genericConstants.UserAgent, ctx.GetHeader(genericConstants.UserAgent)),
				zap.Error(fmt.Errorf(genericConstants.UserAgentKeyMissingError))).Error(genericConstants.UserAgentKeyMissingError)
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.UserAgentKeyMissingError))
			return
		}
		ctx.Set(genericConstants.UserAgent, userAgent)

		timeStamp := ctx.GetHeader(genericConstants.TimeStamp)
		// check if timeStamp is nil or empty
		if timeStamp == "" {
			log.With(
				zap.String(genericConstants.RequestPath, ctx.Request.URL.Path),
				zap.String(genericConstants.RequestMethod, ctx.Request.Method),
				zap.String(genericConstants.ClientIPLogParam, ctx.ClientIP()),
				zap.String(genericConstants.TimeStamp, ctx.GetHeader(genericConstants.TimeStamp)),
				zap.Error(fmt.Errorf(genericConstants.TimeStampKeyMissingError))).Error(genericConstants.TimeStampKeyMissingError)
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.TimeStampKeyMissingError))
			return
		}
		err := validateTimeStamp(timeStamp)
		if err != nil {
			log.With(
				zap.String(genericConstants.RequestPath, ctx.Request.URL.Path),
				zap.String(genericConstants.RequestMethod, ctx.Request.Method),
				zap.String(genericConstants.ClientIPLogParam, ctx.ClientIP()),
				zap.String(genericConstants.TimeStamp, ctx.GetHeader(genericConstants.TimeStamp)),
				zap.Error(fmt.Errorf(genericConstants.TimeStampInvalidFormatError))).Error(genericConstants.TimeStampInvalidFormatError)
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, err)
			return
		}
		ctx.Set(genericConstants.TimeStamp, timeStamp)
		parentSpan.SetAttributes(attribute.String(genericConstants.TimeStamp, timeStamp))

		bypass := ctx.GetHeader(genericConstants.Bypass)
		ctx.Set(genericConstants.Bypass, bypass)
	}
}

func validateTimeStamp(timeStamp string) error {
	// Convert the timestamp string to an integer (milliseconds since epoch)
	timeStampInt, err := strconv.ParseInt(timeStamp, 10, 64)
	if err != nil {
		return fmt.Errorf(genericConstants.TimeStampInvalidFormatError)
	}

	// Convert the epoch timestamp to time.Time
	parsedTime := time.Unix(0, timeStampInt*int64(time.Millisecond))

	// Compare the parsed timestamp with the current date
	currentDate := time.Now().Truncate(24 * time.Hour) // Truncate to midnight for date comparison
	if !parsedTime.Truncate(24 * time.Hour).Equal(currentDate) {
		return fmt.Errorf(genericConstants.TimeStampDateNotMatchError)
	}
	return nil
}
