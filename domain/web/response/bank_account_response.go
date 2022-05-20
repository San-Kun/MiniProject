package response

type BankAccountResponse struct {
	ID            int    `json:"id"`
	UserID        int    `json:"userid" validate:"required"`
	BankID        int    `json:"bankid" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
	AccountName   string `json:"account_name" validate:"required"`
	User          User   `json:"user_id" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Bank          Bank   `json:"bank_id" gorm:"foreignKey:BankID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type BankAccountsResponse struct {
	ID            int    `json:"id"`
	UserID        int    `json:"userid" validate:"required"`
	BankID        int    `json:"bankid" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
	AccountName   string `json:"account_name" validate:"required"`
	User          User   `json:"user_id" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Bank          Bank   `json:"bank_id" gorm:"foreignKey:BankID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
