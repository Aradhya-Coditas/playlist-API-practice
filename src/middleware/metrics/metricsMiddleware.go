package metrics

import (
	"omnenest-backend/src/constants"
	"omnenest-backend/src/utils/metrics"
	"omnenest-backend/src/utils/tracer"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Metric records the basic API metrics like count, latency at each API level
func Metric() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, span := tracer.AddToSpan(ctx.Request.Context(), "Middleware-Metric")
		defer func() {
			if span != nil {
				span.End()
			}
		}()
		path := ctx.FullPath()
		method := ctx.Request.Method

		start := time.Now()

		ist, _ := time.LoadLocation(constants.IST)
		currentDateTimeInHour := time.Now().In(ist).Format(constants.PrometheusDateTimeFormat)
		ctx.Set(constants.APIRequestTime, currentDateTimeInHour)
		ctx.Next()
		serviceName := ctx.GetString(constants.ServiceNameLabel)
		statusCode := strconv.Itoa(ctx.Writer.Status())
		var nestOverallLatency float64
		if nestLatency, ok := ctx.Value(constants.NestOverallLatency).(float64); ok {
			nestOverallLatency = nestLatency
		}
		// increment counter
		metrics.Inc(ctx, constants.HttpRequestMetricLabel, serviceName, path, method, statusCode)
		// record latency histogram
		latency := time.Since(start).Seconds()
		metrics.Record(ctx, constants.HttpRequestMetricLabel, nestOverallLatency, path, constants.NestOverallLatencyLabel, serviceName)
		metrics.Record(ctx, constants.HttpRequestMetricLabel, latency-nestOverallLatency, path, constants.BFFLatencyLabel, serviceName)
		metrics.Record(ctx, constants.HttpRequestMetricLabel, latency, path, constants.APILatencyLabel, serviceName)
	}
}
