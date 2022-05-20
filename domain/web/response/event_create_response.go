package response

type EventCreateResponse struct {
	ID           int    `json:"id"`
	UserID       UserID `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EventName    string `json:"event_name" validate:"required"`
	EventTanggal string `json:"event_tanggal" validate:"required"`
	Capacity     string `json:"capacity" validate:"required"`
}

type UserID struct {
	ID       int    `json:"id"`
	RoleId   int    `json:"roles" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	NoTelp   string `json:"no_telp" validate:"required"`
}
