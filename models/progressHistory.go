package models

import (
	"time"

	"github.com/FranciscoJSB12/gym-tracker-api-go/db"
)

type ProgressHistory struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ExerciseID   uint      `gorm:"not null" json:"exerciseID"` // Foreign key
	WeightUsedKG float32   `gorm:"not null" json:"weightUsedKG"`
	RepsDone     uint      `gorm:"default:0" json:"repsDone"`
	RoundsDone   uint      `gorm:"default:0" json:"roundsDone"`
	Date         time.Time `gorm:"not null" json:"date"`
}

func (p *ProgressHistory) CreateProgressHistory() error {
	result := db.DB.Create(&p)

	if result != nil {
		return result.Error
	}

	return nil
}

func GetProgressByExercise(exerciseID uint, progress *[]ProgressHistory) error {
	result := db.DB.Where("exercise_id = ?", exerciseID).Find(progress)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
