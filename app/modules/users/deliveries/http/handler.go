package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/modules/users"
)

type userHandler struct {
	UserUsecase users.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, userUsecase users.UseCase) {
	handler := userHandler{
		UserUsecase: userUsecase,
	}

	v1 := ginEngine.Group("v1")
	v1.GET("/users", handler.GetAll)
}