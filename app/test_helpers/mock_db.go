package test_helpers

import (
	"github.com/Selvatico/go-mocket"
	"github.com/jinzhu/gorm"
	"log"
)

func NewMockingDB() *gorm.DB {
	gomocket.Catcher.Register()

	db, err := gorm.Open(gomocket.DriverName, "")
	if err != nil {
		log.Fatalf("error mocking up the database %s", err)
	}

	db.LogMode(true)

	return db
}
