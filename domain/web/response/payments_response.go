package response

type PaymentsResponse struct {
	ID            int       `json:"id"`
	EventID       EventID   `json:"event_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PaymentDate   string    `json:"payment_date" validate:"required"`
	PaymentAmount int       `json:"payment_amount" validate:"required"`
	PaytypeId     PaytypeId `json:"paytype_id"`
}
