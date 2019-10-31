package usecases

import "gitlab.com/lightnet-thailand/poc/clean-workshop/app/modules/users"

type usecase struct {
	UserRepository users.Repository
}

func NewUserUsecase(userRepository users.Repository) users.UseCase {
	return &usecase{
		UserRepository: userRepository,
	}
}
