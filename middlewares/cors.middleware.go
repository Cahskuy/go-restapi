package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func CorsHandler(ctx *cors.Cors) gin.HandlerFunc {
	// Return a Gin middleware function
	return func(ctx *gin.Context) {
		// Handle CORS for the request using the CORS middleware
		ctx.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		ctx.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Check for preflight request
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		// Continue processing the request
		ctx.Next()
	}
}
