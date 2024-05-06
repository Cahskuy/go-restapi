package routes

import (
	"github.com/Cahskuy/go-restapi/controllers"
	"github.com/Cahskuy/go-restapi/middlewares"
	"github.com/Cahskuy/go-restapi/schemas"
	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	// Initialize custom validator
	validator := middlewares.NewCustomValidator()

	// Define routes
	route.POST("/posts", middlewares.ValidationHandler(validator, schemas.Post{}), controllers.PostsCreate)
	route.PUT("/posts/:id", middlewares.ValidationHandler(validator, schemas.Post{}), controllers.PostsUpdate)
	route.GET("/posts", controllers.PostsIndex)
	route.GET("/posts/:id", controllers.PostsShow)
	route.DELETE("/posts/:id", controllers.PostsDelete)
}
