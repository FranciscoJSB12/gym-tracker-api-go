package models

import "time"

type WeightHistory struct {
	ID       uint      `gorm:"primaryKey"`
	UserID   uint      `gorm:"not null"` // Foreign key
	Date     time.Time `gorm:"not null"`
	WeightKG float64   `gorm:"not null"`
}
