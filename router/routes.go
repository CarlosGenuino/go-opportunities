package router

import (
	"github.com/CarlosGenuino/go-opportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitilizeHandler()
	v1 := router.Group("/api/v1")

	v1.GET("/opening", handler.ReadOppeningHandler)

	v1.POST("/opening", handler.CreateOpeningHandler)

	v1.PUT("/opening", handler.UpdateOppeningHandler)

	v1.DELETE("/opening", handler.DeleteOpeningHandler)

	v1.GET("/openings", handler.ListOpeningsHandler)
}
