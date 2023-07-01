package service

import (
	"fmt"

	m "go-backend-challenge/core-models-private-library/models/campaign_creator_social_network_actions"
	"go-backend-challenge/internal/model"
	"go-backend-challenge/internal/repository"
)

type ActionsServiceInterface interface {
	CreateActionServiceMethod()
}

type ActionsServiceStruct struct {
	repository.AgenciesDbRepository
}

func (c ActionsServiceStruct) CreateActionServiceMethod(
	requestData model.CreateActionRequestModel,
) (
	model.CreateActionResponseModel,
	error,
) {

	response := model.CreateActionResponseModel{}
	actionToCreate := m.CampaignCreatorSocialNetworkActions{
		CodeName:               requestData.CodeName,
		Quantity:               requestData.Quantity,
		CampaignId:             requestData.CampaignId,
		CreatorId:              requestData.CreatorId,
		CreatorSocialNetworkId: requestData.CreatorSocialNetworkId,
	}

	actionToCreate, err := c.AgenciesDbRepository.CreateAction(actionToCreate)
	if err != nil {
		return response, fmt.Errorf("an error happened trying to create an action: %s", err.Error())
	}

	response.CodeName = actionToCreate.CodeName
	response.Quantity = actionToCreate.Quantity
	response.CampaignId = actionToCreate.CampaignId
	response.CreatorId = actionToCreate.CreatorId
	response.CreatorSocialNetworkId = actionToCreate.CreatorSocialNetworkId

	return response, nil
}
