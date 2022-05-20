package repository

import (
	"github.com/San-Kun/MiniProject/domain"
	"gorm.io/gorm"
)

type bankAccountRepository struct {
	Conn *gorm.DB
}

func NewBankAccountRepository(Conn *gorm.DB) domain.BankAccountRepository {
	return &bankAccountRepository{Conn: Conn}
}

func (u *bankAccountRepository) Create(bankAccount *domain.BankAccount) (*domain.BankAccount, error) {
	if err := u.Conn.Create(&bankAccount).Error; err != nil {
		return nil, err
	}

	return bankAccount, nil
}

func (u *bankAccountRepository) ReadByID(id int) (*domain.BankAccount, error) {
	bankAccount := &domain.BankAccount{ID: id}
	if err := u.Conn.First(&bankAccount).Error; err != nil {
		return nil, err
	}

	return bankAccount, nil
}

func (u *bankAccountRepository) ReadAll() (*domain.BankAccounts, error) {
	bankAccounts := &domain.BankAccounts{}
	u.Conn.Find(&bankAccounts)

	return bankAccounts, nil
}

func (u *bankAccountRepository) Delete(id int) (*domain.BankAccount, error) {
	bankAccount := &domain.BankAccount{ID: id}
	if err := u.Conn.Delete(&bankAccount).Error; err != nil {
		return nil, err
	}
	return bankAccount, nil
}

func (u *bankAccountRepository) Updates(id int) (*domain.BankAccount, error) {
	bankAccount := &domain.BankAccount{ID: id}
	if err := u.Conn.Updates(&bankAccount).Error; err != nil {
		return nil, err
	}

	return bankAccount, nil
}
