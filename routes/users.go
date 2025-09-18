package routes

import (
	"github.com/FranciscoJSB12/gym-tracker-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine) {
	userGroup := server.Group("/users")
	{
		userGroup.POST("/sign-up", controllers.SignUp)
		userGroup.POST("/login", controllers.Login)
	}
}
