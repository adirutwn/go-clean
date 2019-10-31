package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/environments"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/extensions"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/modules/users/deliveries/http"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/modules/users/repositories"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/modules/users/usecases"
)

func main() {
	environments.Init()

	db := extensions.ConnectDB(
		environments.PostgresHost,
		environments.PostgresPort,
		environments.PostgresUser,
		environments.PostgresPassword,
		environments.PostgresDB,
	)

	extensions.MigrateDB(
		environments.PostgresHost,
		environments.PostgresPort,
		environments.PostgresUser,
		environments.PostgresPassword,
		environments.PostgresDB,
		"migrations",
	)

	ginEngine := gin.New()

	userRepository := repositories.NewUserRepository(db)
	userUseCase := usecases.NewUserUsecase(userRepository)
	http.NewEndpointHttpHandler(ginEngine, userUseCase)

	err := ginEngine.Run(fmt.Sprintf(":%s", environments.Port))
	if err != nil {
		panic(err)
	}
}
