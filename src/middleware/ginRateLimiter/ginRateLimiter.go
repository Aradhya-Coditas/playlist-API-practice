package ginratelimiter

import (
	"fmt"
	"net/http"
	"omnenest-backend/src/constants"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/responseUtils"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"go.uber.org/zap"
)

func RateLimitMiddleware(limiter *ratelimit.Bucket) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := logger.GetLoggerWithoutContext()
		log.Info(constants.AvailableRateLimitKey, zap.Int64(constants.AvailableRateLimitKey, limiter.Available()))
		// TODO: Check if limit is less than some threshold then setup alerting on the same
		if limiter.TakeAvailable(1) < 1 {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusTooManyRequests, fmt.Errorf(constants.RateLimitExceeded))
			return
		} else {
			ctx.Next()
		}
	}
}
