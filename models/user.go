package models

type User struct {
	ID            uint            `gorm:"primaryKey"`
	Email         string          `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password      string          `gorm:"type:varchar(255);not null"`
	Name          string          `gorm:"type:varchar(255);not null"`
	HeightCM      float32         `gorm:"not null"`
	Routines      []Routine       `gorm:"foreignKey:UserID"` // One-to-many relationship with Routine
	WeightHistory []WeightHistory `gorm:"foreignKey:UserID"` // One-to-many relationship with WeightHistory
}
