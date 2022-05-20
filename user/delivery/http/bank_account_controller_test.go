package http_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/domain/mocks"
	httpBankAccount "github.com/San-Kun/MiniProject/user/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestBankAccountController_GetBankAccounts(t *testing.T) {
	mockBankAccountUsecase := new(mocks.BankAccountUsecase)
	mockListBankAccount := &domain.BankAccounts{
		domain.BankAccount{
			ID:            1,
			UserID:        1,
			BankID:        1,
			AccountNumber: "987654321",
			AccountName:   "IKHSAN ENDANG PRASETYA",
		},
		domain.BankAccount{
			ID:            2,
			UserID:        2,
			BankID:        2,
			AccountNumber: "1234567890",
			AccountName:   "ENDANG PRASETYA",
		},
	}

	e := echo.New()
	mockBankAccountUsecase.On("ReadAll").Return(mockListBankAccount, nil)
	req, _ := http.NewRequest(echo.GET, "/bankAccounts", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bankAccountController := httpBankAccount.BankAccountController{
		BankAccountUsecase: mockBankAccountUsecase,
	}
	err := bankAccountController.GetBankAccounts(c)
	require.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, 2, len(responseBody["data"].([]interface{})))

	mockBankAccountUsecase.AssertExpectations(t)
}

func TestBankAccountController_GetBankAccountsFailed(t *testing.T) {
	mockBankAccountUsecase := new(mocks.BankAccountUsecase)

	e := echo.New()
	mockBankAccountUsecase.On("ReadAll").Return(nil, errors.New("something error"))
	req, _ := http.NewRequest(echo.GET, "/bankAccounts", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bankAccountController := httpBankAccount.BankAccountController{
		BankAccountUsecase: mockBankAccountUsecase,
	}
	err := bankAccountController.GetBankAccounts(c)
	require.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	assert.Equal(t, 0, len(responseBody))
	assert.Equal(t, 400, rec.Code)

	mockBankAccountUsecase.AssertExpectations(t)
}

func TestBankAccountController_GetBankAccountByID(t *testing.T) {
	mockBankAccountUsecase := new(mocks.BankAccountUsecase)
	mockBankAccount := &domain.BankAccount{
		ID:            1,
		UserID:        1,
		BankID:        1,
		AccountNumber: "987654321",
		AccountName:   "IKHSAN ENDANG PRASETYA",
	}

	e := echo.New()
	mockBankAccountUsecase.On("ReadByID", mock.AnythingOfType("int")).Return(mockBankAccount, nil)
	req, _ := http.NewRequest(echo.GET, "/bankAccounts/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("bankAccounts/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	bankAccountController := httpBankAccount.BankAccountController{
		BankAccountUsecase: mockBankAccountUsecase,
	}
	err := bankAccountController.GetBankAccountByID(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, float64(mockBankAccount.ID), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, mockBankAccount.UserID, responseBody["data"].(map[string]interface{})["user_id"])
	assert.Equal(t, mockBankAccount.BankID, responseBody["data"].(map[string]interface{})["bank_id"])
	assert.Equal(t, mockBankAccount.AccountNumber, responseBody["data"].(map[string]interface{})["account_number"])
	assert.Equal(t, mockBankAccount.AccountName, responseBody["data"].(map[string]interface{})["account_name"])

	mockBankAccountUsecase.AssertExpectations(t)
}

func TestBankAccountController_GetBankAccountByIDNotFound(t *testing.T) {
	mockBankAccountUsecase := new(mocks.BankAccountUsecase)

	e := echo.New()
	mockBankAccountUsecase.On("ReadByID", mock.AnythingOfType("int")).Return(nil, errors.New("not found"))
	req, _ := http.NewRequest(echo.GET, "/bankAccounts/2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("bankAccounts/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	bankAccountController := httpBankAccount.BankAccountController{
		BankAccountUsecase: mockBankAccountUsecase,
	}
	err := bankAccountController.GetBankAccountByID(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))

	mockBankAccountUsecase.AssertExpectations(t)
}

func TestBankAccountController_CreateBankAccount(t *testing.T) {
	mockBankAccountUsecase := new(mocks.BankAccountUsecase)
	mockBankAccount := &domain.BankAccount{
		ID:            1,
		UserID:        1,
		BankID:        1,
		AccountNumber: "987654321",
		AccountName:   "IKHSAN ENDANG PRASETYA",
	}

	e := echo.New()
	mockBankAccountUsecase.On("Create", mock.AnythingOfType("request.BankAccountCreateRequest")).Return(mockBankAccount, nil)
	req, _ := http.NewRequest(echo.POST, "/bankAccounts", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bankAccountController := httpBankAccount.BankAccountController{
		BankAccountUsecase: mockBankAccountUsecase,
	}
	err := bankAccountController.CreateBankAccount(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, float64(mockBankAccount.ID), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, mockBankAccount.UserID, responseBody["data"].(map[string]interface{})["user_id"])
	assert.Equal(t, mockBankAccount.BankID, responseBody["data"].(map[string]interface{})["bank_id"])
	assert.Equal(t, mockBankAccount.AccountNumber, responseBody["data"].(map[string]interface{})["account_number"])
	assert.Equal(t, mockBankAccount.AccountName, responseBody["data"].(map[string]interface{})["account_name"])

	mockBankAccountUsecase.AssertExpectations(t)
}

func TestBankAccountController_CreateBankAccountFailed(t *testing.T) {
	mockBankAccountUsecase := new(mocks.BankAccountUsecase)

	e := echo.New()
	mockBankAccountUsecase.On("Create", mock.AnythingOfType("request.BankAccountCreateRequest")).Return(nil, errors.New("error something"))
	req, _ := http.NewRequest(echo.POST, "/bankAccounts", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bankAccountController := httpBankAccount.BankAccountController{
		BankAccountUsecase: mockBankAccountUsecase,
	}
	err := bankAccountController.CreateBankAccount(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, false, responseBody["status"])

	mockBankAccountUsecase.AssertExpectations(t)
}
