package controller

import (
	"go-backend-challenge/internal/model"
	"go-backend-challenge/internal/service"
	"net/http"

	utils "github.com/Alejandrocuartas/core-utils-private-library"
)

type UsersControllerStruct struct {
	service.UsersServiceStruct
}

type UsersControllerInterface interface {
	CreateUserControllerMethod(http.ResponseWriter, *http.Request)
}

func (c UsersControllerStruct) CreateUserControllerMethod(w http.ResponseWriter, r *http.Request) {

	var err error
	requestData := model.CreateUserRequestModel{}

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

	response, err := c.CreateUserServiceMethod(requestData)
	if err != nil {
		utils.ParseToJson(w, http.StatusInternalServerError, utils.Map{"status": "ERROR", "message": err.Error()})
		return
	}

	utils.ParseToJson(w, http.StatusOK, utils.Map{"status": "SUCCESS", "message": nil, "data": response})
}
