package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(ctx *gin.Context, code int, msg string) {
	var status string
	if code == http.StatusInternalServerError {
		status = "error"
	} else {
		status = "fail"
	}

	ctx.AbortWithStatusJSON(code, gin.H{
		"status":  status,
		"message": msg,
	})
}
