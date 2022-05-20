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

func TestPaymentTypeUsecase_ReadAll(t *testing.T) {
	mockPaymentTypeRepo := new(mocks.PaymentTypeRepository)
	mockListPaymentType := &domain.PaymentTypes{
		domain.PaymentType{
			ID:          1,
			PayTypeName: "Transfer Bank BRI / BNI / BCA",
		},
		domain.PaymentType{
			ID:          2,
			PayTypeName: "Via Dana / GOPAY / OVO",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockPaymentTypeRepo.On("ReadAll").Return(mockListPaymentType, nil).Once()
		uc := usecase.NewPaymentTypeUsecase(mockPaymentTypeRepo)
		res, err := uc.ReadAll()
		assert.NoError(t, err)
		assert.Len(t, *res, len(*mockListPaymentType))
		mockPaymentTypeRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockPaymentTypeRepo.On("ReadAll").Return(nil, errors.New("error something")).Once()
		uc := usecase.NewPaymentTypeUsecase(mockPaymentTypeRepo)
		_, err := uc.ReadAll()
		assert.Error(t, err)
		mockPaymentTypeRepo.AssertExpectations(t)
	})
}

func TestPaymentTypeUsecase_ReadByID(t *testing.T) {
	mockPaymentTypeRepo := new(mocks.PaymentTypeRepository)
	mockPaymentType := &domain.PaymentType{
		ID:          1,
		PayTypeName: "Transfer Bank BRI / BNI / BCA",
	}

	t.Run("success", func(t *testing.T) {
		mockPaymentTypeRepo.On("ReadByID", mock.AnythingOfType("int")).Return(mockPaymentType, nil).Once()
		uc := usecase.NewPaymentTypeUsecase(mockPaymentTypeRepo)
		res, err := uc.ReadByID(1)
		assert.NoError(t, err)
		assert.Equal(t, res, mockPaymentType)
		mockPaymentTypeRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockPaymentTypeRepo.On("ReadByID", mock.AnythingOfType("int")).Return(nil, errors.New("error something")).Once()
		uc := usecase.NewPaymentTypeUsecase(mockPaymentTypeRepo)
		_, err := uc.ReadByID(2)
		assert.Error(t, err)
		mockPaymentTypeRepo.AssertExpectations(t)
	})
}

func TestPaymentTypeUsecase_Create(t *testing.T) {
	mockPaymentTypeRepo := new(mocks.PaymentTypeRepository)
	mockPaymentType := &domain.PaymentType{
		ID:          1,
		PayTypeName: "Transfer Bank BRI / BNI / BCA",
	}

	t.Run("success", func(t *testing.T) {
		mockRequest := &request.PaymentTypeCreateRequest{
			PayTypeName: "Transfer Bank BRI / BNI / BCA",
		}

		mockPaymentTypeRepo.On("Create", mock.AnythingOfType("*domain.PaymentType")).Return(mockPaymentType, nil).Once()
		uc := usecase.NewPaymentTypeUsecase(mockPaymentTypeRepo)
		res, err := uc.Create(*mockRequest)
		assert.NoError(t, err)
		assert.Equal(t, res, mockPaymentType)
		mockPaymentTypeRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockRequest := &request.PaymentTypeCreateRequest{
			PayTypeName: "Transfer Bank BRI / BNI / BCA",
		}

		mockPaymentTypeRepo.On("Create", mock.AnythingOfType("*domain.PaymentType")).Return(nil, errors.New("error something")).Once()
		uc := usecase.NewPaymentTypeUsecase(mockPaymentTypeRepo)
		_, err := uc.Create(*mockRequest)
		assert.Error(t, err)
		mockPaymentTypeRepo.AssertExpectations(t)
	})

	t.Run("paymentType-empty-failed", func(t *testing.T) {
		mockRequest := &request.PaymentTypeCreateRequest{
			PayTypeName: "Transfer Bank BRI / BNI / BCA",
		}

		mockPaymentTypeRepo.On("Create", mock.AnythingOfType("*domain.PaymentType")).Return(nil, errors.New("PaymentType empty")).Once()
		uc := usecase.NewPaymentTypeUsecase(mockPaymentTypeRepo)
		_, err := uc.Create(*mockRequest)
		assert.Error(t, err)
		mockPaymentTypeRepo.AssertExpectations(t)
	})
}
