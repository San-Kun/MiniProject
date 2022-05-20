package http_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/domain/mocks"
	httpEventTicket "github.com/San-Kun/MiniProject/user/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestEventTicketController_GetEventTickets(t *testing.T) {
	mockEventTicketUsecase := new(mocks.EventTicketUsecase)
	mockListEventTicket := &domain.EventTickets{
		domain.EventTicket{
			ID:          1,
			EventID:     1,
			PaymentID:   1,
			PayTypeID:   1,
			PayTypeName: "Transfer Bank BRI / BNI / BCA",
		},
		domain.EventTicket{
			ID:          2,
			EventID:     2,
			PaymentID:   2,
			PayTypeID:   2,
			PayTypeName: "Via Dana / GOPAY / OVO",
		},
	}

	e := echo.New()
	mockEventTicketUsecase.On("ReadAll").Return(mockListEventTicket, nil)
	req, _ := http.NewRequest(echo.GET, "/eventTickets", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	eventTicketController := httpEventTicket.EventTicketController{
		EventTicketUsecase: mockEventTicketUsecase,
	}
	err := eventTicketController.GetEventTickets(c)
	require.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, 2, len(responseBody["data"].([]interface{})))

	mockEventTicketUsecase.AssertExpectations(t)
}

func TestEventTicketController_GetEventTicketsFailed(t *testing.T) {
	mockEventTicketUsecase := new(mocks.EventTicketUsecase)

	e := echo.New()
	mockEventTicketUsecase.On("ReadAll").Return(nil, errors.New("something error"))
	req, _ := http.NewRequest(echo.GET, "/eventTickets", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	eventTicketController := httpEventTicket.EventTicketController{
		EventTicketUsecase: mockEventTicketUsecase,
	}
	err := eventTicketController.GetEventTickets(c)
	require.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	assert.Equal(t, 0, len(responseBody))
	assert.Equal(t, 400, rec.Code)

	mockEventTicketUsecase.AssertExpectations(t)
}

func TestEventTicketController_GetEventTicketByID(t *testing.T) {
	mockEventTicketUsecase := new(mocks.EventTicketUsecase)
	mockEventTicket := &domain.EventTicket{
		ID:          1,
		EventID:     1,
		PaymentID:   1,
		PayTypeID:   1,
		PayTypeName: "Transfer Bank BRI / BNI / BCA",
	}

	e := echo.New()
	mockEventTicketUsecase.On("ReadByID", mock.AnythingOfType("int")).Return(mockEventTicket, nil)
	req, _ := http.NewRequest(echo.GET, "/eventTickets/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("eventTickets/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	eventTicketController := httpEventTicket.EventTicketController{
		EventTicketUsecase: mockEventTicketUsecase,
	}
	err := eventTicketController.GetEventTicketByID(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, float64(mockEventTicket.ID), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, float64(mockEventTicket.ID), responseBody["data"].(map[string]interface{})["event_id"])
	assert.Equal(t, float64(mockEventTicket.ID), responseBody["data"].(map[string]interface{})["payment_id"])
	assert.Equal(t, float64(mockEventTicket.ID), responseBody["data"].(map[string]interface{})["paytype_id"])
	assert.Equal(t, mockEventTicket.PayTypeName, responseBody["data"].(map[string]interface{})["pay_type_name"])

	mockEventTicketUsecase.AssertExpectations(t)
}

func TestEventTicketController_GetEventTicketByIDNotFound(t *testing.T) {
	mockEventTicketUsecase := new(mocks.EventTicketUsecase)

	e := echo.New()
	mockEventTicketUsecase.On("ReadByID", mock.AnythingOfType("int")).Return(nil, errors.New("not found"))
	req, _ := http.NewRequest(echo.GET, "/eventTickets/2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("eventTickets/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	eventTicketController := httpEventTicket.EventTicketController{
		EventTicketUsecase: mockEventTicketUsecase,
	}
	err := eventTicketController.GetEventTicketByID(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))

	mockEventTicketUsecase.AssertExpectations(t)
}

func TestEventTicketController_CreateEventTicket(t *testing.T) {
	mockEventTicketUsecase := new(mocks.EventTicketUsecase)
	mockEventTicket := &domain.EventTicket{
		ID:          1,
		EventID:     1,
		PaymentID:   1,
		PayTypeID:   1,
		PayTypeName: "Transfer Bank BRI / BNI / BCA",
	}

	e := echo.New()
	mockEventTicketUsecase.On("Create", mock.AnythingOfType("request.EventTicketCreateRequest")).Return(mockEventTicket, nil)
	req, _ := http.NewRequest(echo.POST, "/eventTickets", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	eventTicketController := httpEventTicket.EventTicketController{
		EventTicketUsecase: mockEventTicketUsecase,
	}
	err := eventTicketController.CreateEventTicket(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, float64(mockEventTicket.ID), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, float64(mockEventTicket.ID), responseBody["data"].(map[string]interface{})["event_id"])
	assert.Equal(t, float64(mockEventTicket.ID), responseBody["data"].(map[string]interface{})["payment_id"])
	assert.Equal(t, float64(mockEventTicket.ID), responseBody["data"].(map[string]interface{})["paytype_id"])
	assert.Equal(t, mockEventTicket.PayTypeName, responseBody["data"].(map[string]interface{})["pay_type_name"])

	mockEventTicketUsecase.AssertExpectations(t)
}

func TestEventTicketController_CreateEventTicketFailed(t *testing.T) {
	mockEventTicketUsecase := new(mocks.EventTicketUsecase)

	e := echo.New()
	mockEventTicketUsecase.On("Create", mock.AnythingOfType("request.EventTicketCreateRequest")).Return(nil, errors.New("error something"))
	req, _ := http.NewRequest(echo.POST, "/eventTickets", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	eventTicketController := httpEventTicket.EventTicketController{
		EventTicketUsecase: mockEventTicketUsecase,
	}
	err := eventTicketController.CreateEventTicket(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, false, responseBody["status"])

	mockEventTicketUsecase.AssertExpectations(t)
}
