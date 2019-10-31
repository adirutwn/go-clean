package repositories

import (
	"github.com/pkg/errors"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/entities"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/modules/users/repositories/models"
)

func (r *repo) GetAll() ([]entities.User, error) {
	var users []models.User

	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to find user")
	}

	var entityUsers []entities.User
	for _, user := range users {

		entityUser, err := user.ToEntity()
		if err != nil {
			return nil, errors.Wrap(err, "failed to append user to entityUser")
		}

		entityUsers = append(entityUsers, *entityUser)
	}

	return entityUsers, nil
}
