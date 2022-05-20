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

type BankAccountController struct {
	BankAccountUsecase domain.BankAccountUsecase
}

func NewBankAccountController(e *echo.Echo, Usecase domain.BankAccountUsecase) {
	BankAccountController := &BankAccountController{
		BankAccountUsecase: Usecase,
	}

	authMiddleware := mid.NewGoMiddleware().AuthMiddleware()
	e.GET("/bankAccounts/:id", BankAccountController.GetBankAccountByID, authMiddleware)
	e.GET("/bankAccounts", BankAccountController.GetBankAccounts, authMiddleware)
	e.DELETE("/bankAccounts/:id", BankAccountController.DeleteBankAccounts, authMiddleware)
	e.PUT("/bankAccounts/:id", BankAccountController.UpdateBankAccounts, authMiddleware)
	e.POST("/bankAccounts", BankAccountController.CreateBankAccount)
}

func (u *BankAccountController) CreateBankAccount(c echo.Context) error {
	var req request.BankAccountCreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdBankAccount, err := u.BankAccountUsecase.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.BankAccountCreateResponse{
		ID:            int(createdBankAccount.ID),
		UserID:        int(createdBankAccount.UserID),
		BankID:        int(createdBankAccount.BankID),
		AccountNumber: createdBankAccount.AccountNumber,
		AccountName:   createdBankAccount.AccountName,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":   201,
		"status": true,
		"data":   res,
	})
}

func (u *BankAccountController) GetBankAccountByID(c echo.Context) error {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	foundBankAccount, err := u.BankAccountUsecase.ReadByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.BankAccountsResponse{
		ID:            int(foundBankAccount.ID),
		UserID:        int(foundBankAccount.UserID),
		BankID:        int(foundBankAccount.BankID),
		AccountNumber: foundBankAccount.AccountNumber,
		AccountName:   foundBankAccount.AccountName,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *BankAccountController) GetBankAccounts(c echo.Context) error {
	foundBankAccounts, err := u.BankAccountUsecase.ReadAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var res []response.BankAccountsResponse
	for _, foundBankAccount := range *foundBankAccounts {
		res = append(res, response.BankAccountsResponse{
			ID:            int(foundBankAccount.ID),
			UserID:        int(foundBankAccount.UserID),
			BankID:        int(foundBankAccount.BankID),
			AccountNumber: foundBankAccount.AccountNumber,
			AccountName:   foundBankAccount.AccountName,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

func (u *BankAccountController) DeleteBankAccounts(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, e := u.BankAccountUsecase.Delete(id)

	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	config.DB.Unscoped().Delete(&domain.BankAccount{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete bank account with id " + strconv.Itoa(id),
		"code":     200,
	})
}

func (u *BankAccountController) UpdateBankAccounts(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updateBankAccount := domain.BankAccount{}
	err = c.Bind(&updateBankAccount)
	if err != nil {
		return err
	}

	if err := config.DB.Model(&domain.BankAccount{}).Where("id = ?", id).Updates(domain.BankAccount{
		BankID:        updateBankAccount.BankID,
		AccountNumber: updateBankAccount.AccountNumber,
		AccountName:   updateBankAccount.AccountName,
	}).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
			"code":     400,
		})
	}
	foundBankAccount, _ := u.BankAccountUsecase.ReadByID(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update bankAccount with id " + strconv.Itoa(id),
		"code":     200,
		"data":     foundBankAccount,
	})
}
