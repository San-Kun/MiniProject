package response

import "gorm.io/gorm"

type SuccessLogin struct {
	ID     int    `json:"id" form:"id"`
	RoleId int    `json:"rolesid"`
	Name   string `json:"name" form:"name" validate:"required"`
	Email  string `json:"email" form:"email" validate:"required"`
	NoTelp string `json:"notelp" form:"notelp" validate:"required"`
	Token  string `json:"token" form:"token" validate:"required"`
	Roles  Roles  `json:"roles_id" gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Roles struct {
	gorm.Model
	ID          int    `json:"id" form:"id"`
	RoleName    string `json:"role_name" form:"role_name" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
}
