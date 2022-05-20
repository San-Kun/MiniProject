package response

type EventTicketCreateResponse struct {
	ID          int     `json:"id"`
	EventID     int     `json:"event_id" validate:"required"`
	PaymentID   int     `json:"payment_id" validate:"required"`
	PayTypeID   int     `json:"pay_type_id" validate:"required"`
	PayTypeName string  `json:"paytype_name" validate:"required"`
	Event       Event   `json:"events" gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Payment     Payment `json:"payments" gorm:"foreignKey:PaymentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
