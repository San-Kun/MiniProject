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

type PaymentTypeController struct {
	PaymentTypeUsecase domain.PaymentTypeUsecase
}

func NewPaymentTypeController(e *echo.Echo, Usecase domain.PaymentTypeUsecase) {
	PaymentTypeController := &PaymentTypeController{
		PaymentTypeUsecase: Usecase,
	}

	authMiddleware := mid.NewGoMiddleware().AuthMiddleware()
	e.GET("/paymentTypes/:id", PaymentTypeController.GetPaymentTypeByID, authMiddleware)
	e.GET("/paymentTypes", PaymentTypeController.GetPaymentTypes, authMiddleware)
	e.DELETE("/paymentTypes/:id", PaymentTypeController.DeletePaymentTypes, authMiddleware)
	e.PUT("/paymentTypes/:id", PaymentTypeController.UpdatePaymentTypes, authMiddleware)
	e.POST("/paymentTypes", PaymentTypeController.CreatePaymentType)
}

func (u *PaymentTypeController) CreatePaymentType(c echo.Context) error {
	var req request.PaymentTypeCreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdPaymentType, err := u.PaymentTypeUsecase.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.PaymentTypeCreateResponse{
		ID:          int(createdPaymentType.ID),
		PayTypeName: createdPaymentType.PayTypeName,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":   201,
		"status": true,
		"data":   res,
	})
}

func (u *PaymentTypeController) GetPaymentTypeByID(c echo.Context) error {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	foundPaymentType, err := u.PaymentTypeUsecase.ReadByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.PaymentTypesResponse{
		ID:          int(foundPaymentType.ID),
		PayTypeName: foundPaymentType.PayTypeName,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *PaymentTypeController) GetPaymentTypes(c echo.Context) error {
	foundPaymentTypes, err := u.PaymentTypeUsecase.ReadAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var res []response.PaymentTypesResponse
	for _, foundPaymentType := range *foundPaymentTypes {
		res = append(res, response.PaymentTypesResponse{
			ID:          int(foundPaymentType.ID),
			PayTypeName: foundPaymentType.PayTypeName,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *PaymentTypeController) DeletePaymentTypes(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, e := u.PaymentTypeUsecase.Delete(id)

	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	config.DB.Unscoped().Delete(&domain.PaymentType{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete paymentType with id " + strconv.Itoa(id),
		"code":     200,
	})
}

func (u *PaymentTypeController) UpdatePaymentTypes(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updatePaymentType := domain.PaymentType{}
	err = c.Bind(&updatePaymentType)
	if err != nil {
		return err
	}

	if err := config.DB.Model(&domain.PaymentType{}).Where("id = ?", id).Updates(domain.PaymentType{
		PayTypeName: updatePaymentType.PayTypeName,
	}).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
			"code":     400,
		})
	}
	foundPaymentType, _ := u.PaymentTypeUsecase.ReadByID(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update payment type with id " + strconv.Itoa(id),
		"code":     200,
		"data":     foundPaymentType,
	})
}
