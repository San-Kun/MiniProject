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

type UserController struct {
	UserUsecase domain.UserUsecase
}

func NewUserController(e *echo.Echo, Usecase domain.UserUsecase) {
	UserController := &UserController{
		UserUsecase: Usecase,
	}

	e.POST("/login", UserController.Login)
	authMiddleware := mid.NewGoMiddleware().AuthMiddleware()
	e.GET("/users/:id", UserController.GetUserByID, authMiddleware)
	e.DELETE("/users/:id", UserController.DeleteUsers, authMiddleware)
	e.PUT("/users/:id", UserController.UpdateUsers, authMiddleware)
	e.GET("/users", UserController.GetUsers, authMiddleware)
	e.POST("/users", UserController.CreateUser)
}

// Login godoc
// @Summary Login User
// @Description Login User
// @Tags Auth
// @accept json
// @Produce json
// @Router /login [post]
// @Param data body response.SuccessLogin true "required"
// @Success 200 {object} domain.User
// @Failure 401 {object} domain.User
func (u *UserController) Login(c echo.Context) error {
	var req request.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res, err := u.UserUsecase.Login(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    401,
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"status":   true,
		"id_user":  res.ID,
		"roles_id": res.RoleId,
		"name":     res.Name,
		"email":    res.Email,
		"no_telp":  res.NoTelp,
		"token":    res.Token,
	})
}

// CreateUser godoc
// @Summary Create User
// @Description Create User
// @Tags User
// @accept json
// @Produce json
// @Router /users [post]
// @param data body response.UserResponse true "required"
// @Success 200 {object} response.UsersResponse
// @Failure 400 {object} response.UsersResponse
func (u *UserController) CreateUser(c echo.Context) error {
	var req request.UserCreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdUser, err := u.UserUsecase.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.UserCreateResponse{
		ID:       int(createdUser.ID),
		Name:     createdUser.Name,
		Email:    createdUser.Email,
		Password: createdUser.Password,
		NoTelp:   createdUser.NoTelp,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

// GetUserById godoc
// @Summary Get User By Id
// @Description Can Get User By Id
// @Tags User
// @accept json
// @Produce json
// @Router /users/{id} [get]
// @param id path int true "id"
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} response.UserResponse
func (u *UserController) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	foundUser, err := u.UserUsecase.ReadByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    400,
			"status":  false,
			"message": err.Error(),
		})
	}

	res := response.UserResponse{
		ID:       int(foundUser.ID),
		Name:     foundUser.Name,
		Email:    foundUser.Email,
		Password: foundUser.Password,
		NoTelp:   foundUser.NoTelp,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

// GetAllUsers godoc
// @Summary Get All User
// @Description Can Get Users
// @Tags User
// @accept json
// @Produce json
// @Router /users [get]
// @Success 200 {object} response.UsersResponse
func (u *UserController) GetUsers(c echo.Context) error {
	foundUsers, err := u.UserUsecase.ReadAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var res []response.UsersResponse
	for _, foundUser := range *foundUsers {
		res = append(res, response.UsersResponse{
			ID:       int(foundUser.ID),
			Name:     foundUser.Name,
			Email:    foundUser.Email,
			Password: foundUser.Password,
			NoTelp:   foundUser.NoTelp,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": true,
		"data":   res,
	})
}

// DeleteUsers godoc
// @Summary Delete User Student
// @Description Can Delete User
// @Tags User
// @accept json
// @Produce json
// @Router /users/{id} [delete]
// @param id path int true "id"
// @Success 200 {object} response.UsersResponse
// @Failure 404 {object} response.UsersResponse
func (u *UserController) DeleteUsers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, e := u.UserUsecase.Delete(id)

	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	config.DB.Unscoped().Delete(&domain.User{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user with id " + strconv.Itoa(id),
		"code":     200,
	})
}

// UpdateUsers godoc
// @Summary Update Users
// @Description Can Update User
// @Tags User
// @accept json
// @Produce json
// @Router /users/{id} [put]
// @param id path int true "id"
// @param data body response.UsersResponse true "required"
// @Success 200 {object} response.UsersResponse
// @Failure 400 {object} response.UsersResponse
func (u *UserController) UpdateUsers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updateUser := domain.User{}
	err = c.Bind(&updateUser)
	if err != nil {
		return err
	}

	if err := config.DB.Model(&domain.User{}).Where("id = ?", id).Updates(domain.User{
		Name:     updateUser.Name,
		Email:    updateUser.Email,
		Password: updateUser.Password,
		NoTelp:   updateUser.NoTelp,
	}).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
			"code":     400,
		})
	}
	foundUser, _ := u.UserUsecase.ReadByID(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user with id " + strconv.Itoa(id),
		"code":     200,
		"data":     foundUser,
	})
}
