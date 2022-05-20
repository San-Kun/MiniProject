package repository

import (
	"github.com/San-Kun/MiniProject/domain"
	"gorm.io/gorm"
)

type eventRepository struct {
	Conn *gorm.DB
}

func NewEventRepository(Conn *gorm.DB) domain.EventRepository {
	return &eventRepository{Conn: Conn}
}

func (u *eventRepository) Create(event *domain.Event) (*domain.Event, error) {
	if err := u.Conn.Create(&event).Error; err != nil {
		return nil, err
	}

	return event, nil
}

func (u *eventRepository) ReadByID(id int) (*domain.Event, error) {
	event := &domain.Event{ID: id}
	if err := u.Conn.First(&event).Error; err != nil {
		return nil, err
	}

	return event, nil
}

func (u *eventRepository) ReadAll() (*domain.Events, error) {
	events := &domain.Events{}
	u.Conn.Find(&events)

	return events, nil
}

func (u *eventRepository) Delete(id int) (*domain.Event, error) {
	event := &domain.Event{ID: id}
	if err := u.Conn.Delete(&event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (u *eventRepository) Updates(id int) (*domain.Event, error) {
	event := &domain.Event{ID: id}
	if err := u.Conn.Updates(&event).Error; err != nil {
		return nil, err
	}

	return event, nil
}
