package routes

import (
	"net/http"
	"strconv"

	"github.com/FranciscoJSB12/gym-tracker-api-go/dtos"
	"github.com/FranciscoJSB12/gym-tracker-api-go/middlewares"
	"github.com/FranciscoJSB12/gym-tracker-api-go/models"
	"github.com/FranciscoJSB12/gym-tracker-api-go/utils"
	"github.com/gin-gonic/gin"
)

func createProgressHistory(context *gin.Context) {
	var progress models.ProgressHistory

	if err := context.ShouldBindJSON(&progress); err != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Could not parse request data",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	err := progress.CreateProgressHistory()

	if err != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Could not create progress history",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	utils.Respond(context, &dtos.Response{
		Status:  http.StatusCreated,
		Message: "Progress history created successfully",
		Ok:      true,
		Data:    gin.H{"progressId": progress.ID},
	})
}

func getProgressByExercise(context *gin.Context) {
	exerciseID, err := strconv.ParseUint(context.Param("id"), 10, 64)

	if err != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid exercise ID",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	var progress []models.ProgressHistory

	err = models.GetProgressByExercise(uint(exerciseID), &progress)

	if err != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Could not retrieve progress history",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	if len(progress) == 0 {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusOK,
			Message: "No progress history found for this exercise",
			Ok:      true,
			Data:    []models.ProgressHistory{},
		})
		return
	}

	utils.Respond(context, &dtos.Response{
		Status:  http.StatusOK,
		Message: "Progress history retrieved successfully",
		Ok:      true,
		Data:    progress,
	})
}

func registerProgressHistoryRoutes(server *gin.Engine) {
	group := server.Group("/progress-histories")

	group.Use(middlewares.Authenticate)
	group.POST("/", createProgressHistory)
	group.GET("/exercise/:id", getProgressByExercise)
}
