package models

import (
	"time"

	"github.com/FranciscoJSB12/gym-tracker-api-go/db"
)

type WeightHistory struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	UserID   uint      `gorm:"not null" json:"userID"` // Foreign key
	Date     time.Time `gorm:"not null" json:"date"`
	WeightKG float64   `gorm:"not null" json:"weightKG"`
}

func (wh *WeightHistory) CreateWeightHistory() error {
	result := db.DB.Create(&wh)

	if result != nil {
		return result.Error
	}

	return nil
}

func GetWeightHistoryByUser(userID uint, history *[]WeightHistory) error {
	result := db.DB.Where("user_id = ?", userID).Find(history)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
