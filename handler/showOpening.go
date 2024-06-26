package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pires-guilherme/go-jobs-opportunities/schemas"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	error := db.First(&opening, id).Error

	if error != nil {
		sendError(ctx, http.StatusNotFound, "Id not found")
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
