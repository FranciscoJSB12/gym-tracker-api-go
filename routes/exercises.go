package routes

import (
	"github.com/FranciscoJSB12/gym-tracker-api-go/controllers"
	"github.com/FranciscoJSB12/gym-tracker-api-go/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterExerciseRoutes(server *gin.Engine) {
	group := server.Group("/exercises")

	group.Use(middlewares.Authenticate)
	group.POST("/", controllers.CreateExercise)
	group.GET("/routine/:id", controllers.GetAllExercisesByRoutine)
}
