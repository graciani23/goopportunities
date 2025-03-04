package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graciani23/goopportunities/schemas"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List a new job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "listOpenings", openings)
}
