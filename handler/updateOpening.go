package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pires-guilherme/go-jobs-opportunities/schemas"
)

func UpdateOpeningHandler(ctx *gin.Context) {
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

	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	err := db.Updates(&opening).Error

	if err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
	}

	sendSuccess(ctx, "update-opening", opening)
}
