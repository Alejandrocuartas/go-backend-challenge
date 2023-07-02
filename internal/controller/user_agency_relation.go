package controller

import (
	"go-backend-challenge/internal/model"
	"go-backend-challenge/internal/service"
	"net/http"

	utils "github.com/Alejandrocuartas/core-utils-private-library"
)

type UserAgencyRelationsControllerStruct struct {
	service.UserAgencyRelationsServiceStruct
}

type UserAgencyRelationsControllerInterface interface {
	CreateUserAgencyRelationsControllerMethod(http.ResponseWriter, *http.Request)
}

func (c UserAgencyRelationsControllerStruct) CreateUserAgencyRelationsControllerMethod(w http.ResponseWriter, r *http.Request) {

	var err error
	requestData := model.CreateUserAgencyRelationRequestModel{}

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

	response, err := c.CreateUserAgencyRelationServiceMethod(requestData)
	if err != nil {
		utils.ParseToJson(w, http.StatusInternalServerError, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	utils.ParseToJson(w, http.StatusOK, utils.Map{"status": "SUCCESS", "message": nil, "data": response})
}
