package service

import (
	"fmt"

	"go-backend-challenge/internal/model"
	"go-backend-challenge/internal/repository"

	"github.com/Alejandrocuartas/core-models-private-library/models/users"
)

type UsersServiceInterface interface {
	CreateUserServiceMethod()
}

type UsersServiceStruct struct {
	repository.AgenciesDbRepository
}

func (c UsersServiceStruct) CreateUserServiceMethod(
	requestData model.CreateUserRequestModel,
) (
	model.CreateUserResponseModel,
	error,
) {

	response := model.CreateUserResponseModel{}
	userToCreate := users.User{
		FirstName:           requestData.FirstName,
		LastName:            requestData.LastName,
		Email:               requestData.Email,
		CognitoId:           requestData.Email + "/" + requestData.FirstName,
		OnBoardingCompleted: false,
	}

	userToCreate, err := c.AgenciesDbRepository.CreateUser(userToCreate)
	if err != nil {
		return response, fmt.Errorf("an error happened trying to create user: %s", err.Error())
	}

	response.Email = userToCreate.Email
	response.LastName = userToCreate.LastName
	response.FirstName = userToCreate.FirstName

	return response, nil
}
