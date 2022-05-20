package domain

import (
	"github.com/San-Kun/MiniProject/domain/web/request"
)

type BankAccount struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	BankID        int    `json:"bank_id"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
}

type BankAccounts []BankAccount

type BankAccountRepository interface {
	Create(bankAccount *BankAccount) (*BankAccount, error)
	ReadByID(id int) (*BankAccount, error)
	Delete(id int) (*BankAccount, error)
	Updates(id int) (*BankAccount, error)
	ReadAll() (*BankAccounts, error)
}

type BankAccountUsecase interface {
	Create(request request.BankAccountCreateRequest) (*BankAccount, error)
	ReadByID(id int) (*BankAccount, error)
	Delete(id int) (*BankAccount, error)
	Updates(id int) (*BankAccount, error)
	ReadAll() (*BankAccounts, error)
}
