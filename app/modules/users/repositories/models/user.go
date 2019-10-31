package models

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/entities"
)

type User struct {
	FirstName string `gorm:"first_name"`
	LastName string `gorm:"last_name"`
}

func (u *User) ToEntity() (*entities.User, error) {
	var entityUser entities.User
	err := copier.Copy(&entityUser, u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse from model to entity")
	}

	return &entityUser, nil
}
