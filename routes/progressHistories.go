package routes

import (
	"github.com/FranciscoJSB12/gym-tracker-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterProgressHistoryRoutes(server *gin.Engine) {
	routineGroup := server.Group("/progress-histories")
	{
		routineGroup.POST("/", controllers.CreateProgressHistory)
		routineGroup.GET("/exercise/:id", controllers.GetAllProgressByExercise)
	}
}
