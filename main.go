package main

import (
	"github.com/Cahskuy/go-crud/initializers"
	"github.com/Cahskuy/go-crud/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	app := gin.Default()

	routes.InitRoute(app)

	app.Run()
}
