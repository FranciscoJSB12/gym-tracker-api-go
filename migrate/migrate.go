package main

import (
	"github.com/FranciscoJSB12/gym-tracker-api-go/db"
	"github.com/FranciscoJSB12/gym-tracker-api-go/models"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	db.InitDB()
}

func main() {
	db.DB.AutoMigrate(
		&models.User{},
		&models.Routine{},
		&models.Exercise{},
		&models.ProgressHistory{},
		&models.WeightHistory{},
	)
}
