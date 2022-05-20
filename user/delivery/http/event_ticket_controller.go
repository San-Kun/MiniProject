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

type EventTicketController struct {
	EventTicketUsecase domain.EventTicketUsecase
}

func NewEventTicketController(e *echo.Echo, Usecase domain.EventTicketUsecase) {
	EventTicketController := &EventTicketController{
		EventTicketUsecase: Usecase,
	}

	authMiddleware := mid.NewGoMiddleware().AuthMiddleware()
	e.GET("/eventTickets/:id", EventTicketController.GetEventTicketByID, authMiddleware)
	e.GET("/eventTickets", EventTicketController.GetEventTickets, authMiddleware)
	e.DELETE("/eventTickets/:id", EventTicketController.DeleteEventTickets, authMiddleware)
	e.PUT("/eventTickets/:id", EventTicketController.UpdateEventTickets, authMiddleware)
	e.POST("/eventTickets", EventTicketController.CreateEventTicket)
}

func (u *EventTicketController) CreateEventTicket(c echo.Context) error {
	var req request.EventTicketCreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdEventTicket, err := u.EventTicketUsecase.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.EventTicketCreateResponse{
		ID:          int(createdEventTicket.ID),
		EventID:     int(createdEventTicket.EventID),
		PaymentID:   int(createdEventTicket.PaymentID),
		PayTypeID:   int(createdEventTicket.PayTypeID),
		PayTypeName: createdEventTicket.PayTypeName,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":   201,
		"status": true,
		"data":   res,
	})
}

func (u *EventTicketController) GetEventTicketByID(c echo.Context) error {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	foundEventTicket, err := u.EventTicketUsecase.ReadByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.EventTicketsResponse{
		ID:          int(foundEventTicket.ID),
		EventID:     int(foundEventTicket.EventID),
		PaymentID:   int(foundEventTicket.PaymentID),
		PayTypeID:   int(foundEventTicket.PayTypeID),
		PayTypeName: foundEventTicket.PayTypeName,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *EventTicketController) GetEventTickets(c echo.Context) error {
	foundEventTickets, err := u.EventTicketUsecase.ReadAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var res []response.EventTicketsResponse
	for _, foundEventTicket := range *foundEventTickets {
		res = append(res, response.EventTicketsResponse{
			ID:          int(foundEventTicket.ID),
			EventID:     int(foundEventTicket.EventID),
			PaymentID:   int(foundEventTicket.PaymentID),
			PayTypeID:   int(foundEventTicket.PayTypeID),
			PayTypeName: foundEventTicket.PayTypeName,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *EventTicketController) DeleteEventTickets(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, e := u.EventTicketUsecase.Delete(id)

	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	config.DB.Unscoped().Delete(&domain.EventTicket{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete event ticket with id " + strconv.Itoa(id),
		"code":     200,
	})
}

func (u *EventTicketController) UpdateEventTickets(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updateEventTicket := domain.EventTicket{}
	err = c.Bind(&updateEventTicket)
	if err != nil {
		return err
	}

	if err := config.DB.Model(&domain.EventTicket{}).Where("id = ?", id).Updates(domain.EventTicket{
		EventID:     updateEventTicket.EventID,
		PaymentID:   updateEventTicket.PaymentID,
		PayTypeID:   updateEventTicket.PayTypeID,
		PayTypeName: updateEventTicket.PayTypeName,
	}).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
			"code":     400,
		})
	}
	foundEventTicket, _ := u.EventTicketUsecase.ReadByID(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update event ticket with id " + strconv.Itoa(id),
		"code":     200,
		"data":     foundEventTicket,
	})
}
