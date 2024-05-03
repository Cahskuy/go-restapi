package routes

import (
	"github.com/Cahskuy/go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app
	route.POST("/posts", controllers.PostsCreate)
	route.PUT("/posts/:id", controllers.PostsUpdate)
	route.GET("/posts", controllers.PostsIndex)
	route.GET("/posts/:id", controllers.PostsShow)
	route.DELETE("/posts/:id", controllers.PostsDelete)
}
