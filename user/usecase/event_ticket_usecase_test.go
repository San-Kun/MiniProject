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

func TestEventTicketUsecase_ReadAll(t *testing.T) {
	mockEventTicketRepo := new(mocks.EventTicketRepository)
	mockListEventTicket := &domain.EventTickets{
		domain.EventTicket{
			ID:          1,
			EventID:     1,
			PaymentID:   1,
			PayTypeID:   1,
			PayTypeName: "Transfer Bank BRI / BNI / BCA",
		},
		domain.EventTicket{
			ID:          2,
			EventID:     2,
			PaymentID:   2,
			PayTypeID:   2,
			PayTypeName: "Via Dana / GOPAY / OVO",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockEventTicketRepo.On("ReadAll").Return(mockListEventTicket, nil).Once()
		uc := usecase.NewEventTicketUsecase(mockEventTicketRepo)
		res, err := uc.ReadAll()
		assert.NoError(t, err)
		assert.Len(t, *res, len(*mockListEventTicket))
		mockEventTicketRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockEventTicketRepo.On("ReadAll").Return(nil, errors.New("error something")).Once()
		uc := usecase.NewEventTicketUsecase(mockEventTicketRepo)
		_, err := uc.ReadAll()
		assert.Error(t, err)
		mockEventTicketRepo.AssertExpectations(t)
	})
}

func TestEventTicketUsecase_ReadByID(t *testing.T) {
	mockEventTicketRepo := new(mocks.EventTicketRepository)
	mockEventTicket := &domain.EventTicket{
		ID:          1,
		EventID:     1,
		PaymentID:   1,
		PayTypeID:   1,
		PayTypeName: "Transfer Bank BRI / BNI / BCA",
	}

	t.Run("success", func(t *testing.T) {
		mockEventTicketRepo.On("ReadByID", mock.AnythingOfType("int")).Return(mockEventTicket, nil).Once()
		uc := usecase.NewEventTicketUsecase(mockEventTicketRepo)
		res, err := uc.ReadByID(1)
		assert.NoError(t, err)
		assert.Equal(t, res, mockEventTicket)
		mockEventTicketRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockEventTicketRepo.On("ReadByID", mock.AnythingOfType("int")).Return(nil, errors.New("error something")).Once()
		uc := usecase.NewEventTicketUsecase(mockEventTicketRepo)
		_, err := uc.ReadByID(2)
		assert.Error(t, err)
		mockEventTicketRepo.AssertExpectations(t)
	})
}

func TestEventTicketUsecase_Create(t *testing.T) {
	mockEventTicketRepo := new(mocks.EventTicketRepository)
	mockEventTicket := &domain.EventTicket{
		ID:          1,
		EventID:     1,
		PaymentID:   1,
		PayTypeID:   1,
		PayTypeName: "Transfer Bank BRI / BNI / BCA",
	}

	t.Run("success", func(t *testing.T) {
		mockRequest := &request.EventTicketCreateRequest{
			PayTypeName: "Transfer Bank BRI / BNI / BCA",
		}

		mockEventTicketRepo.On("Create", mock.AnythingOfType("*domain.EventTicket")).Return(mockEventTicket, nil).Once()
		uc := usecase.NewEventTicketUsecase(mockEventTicketRepo)
		res, err := uc.Create(*mockRequest)
		assert.NoError(t, err)
		assert.Equal(t, res, mockEventTicket)
		mockEventTicketRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockRequest := &request.EventTicketCreateRequest{
			PayTypeName: "Transfer Bank BRI / BNI / BCA",
		}

		mockEventTicketRepo.On("Create", mock.AnythingOfType("*domain.EventTicket")).Return(nil, errors.New("error something")).Once()
		uc := usecase.NewEventTicketUsecase(mockEventTicketRepo)
		_, err := uc.Create(*mockRequest)
		assert.Error(t, err)
		mockEventTicketRepo.AssertExpectations(t)
	})

	t.Run("bank-empty-failed", func(t *testing.T) {
		mockRequest := &request.EventTicketCreateRequest{
			PayTypeName: "Transfer Bank BRI / BNI / BCA",
		}

		mockEventTicketRepo.On("Create", mock.AnythingOfType("*domain.EventTicket")).Return(nil, errors.New("EventTicket empty")).Once()
		uc := usecase.NewEventTicketUsecase(mockEventTicketRepo)
		_, err := uc.Create(*mockRequest)
		assert.Error(t, err)
		mockEventTicketRepo.AssertExpectations(t)
	})
}
