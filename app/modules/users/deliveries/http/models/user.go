package models

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/entities"
)

type UserResponse struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

func (u *UserResponse) ParseEntity(userEntity entities.User) (*UserResponse, error) {
	err := copier.Copy(u, &userEntity)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse entity to httpModel")
	}

	return u, nil
}
