package response

type PaymentTypeCreateResponse struct {
	ID          int    `json:"id"`
	PayTypeName string `json:"pay_type_name" validate:"required"`
}
