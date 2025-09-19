package routes

import (
	"github.com/FranciscoJSB12/gym-tracker-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine) {
	group := server.Group("/users")

	group.POST("/sign-up", controllers.SignUp)
	group.POST("/login", controllers.Login)

}
