package usecase_test

import (
	"errors"
	"testing"

	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/domain/mocks"
	"github.com/San-Kun/MiniProject/domain/web/request"
	usecase "github.com/San-Kun/MiniProject/user/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBankAccountUsecase_ReadAll(t *testing.T) {
	mockBankAccountRepo := new(mocks.BankAccountRepository)
	mockListBankAccount := &domain.BankAccounts{
		domain.BankAccount{
			ID:            1,
			UserID:        1,
			BankID:        1,
			AccountNumber: "987654321",
			AccountName:   "IKHSAN ENDANG PRASETYA",
		},
		domain.BankAccount{
			ID:            2,
			UserID:        2,
			BankID:        2,
			AccountNumber: "1234567890",
			AccountName:   "ENDANG PRASETYA",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockBankAccountRepo.On("ReadAll").Return(mockListBankAccount, nil).Once()
		uc := usecase.NewBankAccountUsecase(mockBankAccountRepo)
		res, err := uc.ReadAll()
		assert.NoError(t, err)
		assert.Len(t, *res, len(*mockListBankAccount))
		mockBankAccountRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockBankAccountRepo.On("ReadAll").Return(nil, errors.New("error something")).Once()
		uc := usecase.NewBankAccountUsecase(mockBankAccountRepo)
		_, err := uc.ReadAll()
		assert.Error(t, err)
		mockBankAccountRepo.AssertExpectations(t)
	})
}

func TestBankAccountUsecase_ReadByID(t *testing.T) {
	mockBankAccountRepo := new(mocks.BankAccountRepository)
	mockBankAccount := &domain.BankAccount{
		ID:            1,
		UserID:        1,
		BankID:        1,
		AccountNumber: "987654321",
		AccountName:   "IKHSAN ENDANG PRASETYA",
	}

	t.Run("success", func(t *testing.T) {
		mockBankAccountRepo.On("ReadByID", mock.AnythingOfType("int")).Return(mockBankAccount, nil).Once()
		uc := usecase.NewBankAccountUsecase(mockBankAccountRepo)
		res, err := uc.ReadByID(1)
		assert.NoError(t, err)
		assert.Equal(t, res, mockBankAccount)
		mockBankAccountRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockBankAccountRepo.On("ReadByID", mock.AnythingOfType("int")).Return(nil, errors.New("error something")).Once()
		uc := usecase.NewBankAccountUsecase(mockBankAccountRepo)
		_, err := uc.ReadByID(2)
		assert.Error(t, err)
		mockBankAccountRepo.AssertExpectations(t)
	})
}

func TestBankAccountUsecase_Create(t *testing.T) {
	mockBankAccountRepo := new(mocks.BankAccountRepository)
	mockBankAccount := &domain.BankAccount{
		ID:            1,
		UserID:        1,
		BankID:        1,
		AccountNumber: "987654321",
		AccountName:   "IKHSAN ENDANG PRASETYA",
	}

	t.Run("success", func(t *testing.T) {
		mockRequest := &request.BankAccountCreateRequest{
			ID:            1,
			UserID:        1,
			BankID:        1,
			AccountNumber: "987654321",
			AccountName:   "IKHSAN ENDANG PRASETYA",
		}

		mockBankAccountRepo.On("Create", mock.AnythingOfType("*domain.BankAccount")).Return(mockBankAccount, nil).Once()
		uc := usecase.NewBankAccountUsecase(mockBankAccountRepo)
		res, err := uc.Create(*mockRequest)
		assert.NoError(t, err)
		assert.Equal(t, res, mockBankAccount)
		mockBankAccountRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockRequest := &request.BankAccountCreateRequest{
			ID:            1,
			UserID:        1,
			BankID:        1,
			AccountNumber: "987654321",
			AccountName:   "IKHSAN ENDANG PRASETYA",
		}

		mockBankAccountRepo.On("Create", mock.AnythingOfType("*domain.BankAccount")).Return(nil, errors.New("error something")).Once()
		uc := usecase.NewBankAccountUsecase(mockBankAccountRepo)
		_, err := uc.Create(*mockRequest)
		assert.Error(t, err)
		mockBankAccountRepo.AssertExpectations(t)
	})

	t.Run("bankAccount-empty-failed", func(t *testing.T) {
		mockRequest := &request.BankAccountCreateRequest{
			ID:            1,
			UserID:        1,
			BankID:        1,
			AccountNumber: "987654321",
			AccountName:   "IKHSAN ENDANG PRASETYA",
		}

		mockBankAccountRepo.On("Create", mock.AnythingOfType("*domain.BankAccount")).Return(nil, errors.New("BankAccount empty")).Once()
		uc := usecase.NewBankAccountUsecase(mockBankAccountRepo)
		_, err := uc.Create(*mockRequest)
		assert.Error(t, err)
		mockBankAccountRepo.AssertExpectations(t)
	})
}
