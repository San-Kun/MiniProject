package response

type UserResponse struct {
	ID       int    `json:"id"`
	RoleId   int    `json:"roles_id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	NoTelp   string `json:"no_telp" validate:"required"`
	Roles    Roles  `json:"roles" gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UsersResponse struct {
	ID       int    `json:"id"`
	RoleId   int    `json:"roles_id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	NoTelp   string `json:"no_telp" validate:"required"`
	Roles    Roles  `json:"roles" gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
