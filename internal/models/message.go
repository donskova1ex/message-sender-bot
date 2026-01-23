package models

import "time"

type Message struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	PlannedDate time.Time `db:"planned_date"`
	TypeID      int64     `db:"type_id"`
	Text        string    `db:"text"`
	CreatedAt   time.Time `db:"created_at"`
}

type UnsentMessage struct {
	ID          int64     `db:"id"`
	PlannedDate time.Time `db:"planned_date"`
	TypeName    string    `db:"type_name"`
	Text        string    `db:"text"`
}
