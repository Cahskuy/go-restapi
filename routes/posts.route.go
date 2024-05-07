package routes

import (
	"github.com/Cahskuy/go-restapi/controllers"
	"github.com/Cahskuy/go-restapi/middlewares"
	"github.com/Cahskuy/go-restapi/schemas"
	"github.com/gin-gonic/gin"
)

func Posts(g *gin.RouterGroup) {
	// Initialize custom validator
	validator := middlewares.NewValidator()

	g.POST("/", middlewares.ValidationHandler(validator, schemas.Post{}), controllers.PostsCreate)
	g.PUT("/:id", controllers.PostsUpdate)
	g.GET("/", controllers.PostsIndex)
	g.GET("/:id", controllers.PostsShow)
	g.DELETE("/:id", controllers.PostsDelete)
}
