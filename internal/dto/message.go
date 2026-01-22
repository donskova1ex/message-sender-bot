package dto

import "time"

type CreateMessageRequest struct {
	PlannedDate time.Time `json:"planned_date"`
	TypeID      int64     `json:"type_id"`
	Text        string    `json:"text"`
}
