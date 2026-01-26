package dto

import "time"

type Response struct {
	Status       string      `json:"status"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
	ResponseDate time.Time   `json:"response_date,omitempty"`
}
type ScheduleMessageRequest struct {
	PlannedDate time.Time `json:"planned_date"`
	TypeID      int64     `json:"type_id"`
	Text        string    `json:"text"`
}

type MessageTypeRequest struct {
	Name string `json:"name"`
}

type UpdateMessageRequest struct {
	PlannedDate *time.Time `json:"planned_date,omitempty"`
	TypeID      *int64     `json:"type_id,omitempty"`
	Text        *string    `json:"text,omitempty"`
}

type MessagesResponse struct {
	ID          int64     `json:"id"`
	PlannedDate time.Time `json:"planned_date"`
	Type        string    `json:"type_name"`
	Text        string    `json:"text"`
	IsSent      bool      `json:"is_sent"`
}
