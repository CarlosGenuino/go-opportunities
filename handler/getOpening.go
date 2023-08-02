package handler

import (
	"fmt"
	"net/http"

	"github.com/CarlosGenuino/go-opportunities/schema"
	"github.com/gin-gonic/gin"
)

func ReadOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, fmt.Sprintf("requestParameter id is required"))
	}
	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Not found opening for id %s", id))
	}
	
	sendSuccess(ctx, "get-opening", opening)
}