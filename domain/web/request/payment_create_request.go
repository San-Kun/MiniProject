package request

type PaymentCreateRequest struct {
	ID            int       `json:"id"`
	EventId       EventId   `json:"event_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PaymentDate   string    `json:"payment_date" validate:"required"`
	PaymentAmount int       `json:"payment_amount" validate:"required"`
	PaytypeId     PaytypeId `json:"paytype_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type EventId struct {
	ID           int    `json:"id"`
	EventName    string `json:"event_name" validate:"required"`
	EventTanggal string `json:"event_tanggal" validate:"required"`
	Capacity     string `json:"capacity" validate:"required"`
}

type PaytypeId struct {
	ID          int    `json:"id" validate:"required"`
	PaytypeName string `json:"paytype_name" validate:"required"`
}
