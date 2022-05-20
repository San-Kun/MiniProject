package response

type BankCreateResponse struct {
	ID       int    `json:"id" validate:"required"`
	BankName string `json:"bank_name" validate:"required"`
}
