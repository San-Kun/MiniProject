package http_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/domain/mocks"
	httpEvent "github.com/San-Kun/MiniProject/user/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestEventController_GetEvents(t *testing.T) {
	mockEventUsecase := new(mocks.EventUsecase)
	mockListEvent := &domain.Events{
		domain.Event{
			ID:           1,
			EventName:    "Forum Diskusi",
			EventTanggal: "12 Maret 2022",
			Capacity:     "100",
		},
		domain.Event{
			ID:           2,
			EventName:    "Workshop Kebhinekaan",
			EventTanggal: "20 Maret 2022",
			Capacity:     "150",
		},
	}

	e := echo.New()
	mockEventUsecase.On("ReadAll").Return(mockListEvent, nil)
	req, _ := http.NewRequest(echo.GET, "/events", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	eventController := httpEvent.EventController{
		EventUsecase: mockEventUsecase,
	}
	err := eventController.GetEvents(c)
	require.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, 2, len(responseBody["data"].([]interface{})))

	mockEventUsecase.AssertExpectations(t)
}

func TestEventController_GetEventsFailed(t *testing.T) {
	mockEventUsecase := new(mocks.EventUsecase)

	e := echo.New()
	mockEventUsecase.On("ReadAll").Return(nil, errors.New("something error"))
	req, _ := http.NewRequest(echo.GET, "/events", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	eventController := httpEvent.EventController{
		EventUsecase: mockEventUsecase,
	}
	err := eventController.GetEvents(c)
	require.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	assert.Equal(t, 0, len(responseBody))
	assert.Equal(t, 400, rec.Code)

	mockEventUsecase.AssertExpectations(t)
}

func TestEventController_GetEventByID(t *testing.T) {
	mockEventUsecase := new(mocks.EventUsecase)
	mockEvent := &domain.Event{
		ID:           1,
		EventName:    "Forum Diskusi",
		EventTanggal: "12 Maret 2022",
		Capacity:     "100",
	}

	e := echo.New()
	mockEventUsecase.On("ReadByID", mock.AnythingOfType("int")).Return(mockEvent, nil)
	req, _ := http.NewRequest(echo.GET, "/events/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("events/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	eventController := httpEvent.EventController{
		EventUsecase: mockEventUsecase,
	}
	err := eventController.GetEventByID(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, float64(mockEvent.ID), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, mockEvent.EventName, responseBody["data"].(map[string]interface{})["event_name"])
	assert.Equal(t, mockEvent.EventTanggal, responseBody["data"].(map[string]interface{})["event_tanggal"])
	assert.Equal(t, mockEvent.Capacity, responseBody["data"].(map[string]interface{})["capacity"])

	mockEventUsecase.AssertExpectations(t)
}

func TestEventController_GetEventByIDNotFound(t *testing.T) {
	mockEventUsecase := new(mocks.EventUsecase)

	e := echo.New()
	mockEventUsecase.On("ReadByID", mock.AnythingOfType("int")).Return(nil, errors.New("not found"))
	req, _ := http.NewRequest(echo.GET, "/events/2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("events/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	eventController := httpEvent.EventController{
		EventUsecase: mockEventUsecase,
	}
	err := eventController.GetEventByID(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))

	mockEventUsecase.AssertExpectations(t)
}

func TestEventController_CreateEvent(t *testing.T) {
	mockEventUsecase := new(mocks.EventUsecase)
	mockEvent := &domain.Event{
		ID:           1,
		EventName:    "Forum Diskusi",
		EventTanggal: "12 Maret 2022",
		Capacity:     "100",
	}

	e := echo.New()
	mockEventUsecase.On("Create", mock.AnythingOfType("request.EventCreateRequest")).Return(mockEvent, nil)
	req, _ := http.NewRequest(echo.POST, "/events", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	eventController := httpEvent.EventController{
		EventUsecase: mockEventUsecase,
	}
	err := eventController.CreateEvent(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, float64(mockEvent.ID), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, mockEvent.EventName, responseBody["data"].(map[string]interface{})["event_name"])
	assert.Equal(t, mockEvent.EventTanggal, responseBody["data"].(map[string]interface{})["event_tanggal"])
	assert.Equal(t, mockEvent.Capacity, responseBody["data"].(map[string]interface{})["capacity"])

	mockEventUsecase.AssertExpectations(t)
}

func TestEventController_CreateEventFailed(t *testing.T) {
	mockEventUsecase := new(mocks.EventUsecase)

	e := echo.New()
	mockEventUsecase.On("Create", mock.AnythingOfType("request.EventCreateRequest")).Return(nil, errors.New("error something"))
	req, _ := http.NewRequest(echo.POST, "/events", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	eventController := httpEvent.EventController{
		EventUsecase: mockEventUsecase,
	}
	err := eventController.CreateEvent(c)
	assert.NoError(t, err)
	var responseBody map[string]interface{}
	resBody := rec.Body.String()
	err = json.Unmarshal([]byte(resBody), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, false, responseBody["status"])

	mockEventUsecase.AssertExpectations(t)
}
