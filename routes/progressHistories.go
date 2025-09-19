package routes

import (
	"github.com/FranciscoJSB12/gym-tracker-api-go/controllers"
	"github.com/FranciscoJSB12/gym-tracker-api-go/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterProgressHistoryRoutes(server *gin.Engine) {
	group := server.Group("/progress-histories")

	group.Use(middlewares.Authenticate)
	group.POST("/", controllers.CreateProgressHistory)
	group.GET("/exercise/:id", controllers.GetAllProgressByExercise)

}
