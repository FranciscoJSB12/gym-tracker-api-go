package models

type User struct {
	ID       int64
	Email    string  `binding:"required"`
	Password string  `binding:"required"`
	Name     string  `binding:"required"`
	HeightCM float64 `binding:"required"`
}
