package repository

import (
	"github.com/San-Kun/MiniProject/domain"
	"gorm.io/gorm"
)

type paymentTypeRepository struct {
	Conn *gorm.DB
}

func NewPaymentTypeRepository(Conn *gorm.DB) domain.PaymentTypeRepository {
	return &paymentTypeRepository{Conn: Conn}
}

func (u *paymentTypeRepository) Create(paymentType *domain.PaymentType) (*domain.PaymentType, error) {
	if err := u.Conn.Create(&paymentType).Error; err != nil {
		return nil, err
	}

	return paymentType, nil
}

func (u *paymentTypeRepository) ReadByID(id int) (*domain.PaymentType, error) {
	paymentType := &domain.PaymentType{ID: id}
	if err := u.Conn.First(&paymentType).Error; err != nil {
		return nil, err
	}

	return paymentType, nil
}

func (u *paymentTypeRepository) ReadAll() (*domain.PaymentTypes, error) {
	paymentTypes := &domain.PaymentTypes{}
	u.Conn.Find(&paymentTypes)

	return paymentTypes, nil
}

func (u *paymentTypeRepository) Delete(id int) (*domain.PaymentType, error) {
	paymentType := &domain.PaymentType{ID: id}
	if err := u.Conn.Delete(&paymentType).Error; err != nil {
		return nil, err
	}
	return paymentType, nil
}

func (u *paymentTypeRepository) Updates(id int) (*domain.PaymentType, error) {
	paymentType := &domain.PaymentType{ID: id}
	if err := u.Conn.Updates(&paymentType).Error; err != nil {
		return nil, err
	}

	return paymentType, nil
}
