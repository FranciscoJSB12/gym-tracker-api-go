package models

type User struct {
	ID       int64
	Email    string  `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string  `gorm:"type:varchar(255);not null"`
	Name     string  `gorm:"type:varchar(255);not null"`
	HeightCM float64 `gorm:"not null"`
}
