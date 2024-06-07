package controllers

import (
	"net/http"

	"github.com/Cahskuy/go-restapi/initializers"
	"github.com/Cahskuy/go-restapi/models"
	"github.com/Cahskuy/go-restapi/schemas"
	"github.com/Cahskuy/go-restapi/utils"
	"github.com/gin-gonic/gin"
)

func PostsCreate(ctx *gin.Context) {
	post := ctx.MustGet("payload").(*schemas.Post)

	println("INI TITLE", post.Title)
	println("INI BODY", *(*post).Body)

	// err := initializers.DB.Create(&post).Error

	// if err != nil {
	// 	utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": post,
	})
}

func PostsIndex(ctx *gin.Context) {
	var posts []models.Post
	err := initializers.DB.Find(&posts).Error
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": posts,
	})

}

func PostsShow(ctx *gin.Context) {
	id := ctx.Param("id")
	var post models.Post
	err := initializers.DB.First(&post, id).Error

	var httpErrorCode int
	if err != nil {
		switch err.Error() {
		case "record not found":
			httpErrorCode = http.StatusNotFound
		default:
			httpErrorCode = http.StatusInternalServerError
		}

		utils.ErrorResponse(ctx, httpErrorCode, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": post,
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
	err1 := initializers.DB.First(&post, id).Error

	var httpErrorCode1 int
	if err1 != nil {
		switch err1.Error() {
		case "record not found":
			httpErrorCode1 = http.StatusNotFound
		default:
			httpErrorCode1 = http.StatusInternalServerError
		}

		utils.ErrorResponse(ctx, httpErrorCode1, err1.Error())
		return
	}

	err2 := initializers.DB.Model(&post).Updates(models.Post{
		Title: req.Title,
		Body:  req.Body,
	}).Error

	if err2 != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err2.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"posts": post,
	})
}

func PostsDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := initializers.DB.Delete(&models.Post{}, id).Error
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete successfully",
	})
}
