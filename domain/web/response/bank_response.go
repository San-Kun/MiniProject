package response

type BankResponse struct {
	ID       int    `json:"id" validate:"required"`
	BankName string `json:"bank_name" validate:"required"`
}

type BanksResponse struct {
	ID       int    `json:"id" validate:"required"`
	BankName string `json:"bank_name" validate:"required"`
}
