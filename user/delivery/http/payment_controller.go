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

type PaymentController struct {
	PaymentUsecase domain.PaymentUsecase
}

func NewPaymentController(e *echo.Echo, Usecase domain.PaymentUsecase) {
	PaymentController := &PaymentController{
		PaymentUsecase: Usecase,
	}

	authMiddleware := mid.NewGoMiddleware().AuthMiddleware()
	e.GET("/payments/:id", PaymentController.GetPaymentByID, authMiddleware)
	e.GET("/payments", PaymentController.GetPayments, authMiddleware)
	e.DELETE("/payments/:id", PaymentController.DeletePayments, authMiddleware)
	e.PUT("/payments/:id", PaymentController.UpdatePayments, authMiddleware)
	e.POST("/payments", PaymentController.CreatePayment)
}

func (u *PaymentController) CreatePayment(c echo.Context) error {
	var req request.PaymentCreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdPayment, err := u.PaymentUsecase.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.PaymentCreateResponse{
		ID:            int(createdPayment.ID),
		PaymentDate:   createdPayment.PaymentDate,
		PaymentAmount: createdPayment.PaymentAmount,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":   201,
		"status": true,
		"data":   res,
	})
}

func (u *PaymentController) GetPaymentByID(c echo.Context) error {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	foundPayment, err := u.PaymentUsecase.ReadByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.PaymentsResponse{
		ID:            int(foundPayment.ID),
		PaymentDate:   foundPayment.PaymentDate,
		PaymentAmount: foundPayment.PaymentAmount,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *PaymentController) GetPayments(c echo.Context) error {
	foundPayments, err := u.PaymentUsecase.ReadAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var res []response.PaymentsResponse
	for _, foundPayment := range *foundPayments {
		res = append(res, response.PaymentsResponse{
			ID:            int(foundPayment.ID),
			PaymentDate:   foundPayment.PaymentDate,
			PaymentAmount: foundPayment.PaymentAmount,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *PaymentController) DeletePayments(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, e := u.PaymentUsecase.Delete(id)

	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	config.DB.Unscoped().Delete(&domain.Payment{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete payment ticket with id " + strconv.Itoa(id),
		"code":     200,
	})
}

func (u *PaymentController) UpdatePayments(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updatePayment := domain.Payment{}
	err = c.Bind(&updatePayment)
	if err != nil {
		return err
	}

	if err := config.DB.Model(&domain.Payment{}).Where("id = ?", id).Updates(domain.Payment{
		PaymentDate:   updatePayment.PaymentDate,
		PaymentAmount: updatePayment.PaymentAmount,
	}).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
			"code":     400,
		})
	}
	foundPayment, _ := u.PaymentUsecase.ReadByID(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update payment with id " + strconv.Itoa(id),
		"code":     200,
		"data":     foundPayment,
	})
}
