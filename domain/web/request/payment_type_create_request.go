package request

type PaymentTypeCreateRequest struct {
	ID          int    `json:"id"`
	PayTypeName string `json:"pay_type_name" validate:"required"`
}
