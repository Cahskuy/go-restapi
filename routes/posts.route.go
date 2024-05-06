package routes

import (
	"github.com/Cahskuy/go-restapi/controllers"
	"github.com/gin-gonic/gin"
)

func Posts(g *gin.RouterGroup) {
	g.POST("/", controllers.PostsCreate)
	g.PUT("/:id", controllers.PostsUpdate)
	g.GET("/", controllers.PostsIndex)
	g.GET("/:id", controllers.PostsShow)
	g.DELETE("/:id", controllers.PostsDelete)
}
