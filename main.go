package main

import (
	"log"

	"github.com/FranciscoJSB12/gym-tracker-api-go/db"
	"github.com/FranciscoJSB12/gym-tracker-api-go/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	db.InitDB()
}

func main() {
	server := gin.Default()

	routes.RegisterAppRoutes(server)

	log.Println("Server is running!")
	err := server.Run()

	if err != nil {
		log.Fatal("Server failed!", err)
	}
}
