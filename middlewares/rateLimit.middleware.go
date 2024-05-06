package middlewares

import (
	"net/http"
	"time"

	"github.com/Cahskuy/go-restapi/utils"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimitHandler() gin.HandlerFunc {
	// Create a rate limiter with a limit of 100 requests per second
	limiter := ratelimit.NewBucket(time.Duration(time.Second), 100)

	return func(ctx *gin.Context) {
		// Throttle requests using the rate limiter
		if limiter.TakeAvailable(1) == 0 {
			// If rate limit is exceeded, return HTTP 429 Too Many Requests status code
			utils.ErrorResponse(ctx, http.StatusTooManyRequests, "Rate limit exceeded")
			return
		}
		ctx.Next()
	}
}
