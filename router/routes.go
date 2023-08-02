package router

import (
	"github.com/CarlosGenuino/go-opportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitilizeHandler()
	v1 := router.Group("/api/v1")

	v1.GET("/oppening", handler.ReadOppeningHandler)

	v1.POST("/oppening", handler.CreateOppeningHandler)

	v1.PUT("/oppening", handler.UpdateOppeningHandler)

	v1.DELETE("/oppening", handler.DeleteOppeningHandler)

	v1.GET("/oppenings", handler.ListOppeningsHandler)
}
