package domain

import (
	"github.com/San-Kun/MiniProject/domain/web/request"
)

type EventTicket struct {
	ID          int     `json:"id"`
	EventID     int     `json:"event_id"`
	PaymentID   int     `json:"payment_id"`
	PayTypeID   int     `json:"pay_type_id"`
	PayTypeName string  `json:"paytype_name"`
	Event       Event   `json:"events" gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Payment     Payment `json:"payments" gorm:"foreignKey:PaymentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type EventTickets []EventTicket

type EventTicketRepository interface {
	Create(eventTicket *EventTicket) (*EventTicket, error)
	ReadByID(id int) (*EventTicket, error)
	Delete(id int) (*EventTicket, error)
	Updates(id int) (*EventTicket, error)
	ReadAll() (*EventTickets, error)
}

type EventTicketUsecase interface {
	Create(request request.EventTicketCreateRequest) (*EventTicket, error)
	ReadByID(id int) (*EventTicket, error)
	Delete(id int) (*EventTicket, error)
	Updates(id int) (*EventTicket, error)
	ReadAll() (*EventTickets, error)
}
