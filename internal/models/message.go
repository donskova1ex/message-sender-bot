package models

import "time"

type Message struct {
	ID          int       `db:"id"`
	UserID      int       `db:"user_id"`
	PlannedDate time.Time `db:"planned_date"`
	TypeID      int64     `db:"type_id"`
	Text        string    `db:"text"`
	CreatedAt   time.Time `db:"created_at"`
}
