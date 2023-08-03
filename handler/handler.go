package handler

import (
	"github.com/CarlosGenuino/go-opportunities/config"
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
