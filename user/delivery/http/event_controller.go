package http

import (
	"net/http"
	"strconv"

	"github.com/San-Kun/MiniProject/app/config"
	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/domain/web/request"
	"github.com/San-Kun/MiniProject/domain/web/response"
	mid "github.com/San-Kun/MiniProject/user/delivery/http/middleware"
	"github.com/labstack/echo/v4"
)

type EventController struct {
	EventUsecase domain.EventUsecase
}

func NewEventController(e *echo.Echo, Usecase domain.EventUsecase) {
	EventController := &EventController{
		EventUsecase: Usecase,
	}

	authMiddleware := mid.NewGoMiddleware().AuthMiddleware()
	e.GET("/events/:id", EventController.GetEventByID, authMiddleware)
	e.GET("/events", EventController.GetEvents, authMiddleware)
	e.DELETE("/events/:id", EventController.DeleteEvents, authMiddleware)
	e.PUT("/events/:id", EventController.UpdateEvents, authMiddleware)
	e.POST("/events", EventController.CreateEvent)
}

func (u *EventController) CreateEvent(c echo.Context) error {
	var req request.EventCreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdEvent, err := u.EventUsecase.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.EventCreateResponse{
		ID:           int(createdEvent.ID),
		EventName:    createdEvent.EventName,
		EventTanggal: createdEvent.EventTanggal,
		Capacity:     createdEvent.Capacity,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":   201,
		"status": true,
		"data":   res,
	})
}

func (u *EventController) GetEventByID(c echo.Context) error {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	foundEvent, err := u.EventUsecase.ReadByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.EventResponse{
		ID:           int(foundEvent.ID),
		EventName:    foundEvent.EventName,
		EventTanggal: foundEvent.EventTanggal,
		Capacity:     foundEvent.Capacity,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *EventController) GetEvents(c echo.Context) error {
	foundEvents, err := u.EventUsecase.ReadAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var res []response.EventResponse
	for _, foundEvent := range *foundEvents {
		res = append(res, response.EventResponse{
			ID:           int(foundEvent.ID),
			EventName:    foundEvent.EventName,
			EventTanggal: foundEvent.EventTanggal,
			Capacity:     foundEvent.Capacity,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *EventController) DeleteEvents(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, e := u.EventUsecase.Delete(id)

	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	config.DB.Unscoped().Delete(&domain.Event{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete event with id " + strconv.Itoa(id),
		"code":     200,
	})
}

func (u *EventController) UpdateEvents(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updateEvent := domain.Event{}
	err = c.Bind(&updateEvent)
	if err != nil {
		return err
	}

	if err := config.DB.Model(&domain.Event{}).Where("id = ?", id).Updates(domain.Event{
		EventName:    updateEvent.EventName,
		EventTanggal: updateEvent.EventTanggal,
		Capacity:     updateEvent.Capacity,
	}).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
			"code":     400,
		})
	}
	foundEvent, _ := u.EventUsecase.ReadByID(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update event with id " + strconv.Itoa(id),
		"code":     200,
		"data":     foundEvent,
	})
}
