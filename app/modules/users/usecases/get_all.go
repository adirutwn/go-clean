package usecases

import "gitlab.com/lightnet-thailand/poc/clean-workshop/app/entities"

func (uc *usecase) GetAll() ([]entities.User, error) {
	return uc.UserRepository.GetAll()
}
