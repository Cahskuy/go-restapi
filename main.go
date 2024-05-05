package main

import (
	"net/http"

	"github.com/Cahskuy/go-restapi/initializers"
	"github.com/Cahskuy/go-restapi/middlewares"
	"github.com/Cahskuy/go-restapi/routes"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"github.com/rs/cors"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	// Create gin router
	app := gin.Default()

	// Create a rate limiter with a limit of 100 requests per second
	limiter := ratelimit.NewBucketWithRate(100, 100)

	// Specify allowed origins
	allowedOrigins := []string{"http://localhost:3000"}

	// Create CORS options with specified allowed origins
	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	})

	// Enforce rate limiting
	app.Use(middlewares.RateLimitHandler(limiter))

	// Use the CORS middleware
	app.Use(middlewares.CorsHandler(c))

	// Apply secure handler
	app.Use(middlewares.SecureHandler())

	// Apply nosurf used to prevent CSRF attacks
	// app.Use(nosurf.NewPure())

	app.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, world!"})
	})

	routes.InitRoute(app)

	app.Run()
}
