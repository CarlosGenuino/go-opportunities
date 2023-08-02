package handler

import (
	"net/http"

	"github.com/CarlosGenuino/go-opportunities/schema"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schema.Opening{}

	if err := db.Find(openings).Error; err != nil {
		logger.Error("error listing openings")
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
