package domain

import (
	"github.com/San-Kun/MiniProject/domain/web/request"
)

type Event struct {
	ID           int    `json:"id"`
	EventName    string `json:"event_name"`
	EventTanggal string `json:"event_tanggal"`
	Capacity     string `json:"capacity"`
}

type Events []Event

type EventRepository interface {
	Create(event *Event) (*Event, error)
	ReadByID(id int) (*Event, error)
	Delete(id int) (*Event, error)
	Updates(id int) (*Event, error)
	ReadAll() (*Events, error)
}

type EventUsecase interface {
	Create(request request.EventCreateRequest) (*Event, error)
	ReadByID(id int) (*Event, error)
	Delete(id int) (*Event, error)
	Updates(id int) (*Event, error)
	ReadAll() (*Events, error)
}
