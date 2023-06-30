package model

import (
	"errors"
)

type CreateUserRequestModel struct {
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	Email               string `json:"email"`
	CognitoId           string
	OnBoardingCompleted bool
}

func (data CreateUserRequestModel) ValidateData() error {

	if data.Email == "" {
		return errors.New("email is required")
	}

	return nil
}

type CreateUserResponseModel struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
