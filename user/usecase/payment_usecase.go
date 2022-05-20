package usecase

import (
	"errors"

	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/domain/web/request"
)

type paymentUsecase struct {
	PaymentRepo domain.PaymentRepository
}

func NewPaymentUsecase(ur domain.PaymentRepository) domain.PaymentUsecase {
	return &paymentUsecase{PaymentRepo: ur}
}

func (u *paymentUsecase) Create(request request.PaymentCreateRequest) (*domain.Payment, error) {
	if request.PaymentDate == "" {
		return nil, errors.New("email empty")
	}
	payment := &domain.Payment{
		PaymentDate:   request.PaymentDate,
		PaymentAmount: request.PaymentAmount,
	}

	createdPayment, err := u.PaymentRepo.Create(payment)
	if err != nil {
		return nil, err
	}

	return createdPayment, nil
}

func (u *paymentUsecase) ReadByID(id int) (*domain.Payment, error) {
	payment, err := u.PaymentRepo.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return payment, err
}

func (u *paymentUsecase) ReadAll() (*domain.Payments, error) {
	foundPayments, err := u.PaymentRepo.ReadAll()
	if err != nil {
		return nil, err
	}

	return foundPayments, nil
}

func (u *paymentUsecase) Delete(id int) (*domain.Payment, error) {
	payment, err := u.PaymentRepo.Delete(id)
	if err != nil {
		return nil, err
	}

	return payment, err
}

func (u *paymentUsecase) Updates(id int) (*domain.Payment, error) {
	payment, err := u.PaymentRepo.Updates(id)
	if err != nil {
		return nil, err
	}

	return payment, err
}
