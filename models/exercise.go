package models

type Exercise struct {
	ID             uint              `gorm:"primaryKey"`
	Name           string            `gorm:"type:varchar(255);not null"`
	RoutineID      uint              `gorm:"not null"` // Foreign key
	ProposedRounds int64             `gorm:"not null"`
	Progresses     []ProgressHistory `gorm:"foreignKey:ExerciseID"` // One-to-many relationship
}
