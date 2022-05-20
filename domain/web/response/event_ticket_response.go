package response

import "gorm.io/gorm"

type EventTicketResponse struct {
	ID          int     `json:"id"`
	EventID     int     `json:"event_id"`
	PaymentID   int     `json:"payment_id" validate:"required"`
	PayTypeID   int     `json:"pay_type_id" validate:"required"`
	PayTypeName string  `json:"paytype_name" validate:"required"`
	Event       Event   `json:"events" gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Payment     Payment `json:"payments" gorm:"foreignKey:PaymentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type EventTicketsResponse struct {
	ID          int     `json:"id"`
	EventID     int     `json:"event_id"`
	PaymentID   int     `json:"payment_id" validate:"required"`
	PayTypeID   int     `json:"pay_type_id" validate:"required"`
	PayTypeName string  `json:"paytype_name" validate:"required"`
	Event       Event   `json:"events" gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Payment     Payment `json:"payments" gorm:"foreignKey:PaymentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Event struct {
	gorm.Model
	ID           int    `json:"id"`
	EventName    string `json:"event_name"`
	EventTanggal string `json:"event_tanggal"`
	Capacity     string `json:"capacity"`
}

type Payment struct {
	gorm.Model
	ID            int    `json:"id"`
	PaymentDate   string `json:"payment_date"`
	PaymentAmount int    `json:"payment_amount"`
}
