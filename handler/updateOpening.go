package handler

import (
	"fmt"
	"net/http"

	"github.com/CarlosGenuino/go-opportunities/schema"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	update := UpdateOpeningRequest{}
	ctx.BindJSON(&update)

	if err := update.Validate(); err != nil {
		logger.Errorf("validation error: %s", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		logger.Error("queryParameter id is required")
		sendError(ctx, http.StatusBadRequest, "queryParameter id is required")
		return
	}

	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Not found opening for id %s", id))
		return
	}

	if update.Company != "" {
		opening.Company = update.Company
	}

	if update.Link != "" {
		opening.Link = update.Link
	}

	if update.Location != "" {
		opening.Location = update.Location
	}

	if update.Role != "" {
		opening.Role = update.Role
	}

	if update.Remote != nil {
		opening.Remote = *update.Remote
	}

	if update.Salary > 0 {
		opening.Salary = update.Salary
	}

	if err := opening.Validate(); err != nil {
		logger.Errorf("validation entity error: %s", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Error updating opening: %s", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
	}

	sendSuccess(ctx, "update-opening", opening)
}
