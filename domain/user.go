package domain

import (
	"github.com/San-Kun/MiniProject/domain/web/request"
	"github.com/San-Kun/MiniProject/domain/web/response"
)

type User struct {
	ID       int    `json:"id"`
	RoleId   int    `json:"roles_id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	NoTelp   string `json:"no_telp" validate:"required"`
	Roles    Roles  `json:"roles" gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Roles struct {
	ID          int    `json:"id"`
	RoleName    string `json:"role_name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type Users []User

type UserRepository interface {
	Create(user *User) (*User, error)
	ReadByID(id int) (*User, error)
	Delete(id int) (*User, error)
	Updates(id int) (*User, error)
	ReadAll() (*Users, error)
	CheckLogin(user *User) (*User, bool, error)
}

type UserUsecase interface {
	Create(request request.UserCreateRequest) (*User, error)
	ReadByID(id int) (*User, error)
	ReadAll() (*Users, error)
	Delete(id int) (*User, error)
	Updates(id int) (*User, error)
	Login(request request.LoginRequest) (*response.SuccessLogin, error)
}
