package models

import (
	"errors"

	"github.com/FranciscoJSB12/gym-tracker-api-go/db"
	"github.com/FranciscoJSB12/gym-tracker-api-go/utils"
)

type User struct {
	ID            uint            `gorm:"primaryKey"`
	Email         string          `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password      string          `gorm:"type:varchar(255);not null"`
	Name          string          `gorm:"type:varchar(255);not null"`
	HeightCM      float32         `gorm:"not null"`
	Routines      []Routine       `gorm:"foreignKey:UserID"` // One-to-many relationship with Routine
	WeightHistory []WeightHistory `gorm:"foreignKey:UserID"` // One-to-many relationship with WeightHistory
}

func (u *User) Save() error {
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	u.Password = hashedPassword

	result := db.DB.Create(u)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) ValidateCredentials() error {
	enteredPassword := u.Password

	result := db.DB.Where("email = ?", u.Email).First(u)

	if result.Error != nil {
		return errors.New("invalid credentials")
	}

	isPasswordValid := utils.CheckPasswordHash(enteredPassword, u.Password)

	if !isPasswordValid {
		return errors.New("invalid credentials")
	}

	return nil
}
