package response

import "gorm.io/gorm"

type BankAccountCreateResponse struct {
	ID            int    `json:"id"`
	UserID        int    `json:"userid" validate:"required"`
	BankID        int    `json:"bankid" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
	AccountName   string `json:"account_name" validate:"required"`
	User          User   `json:"user_id" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Bank          Bank   `json:"bank_id" gorm:"foreignKey:BankID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type User struct {
	gorm.Model
	ID     int    `json:"id"`
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required"`
	NoTelp string `json:"no_telp" validate:"required"`
}

type Bank struct {
	gorm.Model
	ID       int    `json:"id"`
	BankName string `json:"bank_name" validate:"required"`
}
