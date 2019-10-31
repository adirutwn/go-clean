package environments

import (
	"fmt"
	"syscall"
)

var (
	Port string
	PostgresHost string
	PostgresPort string
	PostgresUser string
	PostgresPassword string
	PostgresDB string
)

func require(envName string) string {
	env, found := syscall.Getenv(envName)
	if !found {
		panic(fmt.Sprintf("%s env is required", envName))
	}

	return env
}

func Init() {
	Port = require("PORT")
	PostgresHost = require("POSTGRES_HOST")
	PostgresPort = require("POSTGRES_PORT")
	PostgresUser = require("POSTGRES_USER")
	PostgresPassword = require("POSTGRES_PASSWORD")
	PostgresDB = require("POSTGRES_DB")
}