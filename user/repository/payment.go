package repository

import (
	"github.com/San-Kun/MiniProject/domain"
	"gorm.io/gorm"
)

type paymentRepository struct {
	Conn *gorm.DB
}

func NewPaymentRepository(Conn *gorm.DB) domain.PaymentRepository {
	return &paymentRepository{Conn: Conn}
}

func (u *paymentRepository) Create(payment *domain.Payment) (*domain.Payment, error) {
	if err := u.Conn.Create(&payment).Error; err != nil {
		return nil, err
	}

	return payment, nil
}

func (u *paymentRepository) ReadByID(id int) (*domain.Payment, error) {
	payment := &domain.Payment{ID: id}
	if err := u.Conn.First(&payment).Error; err != nil {
		return nil, err
	}

	return payment, nil
}

func (u *paymentRepository) ReadAll() (*domain.Payments, error) {
	payments := &domain.Payments{}
	u.Conn.Find(&payments)

	return payments, nil
}

func (u *paymentRepository) Delete(id int) (*domain.Payment, error) {
	payment := &domain.Payment{ID: id}
	if err := u.Conn.Delete(&payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (u *paymentRepository) Updates(id int) (*domain.Payment, error) {
	payment := &domain.Payment{ID: id}
	if err := u.Conn.Updates(&payment).Error; err != nil {
		return nil, err
	}

	return payment, nil
}
