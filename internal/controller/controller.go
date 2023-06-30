package controller

import (
	"go-backend-challenge/internal/repository"
	"go-backend-challenge/internal/service"
)

type CustomControllerStruct struct {
	Campaigns CampaignsControllerInterface
	Agencies  AgenciesControllerInterface
	Users     UsersControllerInterface
}

func NewControl() CustomControllerStruct {
	agenciesDBConnection := repository.InitAgenciesDB()
	return CustomControllerStruct{
		Campaigns: CampaignsControllerStruct{
			CampaignsServiceStruct: service.CampaignsServiceStruct{
				AgenciesDbRepository: repository.AgenciesDbRepository{
					DB: agenciesDBConnection,
				},
			},
		},
		Agencies: AgenciesControllerStruct{
			AgenciesServiceStruct: service.AgenciesServiceStruct{
				AgenciesDbRepository: repository.AgenciesDbRepository{
					DB: agenciesDBConnection,
				},
			},
		},
		Users: UsersControllerStruct{
			UsersServiceStruct: service.UsersServiceStruct{
				AgenciesDbRepository: repository.AgenciesDbRepository{
					DB: agenciesDBConnection,
				},
			},
		},
	}
}
