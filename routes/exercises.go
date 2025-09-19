package routes

import (
	"github.com/FranciscoJSB12/gym-tracker-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterExerciseRoutes(server *gin.Engine) {
	routineGroup := server.Group("/exercises")
	{
		routineGroup.POST("/", controllers.CreateExercise)
		routineGroup.GET("/routine/:id", controllers.GetAllExercisesByRoutine)
	}
}
