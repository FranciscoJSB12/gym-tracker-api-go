package controllers

import (
	"net/http"
	"strconv"

	"github.com/FranciscoJSB12/gym-tracker-api-go/db"
	"github.com/FranciscoJSB12/gym-tracker-api-go/dtos"
	"github.com/FranciscoJSB12/gym-tracker-api-go/models"
	"github.com/FranciscoJSB12/gym-tracker-api-go/utils"
	"github.com/gin-gonic/gin"
)

func CreateExercise(context *gin.Context) {
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

	if result := db.DB.Create(&exercise); result.Error != nil {
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

func GetAllExercisesByRoutine(context *gin.Context) {
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
	if result := db.DB.Where("routine_id = ?", routineID).Find(&exercises); result.Error != nil {
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

	utils.Respond(context, &dtos.Response{
		Status:  http.StatusOK,
		Message: "Exercises retrieved successfully",
		Ok:      true,
		Data:    exercises,
	})
}
