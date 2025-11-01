package models

import "time"

type url struct {
	ID           uint      `gorm:"primaryKey"`
	ShortCode    string    `gorm:"uniqueIndex;size:8;not null"`
	OriginalIURL string    `gorm:"not null"`
	Clicks       int       `gorm:"default:0"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UserID       *uint
}
