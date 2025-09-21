package models

import "github.com/FranciscoJSB12/gym-tracker-api-go/db"

type Routine struct {
	ID        uint       `gorm:"primaryKey"`
	UserID    uint       `gorm:"not null"` // Foreign key
	Name      string     `gorm:"type:varchar(255);not null"`
	Exercises []Exercise `gorm:"foreignKey:RoutineID"` // One-to-many relationship
}

func (r *Routine) CreateRoutine() error {
	result := db.DB.Create(&r)

	if result != nil {
		return result.Error
	}

	return nil
}

func GetRoutinesByUser(userID uint, routines *[]Routine) error {
	result := db.DB.Where("user_id = ?", userID).Find(routines)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
