package handler

import (
	"net/http"

	"github.com/CarlosGenuino/go-opportunities/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitilizeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
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
