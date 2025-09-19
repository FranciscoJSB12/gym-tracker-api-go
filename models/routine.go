package models

type Routine struct {
	ID        uint       `gorm:"primaryKey"`
	UserID    uint       // Foreign key
	Name      string     `gorm:"type:varchar(255);not null"`
	Exercises []Exercise `gorm:"foreignKey:RoutineID"` // One-to-many relationship
}
