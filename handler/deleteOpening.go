package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pires-guilherme/go-jobs-opportunities/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {
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

	error = db.Delete(&opening).Error

	if error != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting id %s : %s", id, error))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)
}
