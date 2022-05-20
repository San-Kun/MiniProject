package usecase

import (
	"errors"

	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/domain/web/request"
)

type bankAccountUsecase struct {
	BankAccountRepo domain.BankAccountRepository
}

func NewBankAccountUsecase(ur domain.BankAccountRepository) domain.BankAccountUsecase {
	return &bankAccountUsecase{BankAccountRepo: ur}
}

func (u *bankAccountUsecase) Create(request request.BankAccountCreateRequest) (*domain.BankAccount, error) {
	if request.AccountNumber == "" {
		return nil, errors.New("account_number empty")
	}
	bankAccount := &domain.BankAccount{
		AccountNumber: request.AccountNumber,
		AccountName:   request.AccountName,
	}

	createdBankAccount, err := u.BankAccountRepo.Create(bankAccount)
	if err != nil {
		return nil, err
	}

	return createdBankAccount, nil
}

func (u *bankAccountUsecase) ReadByID(id int) (*domain.BankAccount, error) {
	bankAccount, err := u.BankAccountRepo.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return bankAccount, err
}

func (u *bankAccountUsecase) ReadAll() (*domain.BankAccounts, error) {
	foundBankAccounts, err := u.BankAccountRepo.ReadAll()
	if err != nil {
		return nil, err
	}

	return foundBankAccounts, nil
}

func (u *bankAccountUsecase) Delete(id int) (*domain.BankAccount, error) {
	bankAccount, err := u.BankAccountRepo.Delete(id)
	if err != nil {
		return nil, err
	}

	return bankAccount, err
}

func (u *bankAccountUsecase) Updates(id int) (*domain.BankAccount, error) {
	bankAccount, err := u.BankAccountRepo.Updates(id)
	if err != nil {
		return nil, err
	}

	return bankAccount, err
}
