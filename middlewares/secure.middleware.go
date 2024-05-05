package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// Function to create and configure secure middleware
func SecureHandler() gin.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny:          true, // Prevent framing of the application
		ContentTypeNosniff: true, // Prevent MIME type sniffing
		BrowserXssFilter:   true, // Enable XSS filter for the browser
	})

	return func(ctx *gin.Context) {
		err := secureMiddleware.Process(ctx.Writer, ctx.Request)

		// If there was an error, do not continue.
		if err != nil {
			ctx.Abort()
			return
		}

		// Avoid header rewrite if response is a redirection.
		if status := ctx.Writer.Status(); status > 300 && status < 399 {
			ctx.Abort()
		}
	}
}
