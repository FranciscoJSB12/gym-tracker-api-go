package routes

import "github.com/gin-gonic/gin"

func RegisterAppRoutes(server *gin.Engine) {
	registerUserRoutes(server)
	registerRoutineRoutes(server)
	registerExerciseRoutes(server)
	registerProgressHistoryRoutes(server)
	registerWeightHistoryRoutes(server)
}
