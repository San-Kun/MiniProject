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

func TestEventUsecase_ReadAll(t *testing.T) {
	mockEventRepo := new(mocks.EventRepository)
	mockListEvent := &domain.Events{
		domain.Event{
			ID:           1,
			EventName:    "Forum Diskusi",
			EventTanggal: "12 Maret 2022",
			Capacity:     "100",
		},
		domain.Event{
			ID:           2,
			EventName:    "Workshop Kebhinekaan",
			EventTanggal: "20 Maret 2022",
			Capacity:     "150",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockEventRepo.On("ReadAll").Return(mockListEvent, nil).Once()
		uc := usecase.NewEventUsecase(mockEventRepo)
		res, err := uc.ReadAll()
		assert.NoError(t, err)
		assert.Len(t, *res, len(*mockListEvent))
		mockEventRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockEventRepo.On("ReadAll").Return(nil, errors.New("error something")).Once()
		uc := usecase.NewEventUsecase(mockEventRepo)
		_, err := uc.ReadAll()
		assert.Error(t, err)
		mockEventRepo.AssertExpectations(t)
	})
}

func TestEventUsecase_ReadByID(t *testing.T) {
	mockEventRepo := new(mocks.EventRepository)
	mockEvent := &domain.Event{
		ID:           1,
		EventName:    "Forum Diskusi",
		EventTanggal: "12 Maret 2022",
		Capacity:     "100",
	}

	t.Run("success", func(t *testing.T) {
		mockEventRepo.On("ReadByID", mock.AnythingOfType("int")).Return(mockEvent, nil).Once()
		uc := usecase.NewEventUsecase(mockEventRepo)
		res, err := uc.ReadByID(1)
		assert.NoError(t, err)
		assert.Equal(t, res, mockEvent)
		mockEventRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockEventRepo.On("ReadByID", mock.AnythingOfType("int")).Return(nil, errors.New("error something")).Once()
		uc := usecase.NewEventUsecase(mockEventRepo)
		_, err := uc.ReadByID(2)
		assert.Error(t, err)
		mockEventRepo.AssertExpectations(t)
	})
}

func TestEventUsecase_Create(t *testing.T) {
	mockEventRepo := new(mocks.EventRepository)
	mockEvent := &domain.Event{
		ID:           1,
		EventName:    "Forum Diskusi",
		EventTanggal: "12 Maret 2022",
		Capacity:     "100",
	}

	t.Run("success", func(t *testing.T) {
		mockRequest := &request.EventCreateRequest{
			EventName:    "Forum Diskusi",
			EventTanggal: "12 Maret 2022",
		}

		mockEventRepo.On("Create", mock.AnythingOfType("*domain.Event")).Return(mockEvent, nil).Once()
		uc := usecase.NewEventUsecase(mockEventRepo)
		res, err := uc.Create(*mockRequest)
		assert.NoError(t, err)
		assert.Equal(t, res, mockEvent)
		mockEventRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockRequest := &request.EventCreateRequest{
			EventName:    "Forum Diskusi",
			EventTanggal: "12 Maret 2022",
		}

		mockEventRepo.On("Create", mock.AnythingOfType("*domain.Event")).Return(nil, errors.New("error something")).Once()
		uc := usecase.NewEventUsecase(mockEventRepo)
		_, err := uc.Create(*mockRequest)
		assert.Error(t, err)
		mockEventRepo.AssertExpectations(t)
	})

	t.Run("event-empty-failed", func(t *testing.T) {
		mockRequest := &request.EventCreateRequest{
			EventName:    "Forum Diskusi",
			EventTanggal: "12 Maret 2022",
		}

		mockEventRepo.On("Create", mock.AnythingOfType("*domain.Event")).Return(nil, errors.New("event empty")).Once()
		uc := usecase.NewEventUsecase(mockEventRepo)
		_, err := uc.Create(*mockRequest)
		assert.Error(t, err)
		mockEventRepo.AssertExpectations(t)
	})
}
