package http_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/domain/mocks"
	httpPaymentType "github.com/San-Kun/MiniProject/user/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPaymentTypeController_GetPaymentTypes(t *testing.T) {
	mockPaymentTypeUsecase := new(mocks.PaymentTypeUsecase)
	mockListPaymentType := &domain.PaymentTypes{
		domain.PaymentType{
			ID:          1,
			PayTypeName: "Transfer Bank BRI / BNI / BCA",
		},
		domain.PaymentType{
			ID:          2,
			PayTypeName: "Via Dana / GOPAY / OVO",
		},
	}

	e := echo.New()
	mockPaymentTypeUsecase.On("ReadAll").Return(mockListPaymentType, nil)
	req, _ := http.NewRequest(echo.GET, "/paymentTypes", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	paymentTypeController := httpPaymentType.PaymentTypeController{
		PaymentTypeUsecase: mockPaymentTypeUsecase,
	}
	err := paymentTypeController.GetPaymentTypes(c)
	require.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, 2, len(responseBody["data"].([]interface{})))

	mockPaymentTypeUsecase.AssertExpectations(t)
}

func TestPaymentTypeController_GetPaymentTypesFailed(t *testing.T) {
	mockPaymentTypeUsecase := new(mocks.PaymentTypeUsecase)

	e := echo.New()
	mockPaymentTypeUsecase.On("ReadAll").Return(nil, errors.New("something error"))
	req, _ := http.NewRequest(echo.GET, "/paymentTypes", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	paymentTypeController := httpPaymentType.PaymentTypeController{
		PaymentTypeUsecase: mockPaymentTypeUsecase,
	}
	err := paymentTypeController.GetPaymentTypes(c)
	require.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	assert.Equal(t, 0, len(responseBody))
	assert.Equal(t, 400, rec.Code)

	mockPaymentTypeUsecase.AssertExpectations(t)
}

func TestPaymentTypeController_GetPaymentTypeByID(t *testing.T) {
	mockPaymentTypeUsecase := new(mocks.PaymentTypeUsecase)
	mockPaymentType := &domain.PaymentType{
		ID:          1,
		PayTypeName: "Transfer Bank BRI / BNI / BCA",
	}

	e := echo.New()
	mockPaymentTypeUsecase.On("ReadByID", mock.AnythingOfType("int")).Return(mockPaymentType, nil)
	req, _ := http.NewRequest(echo.GET, "/paymentTypes/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("paymentTypes/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	paymentTypeController := httpPaymentType.PaymentTypeController{
		PaymentTypeUsecase: mockPaymentTypeUsecase,
	}
	err := paymentTypeController.GetPaymentTypeByID(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, float64(mockPaymentType.ID), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, mockPaymentType.PayTypeName, responseBody["data"].(map[string]interface{})["pay_type_name"])

	mockPaymentTypeUsecase.AssertExpectations(t)
}

func TestPaymentTypeController_GetPaymentTypeByIDNotFound(t *testing.T) {
	mockPaymentTypeUsecase := new(mocks.PaymentTypeUsecase)

	e := echo.New()
	mockPaymentTypeUsecase.On("ReadByID", mock.AnythingOfType("int")).Return(nil, errors.New("not found"))
	req, _ := http.NewRequest(echo.GET, "/paymentTypes/2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("paymentTypes/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	paymentTypeController := httpPaymentType.PaymentTypeController{
		PaymentTypeUsecase: mockPaymentTypeUsecase,
	}
	err := paymentTypeController.GetPaymentTypeByID(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))

	mockPaymentTypeUsecase.AssertExpectations(t)
}

func TestPaymentTypeController_CreatePaymentType(t *testing.T) {
	mockPaymentTypeUsecase := new(mocks.PaymentTypeUsecase)
	mockPaymentType := &domain.PaymentType{
		ID:          1,
		PayTypeName: "Transfer Bank BRI / BNI / BCA",
	}

	e := echo.New()
	mockPaymentTypeUsecase.On("Create", mock.AnythingOfType("request.PaymentTypeCreateRequest")).Return(mockPaymentType, nil)
	req, _ := http.NewRequest(echo.POST, "/paymentTypes", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	paymentTypeController := httpPaymentType.PaymentTypeController{
		PaymentTypeUsecase: mockPaymentTypeUsecase,
	}
	err := paymentTypeController.CreatePaymentType(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, float64(mockPaymentType.ID), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, mockPaymentType.PayTypeName, responseBody["data"].(map[string]interface{})["pay_type_name"])

	mockPaymentTypeUsecase.AssertExpectations(t)
}

func TestPaymentTypeController_CreatePaymentTypeFailed(t *testing.T) {
	mockPaymentTypeUsecase := new(mocks.PaymentTypeUsecase)

	e := echo.New()
	mockPaymentTypeUsecase.On("Create", mock.AnythingOfType("request.PaymentTypeCreateRequest")).Return(nil, errors.New("error something"))
	req, _ := http.NewRequest(echo.POST, "/paymentTypes", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	paymentTypeController := httpPaymentType.PaymentTypeController{
		PaymentTypeUsecase: mockPaymentTypeUsecase,
	}
	err := paymentTypeController.CreatePaymentType(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, false, responseBody["status"])

	mockPaymentTypeUsecase.AssertExpectations(t)
}
