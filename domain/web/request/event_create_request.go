package request

type EventCreateRequest struct {
	ID           int    `json:"id"`
	EventName    string `json:"event_name" validate:"required"`
	EventTanggal string `json:"event_tanggal" validate:"required"`
	Capacity     string `json:"capacity" validate:"required"`
}
