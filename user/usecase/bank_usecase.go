package usecase

import (
	"errors"

	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/domain/web/request"
)

type bankUsecase struct {
	BankRepo domain.BankRepository
}

func NewBankUsecase(ur domain.BankRepository) domain.BankUsecase {
	return &bankUsecase{BankRepo: ur}
}

func (u *bankUsecase) Create(request request.BankCreateRequest) (*domain.Bank, error) {
	if request.BankName == "" {
		return nil, errors.New("bank empty")
	}
	bank := &domain.Bank{
		BankName: request.BankName,
	}

	createdBank, err := u.BankRepo.Create(bank)
	if err != nil {
		return nil, err
	}

	return createdBank, nil
}

func (u *bankUsecase) ReadByID(id int) (*domain.Bank, error) {
	bank, err := u.BankRepo.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return bank, err
}

func (u *bankUsecase) ReadAll() (*domain.Banks, error) {
	foundBanks, err := u.BankRepo.ReadAll()
	if err != nil {
		return nil, err
	}

	return foundBanks, nil
}

func (u *bankUsecase) Delete(id int) (*domain.Bank, error) {
	bank, err := u.BankRepo.Delete(id)
	if err != nil {
		return nil, err
	}

	return bank, err
}

func (u *bankUsecase) Updates(id int) (*domain.Bank, error) {
	bank, err := u.BankRepo.Updates(id)
	if err != nil {
		return nil, err
	}

	return bank, err
}
