package routes

import (
	"github.com/FranciscoJSB12/gym-tracker-api-go/controllers"
	"github.com/FranciscoJSB12/gym-tracker-api-go/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutineRoutes(server *gin.Engine) {
	group := server.Group("/routines")

	group.Use(middlewares.Authenticate)
	group.POST("/", controllers.CreateRoutine)
	group.GET("/", controllers.GetAllRoutinesByUser)

}
