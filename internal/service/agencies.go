package service

import (
	"fmt"

	"go-backend-challenge/core-models-private-library/models/agencies"
	"go-backend-challenge/internal/model"
	"go-backend-challenge/internal/repository"
)

type AgenciesServiceInterface interface {
	CreateAgencyServiceMethod()
}

type AgenciesServiceStruct struct {
	repository.AgenciesDbRepository
}

func (c AgenciesServiceStruct) CreateAgencyServiceMethod(
	requestData model.CreateAgencyRequestModel,
) (
	model.CreateAgencyResponseModel,
	error,
) {

	response := model.CreateAgencyResponseModel{}
	agencyToCreate := agencies.Agency{
		Name: requestData.Name,
	}

	agencyToCreate, err := c.AgenciesDbRepository.CreateAgency(agencyToCreate)
	if err != nil {
		return response, fmt.Errorf("an error happened trying to create agency: %s", err.Error())
	}

	response.AgencyId = agencyToCreate.ID
	response.AgencyName = agencyToCreate.Name

	return response, nil
}
