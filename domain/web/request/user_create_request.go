package request

type UserCreateRequest struct {
	ID       int    `json:"id"`
	RoleId   int    `json:"roles_id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	NoTelp   string `json:"no_telp" validate:"required"`
	Roles    Roles  `json:"roles" gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Roles struct {
	ID       int    `json:"id"`
	RoleName string `json:"role_name" validate:"required"`
}
