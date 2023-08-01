package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	v1.GET("/oppening", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get oppening",
		})
	})

	v1.POST("/oppening", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Post oppening",
		})
	})

	v1.PUT("/oppening", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Put oppening",
		})
	})

	v1.DELETE("/oppening", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Delete oppening",
		})
	})

	v1.GET("/oppenings", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get List of oppening",
		})
	})
}
