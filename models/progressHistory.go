package models

import (
	"time"
)

type ProgressHistory struct {
	ID         uint      `gorm:"primaryKey"`
	ExerciseID uint      `gorm:"not null"` // Foreign key
	WeightUsed float32   `gorm:"not null"`
	Date       time.Time `gorm:"not null"`
}
