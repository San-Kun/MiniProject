package request

import "gorm.io/gorm"

type BankAccountCreateRequest struct {
	ID            int    `json:"id"`
	UserID        int    `json:"users_id" validate:"required"`
	BankID        int    `json:"banks_id" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
	AccountName   string `json:"account_name" validate:"required"`
	User          User   `json:"users" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Bank          Bank   `json:"banks" gorm:"foreignKey:BankID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
