package users

import "gitlab.com/lightnet-thailand/poc/clean-workshop/app/entities"

type UseCase interface {
	GetAll() ([]entities.User, error)
}
