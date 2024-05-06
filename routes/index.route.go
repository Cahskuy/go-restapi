package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	v1 := route.Group("/v1")

	Posts(v1.Group("/posts"))
}
