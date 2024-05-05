package main

import (
	"github.com/Cahskuy/go-restapi/initializers"
	"github.com/Cahskuy/go-restapi/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
