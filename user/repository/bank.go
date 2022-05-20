package repository

import (
	"github.com/San-Kun/MiniProject/domain"
	"gorm.io/gorm"
)

type bankRepository struct {
	Conn *gorm.DB
}

func NewBankRepository(Conn *gorm.DB) domain.BankRepository {
	return &bankRepository{Conn: Conn}
}

func (u *bankRepository) Create(bank *domain.Bank) (*domain.Bank, error) {
	if err := u.Conn.Create(&bank).Error; err != nil {
		return nil, err
	}

	return bank, nil
}

func (u *bankRepository) ReadByID(id int) (*domain.Bank, error) {
	bank := &domain.Bank{ID: id}
	if err := u.Conn.First(&bank).Error; err != nil {
		return nil, err
	}

	return bank, nil
}

func (u *bankRepository) ReadAll() (*domain.Banks, error) {
	banks := &domain.Banks{}
	u.Conn.Find(&banks)

	return banks, nil
}

func (u *bankRepository) Delete(id int) (*domain.Bank, error) {
	bank := &domain.Bank{ID: id}
	if err := u.Conn.Delete(&bank).Error; err != nil {
		return nil, err
	}
	return bank, nil
}

func (u *bankRepository) Updates(id int) (*domain.Bank, error) {
	bank := &domain.Bank{ID: id}
	if err := u.Conn.Updates(&bank).Error; err != nil {
		return nil, err
	}

	return bank, nil
}
