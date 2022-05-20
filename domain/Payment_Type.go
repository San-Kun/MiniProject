package domain

import (
	"github.com/San-Kun/MiniProject/domain/web/request"
)

type PaymentType struct {
	ID          int    `json:"id"`
	PayTypeName string `json:"pay_type_name"`
}

type PaymentTypes []PaymentType

type PaymentTypeRepository interface {
	Create(paymentType *PaymentType) (*PaymentType, error)
	ReadByID(id int) (*PaymentType, error)
	Delete(id int) (*PaymentType, error)
	Updates(id int) (*PaymentType, error)
	ReadAll() (*PaymentTypes, error)
}

type PaymentTypeUsecase interface {
	Create(request request.PaymentTypeCreateRequest) (*PaymentType, error)
	ReadByID(id int) (*PaymentType, error)
	Delete(id int) (*PaymentType, error)
	Updates(id int) (*PaymentType, error)
	ReadAll() (*PaymentTypes, error)
}
