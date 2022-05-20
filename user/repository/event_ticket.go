package repository

import (
	"github.com/San-Kun/MiniProject/domain"
	"gorm.io/gorm"
)

type eventTicketRepository struct {
	Conn *gorm.DB
}

func NewEventTicketRepository(Conn *gorm.DB) domain.EventTicketRepository {
	return &eventTicketRepository{Conn: Conn}
}

func (u *eventTicketRepository) Create(eventTicket *domain.EventTicket) (*domain.EventTicket, error) {
	if err := u.Conn.Create(&eventTicket).Error; err != nil {
		return nil, err
	}

	return eventTicket, nil
}

func (u *eventTicketRepository) ReadByID(id int) (*domain.EventTicket, error) {
	eventTicket := &domain.EventTicket{ID: id}
	if err := u.Conn.First(&eventTicket).Error; err != nil {
		return nil, err
	}

	return eventTicket, nil
}

func (u *eventTicketRepository) ReadAll() (*domain.EventTickets, error) {
	eventTickets := &domain.EventTickets{}
	u.Conn.Find(&eventTickets)

	return eventTickets, nil
}

func (u *eventTicketRepository) Delete(id int) (*domain.EventTicket, error) {
	eventTicket := &domain.EventTicket{ID: id}
	if err := u.Conn.Delete(&eventTicket).Error; err != nil {
		return nil, err
	}
	return eventTicket, nil
}

func (u *eventTicketRepository) Updates(id int) (*domain.EventTicket, error) {
	eventTicket := &domain.EventTicket{ID: id}
	if err := u.Conn.Updates(&eventTicket).Error; err != nil {
		return nil, err
	}

	return eventTicket, nil
}
