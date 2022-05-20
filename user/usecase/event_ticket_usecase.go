package usecase

import (
	"errors"

	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/domain/web/request"
)

type eventTicketUsecase struct {
	EventTicketRepo domain.EventTicketRepository
}

func NewEventTicketUsecase(ur domain.EventTicketRepository) domain.EventTicketUsecase {
	return &eventTicketUsecase{EventTicketRepo: ur}
}

func (u *eventTicketUsecase) Create(request request.EventTicketCreateRequest) (*domain.EventTicket, error) {
	if request.PayTypeName == "" {
		return nil, errors.New("eventTicket empty")
	}
	eventTicket := &domain.EventTicket{
		PayTypeName: request.PayTypeName,
	}

	createdEventTicket, err := u.EventTicketRepo.Create(eventTicket)
	if err != nil {
		return nil, err
	}

	return createdEventTicket, nil
}

func (u *eventTicketUsecase) ReadByID(id int) (*domain.EventTicket, error) {
	eventTicket, err := u.EventTicketRepo.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return eventTicket, err
}

func (u *eventTicketUsecase) ReadAll() (*domain.EventTickets, error) {
	foundEventTickets, err := u.EventTicketRepo.ReadAll()
	if err != nil {
		return nil, err
	}

	return foundEventTickets, nil
}

func (u *eventTicketUsecase) Delete(id int) (*domain.EventTicket, error) {
	eventTicket, err := u.EventTicketRepo.Delete(id)
	if err != nil {
		return nil, err
	}

	return eventTicket, err
}

func (u *eventTicketUsecase) Updates(id int) (*domain.EventTicket, error) {
	eventTicket, err := u.EventTicketRepo.Updates(id)
	if err != nil {
		return nil, err
	}

	return eventTicket, err
}
