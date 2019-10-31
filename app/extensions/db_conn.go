package extensions

import (
	"fmt"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"math"
)

func ConnectDB(host, port, user, password, dbName string) *gorm.DB {
	connection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName)
	db, err := gorm.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("database is connected")
	}

	return db
}

func MigrateDB(host, port, user, password, dbName, migrationFolder string) {
	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbName)

	migrationEngine, err := migrate.New("file://../app/"+migrationFolder, databaseURL)
	if err != nil {
		panic(err)
	}
	_ = migrationEngine.Steps(math.MaxInt32)
}


