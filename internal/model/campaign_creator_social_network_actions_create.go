package model

import (
	"errors"
)

type CreateActionRequestModel struct {
	CodeName               string `json:"code_name"`
	Quantity               int    `json:"quantity"`
	CampaignId             uint   `json:"campaign_id"`
	CreatorId              uint   `json:"creator_id"`
	CreatorSocialNetworkId uint   `json:"creator_social_network_id"`
}

func (data CreateActionRequestModel) ValidateData() error {

	if data.CodeName == "" {
		return errors.New("code_name is required")
	} else if data.Quantity == 0 {
		return errors.New("quantity is required")
	} else if data.CampaignId == 0 {
		return errors.New("campaign_id is required")
	} else if data.CreatorId == 0 {
		return errors.New("creator_id is required")
	} else if data.CreatorSocialNetworkId == 0 {
		return errors.New("creator_social_network_id is required")
	}

	return nil
}

type CreateActionResponseModel struct {
	CodeName               string `json:"code_name"`
	Quantity               int    `json:"quantity"`
	CampaignId             uint   `json:"campaign_id"`
	CreatorId              uint   `json:"creator_id"`
	CreatorSocialNetworkId uint   `json:"creator_social_network_id"`
}
