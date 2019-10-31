package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/modules/users/deliveries/http/models"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/utils"
	"net/http"
)

func (uh *userHandler) GetAll(c *gin.Context) {
	users, err := uh.UserUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	var usersResponse []models.UserResponse
	for _, user := range users {
		tmpModel, err := new(models.UserResponse).ParseEntity(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
			return
		}

		usersResponse = append(usersResponse, *tmpModel)
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(usersResponse))
}
