package domain

import (
	"github.com/San-Kun/MiniProject/domain/web/request"
)

type Bank struct {
	ID       int    `json:"id"`
	BankName string `json:"bank_name"`
}

type Banks []Bank

type BankRepository interface {
	Create(bank *Bank) (*Bank, error)
	ReadByID(id int) (*Bank, error)
	Delete(id int) (*Bank, error)
	Updates(id int) (*Bank, error)
	ReadAll() (*Banks, error)
}

type BankUsecase interface {
	Create(request request.BankCreateRequest) (*Bank, error)
	ReadByID(id int) (*Bank, error)
	Delete(id int) (*Bank, error)
	Updates(id int) (*Bank, error)
	ReadAll() (*Banks, error)
}
