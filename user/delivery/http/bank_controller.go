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

type BankController struct {
	BankUsecase domain.BankUsecase
}

func NewBankController(e *echo.Echo, Usecase domain.BankUsecase) {
	BankController := &BankController{
		BankUsecase: Usecase,
	}

	authMiddleware := mid.NewGoMiddleware().AuthMiddleware()
	e.GET("/banks/:id", BankController.GetBankByID, authMiddleware)
	e.GET("/banks", BankController.GetBanks, authMiddleware)
	e.DELETE("/banks/:id", BankController.DeleteBanks, authMiddleware)
	e.PUT("/banks/:id", BankController.UpdateBanks, authMiddleware)
	e.POST("/banks", BankController.CreateBank)
}

func (u *BankController) CreateBank(c echo.Context) error {
	var req request.BankCreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdBank, err := u.BankUsecase.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.BankCreateResponse{
		ID:       int(createdBank.ID),
		BankName: createdBank.BankName,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":   201,
		"status": true,
		"data":   res,
	})
}

func (u *BankController) GetBankByID(c echo.Context) error {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	foundBank, err := u.BankUsecase.ReadByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.BanksResponse{
		ID:       int(foundBank.ID),
		BankName: foundBank.BankName,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *BankController) GetBanks(c echo.Context) error {
	foundBanks, err := u.BankUsecase.ReadAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var res []response.BanksResponse
	for _, foundBank := range *foundBanks {
		res = append(res, response.BanksResponse{
			ID:       int(foundBank.ID),
			BankName: foundBank.BankName,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *BankController) DeleteBanks(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, e := u.BankUsecase.Delete(id)

	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	config.DB.Unscoped().Delete(&domain.Bank{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete bank with id " + strconv.Itoa(id),
		"code":     200,
	})
}

func (u *BankController) UpdateBanks(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updateBank := domain.Bank{}
	err = c.Bind(&updateBank)
	if err != nil {
		return err
	}

	if err := config.DB.Model(&domain.Bank{}).Where("id = ?", id).Updates(domain.Bank{
		BankName: updateBank.BankName,
	}).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
			"code":     400,
		})
	}
	foundBank, _ := u.BankUsecase.ReadByID(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update bank with id " + strconv.Itoa(id),
		"code":     200,
		"data":     foundBank,
	})
}
