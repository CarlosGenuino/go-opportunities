package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create Oppening",
	})
}

func ReadOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Read Oppening",
	})
}

func UpdateOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update Oppening",
	})
}

func DeleteOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete Oppening",
	})
}

func ListOppeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List of Oppening",
	})
}
