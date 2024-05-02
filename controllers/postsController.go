package controllers

import (
	"fmt"

	"github.com/Cahskuy/go-crud/initializers"
	"github.com/Cahskuy/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(ctx *gin.Context) {
	fmt.Println(ctx.Request.Body)
	var req struct {
		Title string
		Body  string
	}

	ctx.Bind(&req)

	fmt.Println(ctx.Request.Body)

	post := models.Post{Title: req.Title, Body: req.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, gin.H{
		"message": post,
	})
}

func PostsIndex(ctx *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	ctx.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(ctx *gin.Context) {
	id := ctx.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	ctx.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	var req struct {
		Title string
		Body  string
	}

	ctx.Bind(&req)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: req.Title,
		Body:  req.Body,
	})

	ctx.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	ctx.JSON(200, gin.H{
		"message": "Delete successfully",
	})
}
