package routes

import (
	"github.com/FranciscoJSB12/gym-tracker-api-go/controllers"
	"github.com/FranciscoJSB12/gym-tracker-api-go/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterWeightHistoryRoutes(server *gin.Engine) {
	group := server.Group("/weight-histories")
	group.Use(middlewares.Authenticate)
	group.POST("/", controllers.CreateWeightHistory)
	group.GET("/", controllers.GetAllWeightHistoryByUser)

}
