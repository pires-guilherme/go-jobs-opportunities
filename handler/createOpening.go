package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pires-guilherme/go-jobs-opportunities/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	err := request.Validate()

	if err != nil {
		logger.Errorf("validation error : %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Link:     request.Link,
		Remote:   *request.Remote,
		Salary:   request.Salary,
	}

	err = db.Create(&opening).Error

	if err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
	}

	sendSuccess(ctx, "create-opening", opening)

	//logger.Infof("request received %+v", request)
}
