package model

import (
	"errors"
)

type CreateUserAgencyRelationRequestModel struct {
	UserId   uint `json:"user_id"`
	AgencyId uint `json:"agency_id"`
}

func (data CreateUserAgencyRelationRequestModel) ValidateData() error {

	if data.UserId == 0 {
		return errors.New("user_id is required")
	} else if data.AgencyId == 0 {
		return errors.New("agency_id is required")
	}

	return nil
}

type CreateUserAgencyRelationResponseModel struct {
	UserId   uint `json:"user_id"`
	AgencyId uint `json:"agency_id"`
}
