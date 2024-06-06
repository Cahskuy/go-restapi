package main

import (
	"net/http"

	"github.com/Cahskuy/go-restapi/initializers"
	"github.com/Cahskuy/go-restapi/middlewares"
	"github.com/Cahskuy/go-restapi/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {

	app := gin.Default()

	app.Use(middlewares.RateLimitHandler())

	app.Use(middlewares.CorsHandler())

	app.Use(middlewares.SecureHandler())

	app.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, world!"})
	})

	routes.SetupRoutes(app)

	app.Run()
}
