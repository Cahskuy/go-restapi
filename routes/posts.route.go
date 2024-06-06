package routes

import (
	"github.com/Cahskuy/go-restapi/controllers"
	"github.com/Cahskuy/go-restapi/middlewares"
	"github.com/Cahskuy/go-restapi/schemas"
	"github.com/gin-gonic/gin"
)

func PostsRoutes(router *gin.RouterGroup) {
	postsRoutes := router.Group("/posts")
	{
		postsRoutes.POST("/", middlewares.ValidationHandler(schemas.Post{}), controllers.PostsCreate)
		postsRoutes.PUT("/:id", controllers.PostsUpdate)
		postsRoutes.GET("/", controllers.PostsIndex)
		postsRoutes.GET("/:id", controllers.PostsShow)
		postsRoutes.DELETE("/:id", controllers.PostsDelete)
	}

}
