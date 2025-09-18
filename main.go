package main

import (
	"github.com/FranciscoJSB12/gym-tracker-api-go/db"
	"github.com/FranciscoJSB12/gym-tracker-api-go/routes"
	"github.com/FranciscoJSB12/gym-tracker-api-go/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadEnvVariables()
	db.InitDB()
}

func main() {
	server := gin.Default()

	routes.RegisterUserRoutes(server)

	server.Run()
}
