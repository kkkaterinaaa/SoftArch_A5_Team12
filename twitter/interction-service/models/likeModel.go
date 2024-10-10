package models

import "time"

type Like struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	MessageID uint `gorm:"not null"`
	CreatedAt time.Time
}
