package response

type PaymentTypeResponse struct {
	ID          int    `json:"id"`
	PayTypeName string `json:"pay_type_name" validate:"required"`
}

type PaymentTypesResponse struct {
	ID          int    `json:"id"`
	PayTypeName string `json:"pay_type_name" validate:"required"`
}
