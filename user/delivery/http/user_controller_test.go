package http_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/domain/mocks"
	"github.com/San-Kun/MiniProject/domain/web/response"
	httpUser "github.com/San-Kun/MiniProject/user/delivery/http"
	"github.com/San-Kun/MiniProject/user/delivery/http/helper"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUserController_GetUsers(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecase)
	mockListUser := &domain.Users{
		domain.User{
			ID:       1,
			Name:     "ikhsan",
			Email:    "ikhsan@gmail.com",
			Password: "12345678",
			NoTelp:   "081280948137",
		},
		domain.User{
			ID:       2,
			Name:     "endang",
			Email:    "endang@gmail.com",
			Password: "12345678",
			NoTelp:   "085728372938",
		},
	}

	e := echo.New()
	mockUserUsecase.On("ReadAll").Return(mockListUser, nil)
	req, _ := http.NewRequest(echo.GET, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userController := httpUser.UserController{
		UserUsecase: mockUserUsecase,
	}
	err := userController.GetUsers(c)
	require.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, 2, len(responseBody["data"].([]interface{})))

	mockUserUsecase.AssertExpectations(t)
}

func TestUserController_GetUsersFailed(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecase)

	e := echo.New()
	mockUserUsecase.On("ReadAll").Return(nil, errors.New("something error"))
	req, _ := http.NewRequest(echo.GET, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userController := httpUser.UserController{
		UserUsecase: mockUserUsecase,
	}
	err := userController.GetUsers(c)
	require.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	assert.Equal(t, 0, len(responseBody))
	assert.Equal(t, 400, rec.Code)

	mockUserUsecase.AssertExpectations(t)
}

func TestUserController_GetUserByID(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecase)
	mockUser := &domain.User{
		ID:       1,
		Name:     "ikhsan",
		Email:    "ikhsan@gmail.com",
		Password: "12345678",
		NoTelp:   "081280948137",
	}

	e := echo.New()
	mockUserUsecase.On("ReadByID", mock.AnythingOfType("int")).Return(mockUser, nil)
	req, _ := http.NewRequest(echo.GET, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	userController := httpUser.UserController{
		UserUsecase: mockUserUsecase,
	}
	err := userController.GetUserByID(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, float64(mockUser.ID), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, mockUser.Email, responseBody["data"].(map[string]interface{})["email"])
	assert.Equal(t, mockUser.Password, responseBody["data"].(map[string]interface{})["password"])

	mockUserUsecase.AssertExpectations(t)
}

func TestUserController_GetUserByIDNotFound(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecase)

	e := echo.New()
	mockUserUsecase.On("ReadByID", mock.AnythingOfType("int")).Return(nil, errors.New("not found"))
	req, _ := http.NewRequest(echo.GET, "/users/2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("users/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	userController := httpUser.UserController{
		UserUsecase: mockUserUsecase,
	}
	err := userController.GetUserByID(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))

	mockUserUsecase.AssertExpectations(t)
}

func TestUserController_CreateUser(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecase)
	mockUser := &domain.User{
		ID:       1,
		Name:     "ikhsan",
		Email:    "ikhsan@gmail.com",
		Password: "12345678",
		NoTelp:   "081280948137",
	}

	e := echo.New()
	mockUserUsecase.On("Create", mock.AnythingOfType("request.UserCreateRequest")).Return(mockUser, nil)
	req, _ := http.NewRequest(echo.POST, "/users", strings.NewReader(`{"email": "ikhsan@gmail.com", "password": "12345678" }`))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userController := httpUser.UserController{
		UserUsecase: mockUserUsecase,
	}
	err := userController.CreateUser(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, float64(mockUser.ID), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, mockUser.Email, responseBody["data"].(map[string]interface{})["email"])

	mockUserUsecase.AssertExpectations(t)
}

func TestUserController_CreateUserFailed(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecase)

	e := echo.New()
	mockUserUsecase.On("Create", mock.AnythingOfType("request.UserCreateRequest")).Return(nil, errors.New("error something"))
	req, _ := http.NewRequest(echo.POST, "/users", strings.NewReader(`{"email": "", "password": "12345678" }`))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userController := httpUser.UserController{
		UserUsecase: mockUserUsecase,
	}
	err := userController.CreateUser(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, false, responseBody["status"])

	mockUserUsecase.AssertExpectations(t)
}

func TestUserController_Login(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecase)
	token, _ := helper.NewGoJWT().CreateTokenJWT(1, "ikhsan@gmail.com")
	mockResponse := &response.SuccessLogin{
		ID:    1,
		Email: "ikhsan@gmail.com",
		Token: token,
	}

	e := echo.New()
	mockUserUsecase.On("Login", mock.AnythingOfType("request.LoginRequest")).Return(mockResponse, nil)
	req, _ := http.NewRequest(echo.POST, "/login", strings.NewReader(`{"email": "ikhsan@gmail.com", "password": "12345678" }`))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userController := httpUser.UserController{
		UserUsecase: mockUserUsecase,
	}
	err := userController.Login(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, float64(1), responseBody["id_user"])
	assert.Equal(t, mockResponse.Email, responseBody["email"])

	mockUserUsecase.AssertExpectations(t)
}

func TestUserController_LoginFailed(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecase)

	e := echo.New()
	mockUserUsecase.On("Login", mock.AnythingOfType("request.LoginRequest")).Return(nil, errors.New("something error"))
	req, _ := http.NewRequest(echo.POST, "/login", strings.NewReader(`{"email": "iikhsan@gmail.com", "password": "12345678" }`))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userController := httpUser.UserController{
		UserUsecase: mockUserUsecase,
	}
	err := userController.Login(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 401, int(responseBody["code"].(float64)))

	mockUserUsecase.AssertExpectations(t)
}
