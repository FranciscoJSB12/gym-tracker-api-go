package routes

import (
	"github.com/FranciscoJSB12/gym-tracker-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutineRoutes(server *gin.Engine) {
	routineGroup := server.Group("/routines")
	{
		routineGroup.POST("/", controllers.CreateRoutine)
		routineGroup.GET("/", controllers.GetAllRoutinesByUser)
	}
}
