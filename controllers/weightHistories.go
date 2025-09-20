package controllers

import (
	"net/http"

	"github.com/FranciscoJSB12/gym-tracker-api-go/db"
	"github.com/FranciscoJSB12/gym-tracker-api-go/dtos"
	"github.com/FranciscoJSB12/gym-tracker-api-go/models"
	"github.com/FranciscoJSB12/gym-tracker-api-go/utils"
	"github.com/gin-gonic/gin"
)

func CreateWeightHistory(context *gin.Context) {
	userID := context.GetUint("userID")

	var weightHistory models.WeightHistory

	if err := context.ShouldBindJSON(&weightHistory); err != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Could not parse request data",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	weightHistory.UserID = userID

	if result := db.DB.Create(&weightHistory); result.Error != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Could not create weight history",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	utils.Respond(context, &dtos.Response{
		Status:  http.StatusCreated,
		Message: "Weight history created successfully",
		Ok:      true,
		Data:    gin.H{"weightHistoryId": weightHistory.ID},
	})
}

func GetAllWeightHistoryByUser(context *gin.Context) {
	userID := context.GetUint("userID")

	var weightHistory []models.WeightHistory

	if result := db.DB.Where("user_id = ?", userID).Find(&weightHistory); result.Error != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Could not retrieve weight history",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	if len(weightHistory) == 0 {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusOK,
			Message: "No weight history found for this user",
			Ok:      true,
			Data:    []models.WeightHistory{},
		})
		return
	}

	utils.Respond(context, &dtos.Response{
		Status:  http.StatusOK,
		Message: "Weight history retrieved successfully",
		Ok:      true,
		Data:    weightHistory,
	})
}
