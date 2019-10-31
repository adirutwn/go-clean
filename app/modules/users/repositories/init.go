package repositories

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/modules/users"
)

type repo struct {
	DB *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) users.Repository {
	return &repo{
		DB: dbConn,
	}
}
