package model

import (
	"errors"
)

type CreateAgencyRequestModel struct {
	Name string `json:"name"`
}

func (data CreateAgencyRequestModel) ValidateData() error {

	if data.Name == "" {
		return errors.New("agency_name is required")
	}

	return nil
}

type CreateAgencyResponseModel struct {
	AgencyId   uint   `json:"agency_id"`
	AgencyName string `json:"agency_name"`
}
