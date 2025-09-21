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

func createExercise(context *gin.Context) {
	var exercise models.Exercise

	if err := context.ShouldBindJSON(&exercise); err != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Could not parse request data",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	err := exercise.CreateExcercise()

	if err != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Could not create exercise",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	utils.Respond(context, &dtos.Response{
		Status:  http.StatusCreated,
		Message: "Exercise created successfully",
		Ok:      true,
		Data:    gin.H{"exerciseId": exercise.ID},
	})
}

func getExercisesByRoutine(context *gin.Context) {
	routineID, err := strconv.ParseUint(context.Param("id"), 10, 64)

	if err != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid routine ID",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	var exercises []models.Exercise

	err = models.GetExercisesByRoutine(uint(routineID), &exercises)

	if err != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Could not retrieve exercises",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	if len(exercises) == 0 {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusOK,
			Message: "No exercises found for this routine",
			Ok:      true,
			Data:    []models.Exercise{},
		})
		return
	}

	var exerciseResponses []dtos.ExerciseResponse

	for _, exercise := range exercises {
		exerciseResponses = append(exerciseResponses, dtos.ExerciseResponse{
			ID:             exercise.ID,
			Name:           exercise.Name,
			RoutineID:      exercise.RoutineID,
			ProposedRounds: exercise.ProposedRounds,
			Repetitions:    exercise.Repetitions,
		})
	}

	utils.Respond(context, &dtos.Response{
		Status:  http.StatusOK,
		Message: "Exercises retrieved successfully",
		Ok:      true,
		Data:    exerciseResponses,
	})
}

func registerExerciseRoutes(server *gin.Engine) {
	group := server.Group("/exercises")

	group.Use(middlewares.Authenticate)
	group.POST("/", createExercise)
	group.GET("/routine/:id", getExercisesByRoutine)
}
