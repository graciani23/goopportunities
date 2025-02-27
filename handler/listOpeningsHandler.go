package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graciani23/goopportunities/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "listOpenings", openings)
}
