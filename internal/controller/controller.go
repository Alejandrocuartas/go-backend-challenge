package controller

import (
	"go-backend-challenge/internal/repository"
	"go-backend-challenge/internal/service"
)

type CustomControllerStruct struct {
	Campaigns           CampaignsControllerInterface
	Agencies            AgenciesControllerInterface
	Users               UsersControllerInterface
	UserAgencyRelations UserAgencyRelationsControllerInterface
	Actions             ActionsControllerInterface
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
		UserAgencyRelations: UserAgencyRelationsControllerStruct{
			UserAgencyRelationsServiceStruct: service.UserAgencyRelationsServiceStruct{
				AgenciesDbRepository: repository.AgenciesDbRepository{
					DB: agenciesDBConnection,
				},
			},
		},
		Actions: ActionsControllerStruct{
			ActionsServiceStruct: service.ActionsServiceStruct{
				AgenciesDbRepository: repository.AgenciesDbRepository{
					DB: agenciesDBConnection,
				},
			},
		},
	}
}
