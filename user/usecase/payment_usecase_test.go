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

func TestPaymentUsecase_ReadAll(t *testing.T) {
	mockPaymentRepo := new(mocks.PaymentRepository)
	mockListPayment := &domain.Payments{
		domain.Payment{
			ID:            1,
			PaymentDate:   "12 Maret 2022",
			PaymentAmount: 45000,
		},
		domain.Payment{
			ID:            2,
			PaymentDate:   "20 Maret 2022",
			PaymentAmount: 50000,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockPaymentRepo.On("ReadAll").Return(mockListPayment, nil).Once()
		uc := usecase.NewPaymentUsecase(mockPaymentRepo)
		res, err := uc.ReadAll()
		assert.NoError(t, err)
		assert.Len(t, *res, len(*mockListPayment))
		mockPaymentRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockPaymentRepo.On("ReadAll").Return(nil, errors.New("error something")).Once()
		uc := usecase.NewPaymentUsecase(mockPaymentRepo)
		_, err := uc.ReadAll()
		assert.Error(t, err)
		mockPaymentRepo.AssertExpectations(t)
	})
}

func TestPaymentUsecase_ReadByID(t *testing.T) {
	mockPaymentRepo := new(mocks.PaymentRepository)
	mockPayment := &domain.Payment{
		ID:            1,
		PaymentDate:   "12 Maret 2022",
		PaymentAmount: 45000,
	}

	t.Run("success", func(t *testing.T) {
		mockPaymentRepo.On("ReadByID", mock.AnythingOfType("int")).Return(mockPayment, nil).Once()
		uc := usecase.NewPaymentUsecase(mockPaymentRepo)
		res, err := uc.ReadByID(1)
		assert.NoError(t, err)
		assert.Equal(t, res, mockPayment)
		mockPaymentRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockPaymentRepo.On("ReadByID", mock.AnythingOfType("int")).Return(nil, errors.New("error something")).Once()
		uc := usecase.NewPaymentUsecase(mockPaymentRepo)
		_, err := uc.ReadByID(2)
		assert.Error(t, err)
		mockPaymentRepo.AssertExpectations(t)
	})
}

func TestPaymentUsecase_Create(t *testing.T) {
	mockPaymentRepo := new(mocks.PaymentRepository)
	mockPayment := &domain.Payment{
		ID:            1,
		PaymentDate:   "12 Maret 2022",
		PaymentAmount: 45000,
	}

	t.Run("success", func(t *testing.T) {
		mockRequest := &request.PaymentCreateRequest{
			PaymentDate:   "12 Maret 2022",
			PaymentAmount: 45000,
		}

		mockPaymentRepo.On("Create", mock.AnythingOfType("*domain.Payment")).Return(mockPayment, nil).Once()
		uc := usecase.NewPaymentUsecase(mockPaymentRepo)
		res, err := uc.Create(*mockRequest)
		assert.NoError(t, err)
		assert.Equal(t, res, mockPayment)
		mockPaymentRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockRequest := &request.PaymentCreateRequest{
			PaymentDate:   "12 Maret 2022",
			PaymentAmount: 45000,
		}

		mockPaymentRepo.On("Create", mock.AnythingOfType("*domain.Payment")).Return(nil, errors.New("error something")).Once()
		uc := usecase.NewPaymentUsecase(mockPaymentRepo)
		_, err := uc.Create(*mockRequest)
		assert.Error(t, err)
		mockPaymentRepo.AssertExpectations(t)
	})

	t.Run("payment-empty-failed", func(t *testing.T) {
		mockRequest := &request.PaymentCreateRequest{
			PaymentDate:   "12 Maret 2022",
			PaymentAmount: 45000,
		}

		mockPaymentRepo.On("Create", mock.AnythingOfType("*domain.Payment")).Return(nil, errors.New("Payment empty")).Once()
		uc := usecase.NewPaymentUsecase(mockPaymentRepo)
		_, err := uc.Create(*mockRequest)
		assert.Error(t, err)
		mockPaymentRepo.AssertExpectations(t)
	})
}
