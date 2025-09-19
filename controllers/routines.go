package controllers

import (
	"net/http"

	"github.com/FranciscoJSB12/gym-tracker-api-go/db"
	"github.com/FranciscoJSB12/gym-tracker-api-go/dtos"
	"github.com/FranciscoJSB12/gym-tracker-api-go/models"
	"github.com/FranciscoJSB12/gym-tracker-api-go/utils"
	"github.com/gin-gonic/gin"
)

func CreateRoutine(context *gin.Context) {
	var routine models.Routine

	if err := context.ShouldBindJSON(&routine); err != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Could not parse request data",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	if result := db.DB.Create(&routine); result.Error != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Could not create routine",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	utils.Respond(context, &dtos.Response{
		Status:  http.StatusCreated,
		Message: "Routine created successfully",
		Ok:      true,
		Data:    gin.H{"routineId": routine.ID},
	})
}

func GetAllRoutinesByUser(context *gin.Context) {
	userID := context.GetUint("userID")

	var routines []models.Routine

	if result := db.DB.Where("user_id = ?", userID).Find(&routines); result.Error != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Could not retrieve routines",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	if len(routines) == 0 {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusOK,
			Message: "No routines found for this user",
			Ok:      true,
			Data:    []dtos.RoutineResponse{},
		})
		return
	}

	// Map the models.Routine slice to dtosRoutine.RoutineResponse slice
	var routineResponses []dtos.RoutineResponse
	for _, routine := range routines {
		routineResponses = append(routineResponses, dtos.RoutineResponse{
			ID:     routine.ID,
			UserID: routine.UserID,
			Name:   routine.Name,
		})
	}

	utils.Respond(context, &dtos.Response{
		Status:  http.StatusOK,
		Message: "Routines retrieved successfully",
		Ok:      true,
		Data:    routineResponses,
	})
}
