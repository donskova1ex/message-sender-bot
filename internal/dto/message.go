package dto

import "time"

type ScheduleMessageRequest struct {
	PlannedDate time.Time `json:"planned_date"`
	TypeID      int64     `json:"type_id"`
	Text        string    `json:"text"`
}

type MessageTypeRequest struct {
	Name string `json:"name"`
}