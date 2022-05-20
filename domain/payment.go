package domain

import (
	"github.com/San-Kun/MiniProject/domain/web/request"
)

type Payment struct {
	ID            int    `json:"id"`
	PaymentDate   string `json:"payment_date"`
	PaymentAmount int    `json:"payment_amount"`
}

type Payments []Payment

type PaymentRepository interface {
	Create(payment *Payment) (*Payment, error)
	ReadByID(id int) (*Payment, error)
	Delete(id int) (*Payment, error)
	Updates(id int) (*Payment, error)
	ReadAll() (*Payments, error)
}

type PaymentUsecase interface {
	Create(request request.PaymentCreateRequest) (*Payment, error)
	ReadByID(id int) (*Payment, error)
	Delete(id int) (*Payment, error)
	Updates(id int) (*Payment, error)
	ReadAll() (*Payments, error)
}
