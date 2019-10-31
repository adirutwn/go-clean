package usecases_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/modules/users/mocks"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/modules/users/usecases"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/test_helpers"
	"testing"
)

func TestUsecase_GetAll(t *testing.T) {
	t.Run("Happy", func(t *testing.T) {
		mockedUsers := test_helpers.UsersEntity

		mockedRepo := new(mocks.Repository)
		mockedRepo.On("GetAll").Return(mockedUsers, nil)

		uc := usecases.NewUserUsecase(mockedRepo)
		users, err := uc.GetAll()

		assert.NoError(t, err)
		assert.NotEmpty(t, users)
		assert.Equal(t, 2, len(users))
	})

	t.Run("Error - repository returns error", func(t *testing.T) {
		mockedRepo := new(mocks.Repository)
		mockedRepo.On("GetAll").Return(nil, errors.New("error here"))

		uc := usecases.NewUserUsecase(mockedRepo)
		users, err := uc.GetAll()

		assert.Error(t, err)
		assert.Empty(t, users)
		assert.Equal(t, 0, len(users))
	})
}
