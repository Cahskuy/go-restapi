package main

import (
	"github.com/Cahskuy/go-crud/initializers"
	"github.com/Cahskuy/go-crud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main()  {
	initializers.DB.AutoMigrate(&models.Post{})
}