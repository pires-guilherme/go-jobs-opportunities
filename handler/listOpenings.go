package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pires-guilherme/go-jobs-opportunities/schemas"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	err := db.Find(&openings).Error

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error listing openings %s", err))
		return
	}

	sendSuccess(ctx, "list-openings", openings)

}
