package service

import (
	"fmt"

	"go-backend-challenge/internal/model"
	"go-backend-challenge/internal/repository"

	uar "github.com/Alejandrocuartas/core-models-private-library/models/user_agency_relations"
)

type UserAgencyRelationsServiceInterface interface {
	CreateUserAgencyRelationServiceMethod()
}

type UserAgencyRelationsServiceStruct struct {
	repository.AgenciesDbRepository
}

func (c UserAgencyRelationsServiceStruct) CreateUserAgencyRelationServiceMethod(
	requestData model.CreateUserAgencyRelationRequestModel,
) (
	model.CreateUserAgencyRelationResponseModel,
	error,
) {

	response := model.CreateUserAgencyRelationResponseModel{}
	UserAgencyRelationToCreate := uar.UserAgencyRelation{
		UserId:   requestData.UserId,
		AgencyId: requestData.AgencyId,
	}

	UserAgencyRelationToCreate, err := c.AgenciesDbRepository.CreateUserAgencyRelation(UserAgencyRelationToCreate)
	if err != nil {
		return response, fmt.Errorf("an error happened trying to create user agency relation: %s", err.Error())
	}

	response.AgencyId = UserAgencyRelationToCreate.AgencyId
	response.UserId = UserAgencyRelationToCreate.UserId

	return response, nil
}
