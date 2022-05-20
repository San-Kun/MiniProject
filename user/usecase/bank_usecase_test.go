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

func TestBankUsecase_ReadAll(t *testing.T) {
	mockBankRepo := new(mocks.BankRepository)
	mockListBank := &domain.Banks{
		domain.Bank{
			ID:       1,
			BankName: "BANK BRI",
		},
		domain.Bank{
			ID:       2,
			BankName: "BANK BCA",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockBankRepo.On("ReadAll").Return(mockListBank, nil).Once()
		uc := usecase.NewBankUsecase(mockBankRepo)
		res, err := uc.ReadAll()
		assert.NoError(t, err)
		assert.Len(t, *res, len(*mockListBank))
		mockBankRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockBankRepo.On("ReadAll").Return(nil, errors.New("error something")).Once()
		uc := usecase.NewBankUsecase(mockBankRepo)
		_, err := uc.ReadAll()
		assert.Error(t, err)
		mockBankRepo.AssertExpectations(t)
	})
}

func TestBankUsecase_ReadByID(t *testing.T) {
	mockBankRepo := new(mocks.BankRepository)
	mockBank := &domain.Bank{
		ID:       1,
		BankName: "BANK BRI",
	}

	t.Run("success", func(t *testing.T) {
		mockBankRepo.On("ReadByID", mock.AnythingOfType("int")).Return(mockBank, nil).Once()
		uc := usecase.NewBankUsecase(mockBankRepo)
		res, err := uc.ReadByID(1)
		assert.NoError(t, err)
		assert.Equal(t, res, mockBank)
		mockBankRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockBankRepo.On("ReadByID", mock.AnythingOfType("int")).Return(nil, errors.New("error something")).Once()
		uc := usecase.NewBankUsecase(mockBankRepo)
		_, err := uc.ReadByID(2)
		assert.Error(t, err)
		mockBankRepo.AssertExpectations(t)
	})
}

func TestBankUsecase_Create(t *testing.T) {
	mockBankRepo := new(mocks.BankRepository)
	mockBank := &domain.Bank{
		ID:       1,
		BankName: "BANK BRI",
	}

	t.Run("success", func(t *testing.T) {
		mockRequest := &request.BankCreateRequest{
			BankName: "BANK BRI",
		}

		mockBankRepo.On("Create", mock.AnythingOfType("*domain.Bank")).Return(mockBank, nil).Once()
		uc := usecase.NewBankUsecase(mockBankRepo)
		res, err := uc.Create(*mockRequest)
		assert.NoError(t, err)
		assert.Equal(t, res, mockBank)
		mockBankRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockRequest := &request.BankCreateRequest{
			BankName: "BANK BRI",
		}

		mockBankRepo.On("Create", mock.AnythingOfType("*domain.Bank")).Return(nil, errors.New("error something")).Once()
		uc := usecase.NewBankUsecase(mockBankRepo)
		_, err := uc.Create(*mockRequest)
		assert.Error(t, err)
		mockBankRepo.AssertExpectations(t)
	})

	t.Run("bank-empty-failed", func(t *testing.T) {
		mockRequest := &request.BankCreateRequest{
			BankName: "BANK BRI",
		}

		mockBankRepo.On("Create", mock.AnythingOfType("*domain.Bank")).Return(nil, errors.New("Bank empty")).Once()
		uc := usecase.NewBankUsecase(mockBankRepo)
		_, err := uc.Create(*mockRequest)
		assert.Error(t, err)
		mockBankRepo.AssertExpectations(t)
	})
}
