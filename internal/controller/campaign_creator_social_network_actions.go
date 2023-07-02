package controller

import (
	"go-backend-challenge/internal/model"
	"go-backend-challenge/internal/service"
	"net/http"

	utils "github.com/Alejandrocuartas/core-utils-private-library"
)

type ActionsControllerStruct struct {
	service.ActionsServiceStruct
}

type ActionsControllerInterface interface {
	CreateActionControllerMethod(http.ResponseWriter, *http.Request)
}

func (c ActionsControllerStruct) CreateActionControllerMethod(w http.ResponseWriter, r *http.Request) {

	var err error
	requestData := model.CreateActionRequestModel{}

	err = utils.DecodeData(r, &requestData)
	if err != nil {
		utils.ParseToJson(w, http.StatusBadRequest, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	err = requestData.ValidateData()
	if err != nil {
		utils.ParseToJson(w, http.StatusBadRequest, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	response, err := c.CreateActionServiceMethod(requestData)
	if err != nil {
		utils.ParseToJson(w, http.StatusInternalServerError, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	utils.ParseToJson(w, http.StatusOK, utils.Map{"status": "SUCCESS", "message": nil, "data": response})
}
