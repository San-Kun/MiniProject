package usecase

import (
	"errors"

	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/domain/web/request"
)

type eventUsecase struct {
	EventRepo domain.EventRepository
}

func NewEventUsecase(ur domain.EventRepository) domain.EventUsecase {
	return &eventUsecase{EventRepo: ur}
}

func (u *eventUsecase) Create(request request.EventCreateRequest) (*domain.Event, error) {
	if request.EventName == "" {
		return nil, errors.New("email empty")
	}
	event := &domain.Event{
		EventName:    request.EventName,
		EventTanggal: request.EventTanggal,
	}

	createdEvent, err := u.EventRepo.Create(event)
	if err != nil {
		return nil, err
	}

	return createdEvent, nil
}

func (u *eventUsecase) ReadByID(id int) (*domain.Event, error) {
	event, err := u.EventRepo.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return event, err
}

func (u *eventUsecase) ReadAll() (*domain.Events, error) {
	foundEvents, err := u.EventRepo.ReadAll()
	if err != nil {
		return nil, err
	}

	return foundEvents, nil
}

func (u *eventUsecase) Delete(id int) (*domain.Event, error) {
	event, err := u.EventRepo.Delete(id)
	if err != nil {
		return nil, err
	}

	return event, err
}

func (u *eventUsecase) Updates(id int) (*domain.Event, error) {
	event, err := u.EventRepo.Updates(id)
	if err != nil {
		return nil, err
	}

	return event, err
}
