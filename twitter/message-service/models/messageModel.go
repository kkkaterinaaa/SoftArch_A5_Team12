package models

import "time"

type Message struct {
	ID        uint   `gorm:"primaryKey"` // id: ...
	UserID    uint   `gorm:"not null"`   // user_id: ..
	Content   string `gorm:"type:varchar(400);not null"`
	CreatedAt time.Time
}
