package routes

import (
	"github.com/FranciscoJSB12/gym-tracker-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterWeightHistoryRoutes(server *gin.Engine) {
	routineGroup := server.Group("/weight-histories")
	{
		routineGroup.POST("/", controllers.CreateProgressHistory)
		routineGroup.GET("/", controllers.GetAllProgressByExercise)
	}
}
