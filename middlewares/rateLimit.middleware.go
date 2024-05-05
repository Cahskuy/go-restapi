package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimitHandler(limiter *ratelimit.Bucket) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Throttle requests using the rate limiter
		if limiter.TakeAvailable(1) < 1 {
			// If rate limit is exceeded, return HTTP 429 Too Many Requests status code
			ctx.JSON(http.StatusTooManyRequests, gin.H{"message": "Rate limit exceeded"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
