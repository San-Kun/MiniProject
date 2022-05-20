package usecase

import (
	"errors"

	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/domain/web/request"
)

type paymentTypeUsecase struct {
	PaymentTypeRepo domain.PaymentTypeRepository
}

func NewPaymentTypeUsecase(ur domain.PaymentTypeRepository) domain.PaymentTypeUsecase {
	return &paymentTypeUsecase{PaymentTypeRepo: ur}
}

func (u *paymentTypeUsecase) Create(request request.PaymentTypeCreateRequest) (*domain.PaymentType, error) {
	if request.PayTypeName == "" {
		return nil, errors.New("PaymentType empty")
	}
	paymentType := &domain.PaymentType{
		PayTypeName: request.PayTypeName,
	}

	createdPaymentType, err := u.PaymentTypeRepo.Create(paymentType)
	if err != nil {
		return nil, err
	}

	return createdPaymentType, nil
}

func (u *paymentTypeUsecase) ReadByID(id int) (*domain.PaymentType, error) {
	paymentType, err := u.PaymentTypeRepo.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return paymentType, err
}

func (u *paymentTypeUsecase) ReadAll() (*domain.PaymentTypes, error) {
	foundPaymentTypes, err := u.PaymentTypeRepo.ReadAll()
	if err != nil {
		return nil, err
	}

	return foundPaymentTypes, nil
}

func (u *paymentTypeUsecase) Delete(id int) (*domain.PaymentType, error) {
	paymentType, err := u.PaymentTypeRepo.Delete(id)
	if err != nil {
		return nil, err
	}

	return paymentType, err
}

func (u *paymentTypeUsecase) Updates(id int) (*domain.PaymentType, error) {
	paymentType, err := u.PaymentTypeRepo.Updates(id)
	if err != nil {
		return nil, err
	}

	return paymentType, err
}
