package models

import "time"

type WeightHistory struct {
	ID     uint      `gorm:"primaryKey"`
	UserID int64     `gorm:"not null"` // Foreign key
	Date   time.Time `gorm:"not null"`
	Weight float64   `gorm:"not null"`
}
