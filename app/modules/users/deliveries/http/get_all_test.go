package http_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	_userHttp "gitlab.com/lightnet-thailand/poc/clean-workshop/app/modules/users/deliveries/http"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/modules/users/mocks"
	"gitlab.com/lightnet-thailand/poc/clean-workshop/app/test_helpers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_GetAll(t *testing.T) {
	t.Run("Happy", func(t *testing.T) {
		mockedUserEntity := test_helpers.UsersEntity

		mockUseCase := new(mocks.UseCase)
		mockUseCase.On("GetAll").Return(mockedUserEntity, nil)

		ginEngine := gin.Default()
		_userHttp.NewEndpointHttpHandler(ginEngine, mockUseCase)

		requestUrl := "/v1/users"
		req, err := http.NewRequest("GET", requestUrl, nil)
		assert.NoError(t, err)

		res := httptest.NewRecorder()
		ginEngine.ServeHTTP(res, req)

		var respBody map[string]interface{}
		_ = json.Unmarshal(res.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "success", respBody["status"].(string))
	})

	t.Run("Error - UseCase return error", func(t *testing.T) {
		mockUseCase := new(mocks.UseCase)
		mockUseCase.On("GetAll").Return(nil, errors.New("error here"))

		ginEngine := gin.Default()
		_userHttp.NewEndpointHttpHandler(ginEngine, mockUseCase)

		requestUrl := "/v1/users"
		req, err := http.NewRequest("GET", requestUrl, nil)
		assert.NoError(t, err)

		res := httptest.NewRecorder()
		ginEngine.ServeHTTP(res, req)

		var respBody map[string]interface{}
		_ = json.Unmarshal(res.Body.Bytes(), &respBody)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "fail", respBody["status"].(string))
	})
}
