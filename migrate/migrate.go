package main

import (
	"github.com/FranciscoJSB12/gym-tracker-api-go/db"
	"github.com/FranciscoJSB12/gym-tracker-api-go/models"
	"github.com/FranciscoJSB12/gym-tracker-api-go/utils"
)

func init() {
	utils.LoadEnvVariables()
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
