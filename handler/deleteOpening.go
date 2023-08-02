package handler

import (
	"fmt"
	"net/http"

	"github.com/CarlosGenuino/go-opportunities/schema"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		logger.Errorf("requestParameter id is required. The field is %s", id)
		sendError(ctx, http.StatusBadRequest, "requestParameter id is required")
		return
	}
	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("openning with id %s not found", id)
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("openning with id %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		logger.Errorf("erro deleting opening with id %s", id)
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening",opening)
}
