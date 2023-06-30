package controller

import (
	utils "go-backend-challenge/core-utils-private-library"
	"go-backend-challenge/internal/model"
	"go-backend-challenge/internal/service"
	"net/http"
)

type AgenciesControllerStruct struct {
	service.AgenciesServiceStruct
}

type AgenciesControllerInterface interface {
	CreateAgencyControllerMethod(http.ResponseWriter, *http.Request)
}

func (c AgenciesControllerStruct) CreateAgencyControllerMethod(w http.ResponseWriter, r *http.Request) {

	var err error
	requestData := model.CreateAgencyRequestModel{}

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

	response, err := c.CreateAgencyServiceMethod(requestData)
	if err != nil {
		utils.ParseToJson(w, http.StatusInternalServerError, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	utils.ParseToJson(w, http.StatusOK, utils.Map{"status": "SUCCESS", "message": nil, "data": response})
}
