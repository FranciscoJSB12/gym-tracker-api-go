package models

import (
	"github.com/FranciscoJSB12/gym-tracker-api-go/db"
)

type Exercise struct {
	ID             uint              `gorm:"primaryKey"`
	Name           string            `gorm:"type:varchar(255);not null"`
	RoutineID      uint              `gorm:"not null"` // Foreign key
	ProposedRounds uint              `gorm:"not null"`
	Repetitions    uint              `gorm:"not null"`
	Progresses     []ProgressHistory `gorm:"foreignKey:ExerciseID"` // One-to-many relationship
}

func (e *Exercise) CreateExcercise() error {
	result := db.DB.Create(e)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetExercisesByRoutine(routineID uint, exercises *[]Exercise) error {
	result := db.DB.Where("routine_id = ?", routineID).Find(exercises)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
