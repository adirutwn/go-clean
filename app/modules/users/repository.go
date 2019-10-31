package users

import "gitlab.com/lightnet-thailand/poc/clean-workshop/app/entities"

type Repository interface {
	GetAll() ([]entities.User, error)
}
