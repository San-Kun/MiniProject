package http_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/domain/mocks"
	httpPayment "github.com/San-Kun/MiniProject/user/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPaymentController_GetPayments(t *testing.T) {
	mockPaymentUsecase := new(mocks.PaymentUsecase)
	mockListPayment := &domain.Payments{
		domain.Payment{
			ID:            1,
			PaymentDate:   "20 Maret 2022",
			PaymentAmount: 45000,
		},
		domain.Payment{
			ID:            2,
			PaymentDate:   "20 Maret 2022",
			PaymentAmount: 50000,
		},
	}

	e := echo.New()
	mockPaymentUsecase.On("ReadAll").Return(mockListPayment, nil)
	req, _ := http.NewRequest(echo.GET, "/payments", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	paymentController := httpPayment.PaymentController{
		PaymentUsecase: mockPaymentUsecase,
	}
	err := paymentController.GetPayments(c)
	require.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, 2, len(responseBody["data"].([]interface{})))

	mockPaymentUsecase.AssertExpectations(t)
}

func TestPaymentController_GetPaymentsFailed(t *testing.T) {
	mockPaymentUsecase := new(mocks.PaymentUsecase)

	e := echo.New()
	mockPaymentUsecase.On("ReadAll").Return(nil, errors.New("something error"))
	req, _ := http.NewRequest(echo.GET, "/payments", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	paymentController := httpPayment.PaymentController{
		PaymentUsecase: mockPaymentUsecase,
	}
	err := paymentController.GetPayments(c)
	require.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	assert.Equal(t, 0, len(responseBody))
	assert.Equal(t, 400, rec.Code)

	mockPaymentUsecase.AssertExpectations(t)
}

func TestPaymentController_GetPaymentByID(t *testing.T) {
	mockPaymentUsecase := new(mocks.PaymentUsecase)
	mockPayment := &domain.Payment{
		ID:            1,
		PaymentDate:   "20 Maret 2022",
		PaymentAmount: 45000,
	}

	e := echo.New()
	mockPaymentUsecase.On("ReadByID", mock.AnythingOfType("int")).Return(mockPayment, nil)
	req, _ := http.NewRequest(echo.GET, "/payments/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("payments/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	paymentController := httpPayment.PaymentController{
		PaymentUsecase: mockPaymentUsecase,
	}
	err := paymentController.GetPaymentByID(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, float64(mockPayment.ID), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, mockPayment.PaymentDate, responseBody["data"].(map[string]interface{})["payment_date"])
	assert.Equal(t, mockPayment.PaymentAmount, responseBody["data"].(map[string]interface{})["payment_amount"])

	mockPaymentUsecase.AssertExpectations(t)
}

func TestPaymentController_GetPaymentByIDNotFound(t *testing.T) {
	mockPaymentUsecase := new(mocks.PaymentUsecase)

	e := echo.New()
	mockPaymentUsecase.On("ReadByID", mock.AnythingOfType("int")).Return(nil, errors.New("not found"))
	req, _ := http.NewRequest(echo.GET, "/payments/2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("payments/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	paymentController := httpPayment.PaymentController{
		PaymentUsecase: mockPaymentUsecase,
	}
	err := paymentController.GetPaymentByID(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))

	mockPaymentUsecase.AssertExpectations(t)
}

func TestPaymentController_CreatePayment(t *testing.T) {
	mockPaymentUsecase := new(mocks.PaymentUsecase)
	mockPayment := &domain.Payment{
		ID:            1,
		PaymentDate:   "20 Maret 2022",
		PaymentAmount: 45000,
	}

	e := echo.New()
	mockPaymentUsecase.On("Create", mock.AnythingOfType("request.PaymentCreateRequest")).Return(mockPayment, nil)
	req, _ := http.NewRequest(echo.POST, "/payments", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	paymentController := httpPayment.PaymentController{
		PaymentUsecase: mockPaymentUsecase,
	}
	err := paymentController.CreatePayment(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, float64(mockPayment.ID), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, mockPayment.PaymentDate, responseBody["data"].(map[string]interface{})["payment_date"])
	assert.Equal(t, mockPayment.PaymentAmount, responseBody["data"].(map[string]interface{})["payment_amount"])

	mockPaymentUsecase.AssertExpectations(t)
}

func TestPaymentController_CreatePaymentFailed(t *testing.T) {
	mockPaymentUsecase := new(mocks.PaymentUsecase)

	e := echo.New()
	mockPaymentUsecase.On("Create", mock.AnythingOfType("request.PaymentCreateRequest")).Return(nil, errors.New("error something"))
	req, _ := http.NewRequest(echo.POST, "/payments", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	paymentController := httpPayment.PaymentController{
		PaymentUsecase: mockPaymentUsecase,
	}
	err := paymentController.CreatePayment(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, false, responseBody["status"])

	mockPaymentUsecase.AssertExpectations(t)
}
