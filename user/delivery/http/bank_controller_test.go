package http_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/domain/mocks"
	httpBank "github.com/San-Kun/MiniProject/user/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestBankController_GetBanks(t *testing.T) {
	mockBankUsecase := new(mocks.BankUsecase)
	mockListBank := &domain.Banks{
		domain.Bank{
			ID:       1,
			BankName: "BANK BRI",
		},
		domain.Bank{
			ID:       2,
			BankName: "BANK BCA",
		},
	}

	e := echo.New()
	mockBankUsecase.On("ReadAll").Return(mockListBank, nil)
	req, _ := http.NewRequest(echo.GET, "/banks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bankController := httpBank.BankController{
		BankUsecase: mockBankUsecase,
	}
	err := bankController.GetBanks(c)
	require.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, 2, len(responseBody["data"].([]interface{})))

	mockBankUsecase.AssertExpectations(t)
}

func TestBankController_GetBanksFailed(t *testing.T) {
	mockBankUsecase := new(mocks.BankUsecase)

	e := echo.New()
	mockBankUsecase.On("ReadAll").Return(nil, errors.New("something error"))
	req, _ := http.NewRequest(echo.GET, "/banks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bankController := httpBank.BankController{
		BankUsecase: mockBankUsecase,
	}
	err := bankController.GetBanks(c)
	require.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	assert.Equal(t, 0, len(responseBody))
	assert.Equal(t, 400, rec.Code)

	mockBankUsecase.AssertExpectations(t)
}

func TestBankController_GetBankByID(t *testing.T) {
	mockBankUsecase := new(mocks.BankUsecase)
	mockBank := &domain.Bank{
		ID:       1,
		BankName: "BANK BRI",
	}

	e := echo.New()
	mockBankUsecase.On("ReadByID", mock.AnythingOfType("int")).Return(mockBank, nil)
	req, _ := http.NewRequest(echo.GET, "/banks/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("banks/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	bankController := httpBank.BankController{
		BankUsecase: mockBankUsecase,
	}
	err := bankController.GetBankByID(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, float64(mockBank.ID), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, mockBank.BankName, responseBody["data"].(map[string]interface{})["bank_name"])

	mockBankUsecase.AssertExpectations(t)
}

func TestBankController_GetBankByIDNotFound(t *testing.T) {
	mockBankUsecase := new(mocks.BankUsecase)

	e := echo.New()
	mockBankUsecase.On("ReadByID", mock.AnythingOfType("int")).Return(nil, errors.New("not found"))
	req, _ := http.NewRequest(echo.GET, "/banks/2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("banks/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	bankController := httpBank.BankController{
		BankUsecase: mockBankUsecase,
	}
	err := bankController.GetBankByID(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))

	mockBankUsecase.AssertExpectations(t)
}

func TestBankController_CreateBank(t *testing.T) {
	mockBankUsecase := new(mocks.BankUsecase)
	mockBank := &domain.Bank{
		ID:       1,
		BankName: "BANK BRI",
	}

	e := echo.New()
	mockBankUsecase.On("Create", mock.AnythingOfType("request.BankCreateRequest")).Return(mockBank, nil)
	req, _ := http.NewRequest(echo.POST, "/banks", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bankController := httpBank.BankController{
		BankUsecase: mockBankUsecase,
	}
	err := bankController.CreateBank(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, float64(mockBank.ID), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, mockBank.BankName, responseBody["data"].(map[string]interface{})["bank_name"])

	mockBankUsecase.AssertExpectations(t)
}

func TestBankController_CreateBankFailed(t *testing.T) {
	mockBankUsecase := new(mocks.BankUsecase)

	e := echo.New()
	mockBankUsecase.On("Create", mock.AnythingOfType("request.BankCreateRequest")).Return(nil, errors.New("error something"))
	req, _ := http.NewRequest(echo.POST, "/banks", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bankController := httpBank.BankController{
		BankUsecase: mockBankUsecase,
	}
	err := bankController.CreateBank(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, false, responseBody["status"])

	mockBankUsecase.AssertExpectations(t)
}
