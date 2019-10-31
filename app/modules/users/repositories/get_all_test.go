package repositories_test

import (
	"errors"
	gomocket "github.com/Selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/modules/users/repositories"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/test_helpers"
	"testing"
)

func TestRepo_GetAll(t *testing.T) {
	t.Run("Happy", func(t *testing.T) {
		mockedDB := test_helpers.NewMockingDB()
		defer mockedDB.Close()

		mockedSelectedUser := []map[string]interface{}{{
			"firstName":              "Adirut",
			"lastName": 		"Nithilerdviwat",
		}, {
			"firstName": "Peach",
			"lastName": "Jura",
		}}

		gomocket.Catcher.Reset().NewMock().WithQuery(`SELECT * FROM "users"`).WithReply(mockedSelectedUser)

		repo := repositories.NewUserRepository(mockedDB)
		users, err := repo.GetAll()

		assert.NoError(t, err)
		assert.NotEmpty(t, users)
		assert.Equal(t, 2, len(users))
	})

	t.Run("Error - DB returns error", func(t *testing.T) {
		mockedDB := test_helpers.NewMockingDB()
		defer mockedDB.Close()

		gomocket.Catcher.Reset().NewMock().WithQuery(`SELECT * FROM "users"`).WithError(errors.New("error here"))

		repo := repositories.NewUserRepository(mockedDB)
		users, err := repo.GetAll()

		assert.Error(t, err)
		assert.Empty(t, users)
		assert.Equal(t, 0, len(users))
	})
}


