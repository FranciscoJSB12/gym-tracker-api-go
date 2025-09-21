package routes

import (
	"net/http"

	"github.com/FranciscoJSB12/gym-tracker-api-go/dtos"
	"github.com/FranciscoJSB12/gym-tracker-api-go/middlewares"
	"github.com/FranciscoJSB12/gym-tracker-api-go/models"
	"github.com/FranciscoJSB12/gym-tracker-api-go/utils"
	"github.com/gin-gonic/gin"
)

func createWeightHistory(context *gin.Context) {
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

	userID := context.GetUint("userID")

	weightHistory.UserID = userID

	err := weightHistory.CreateWeightHistory()

	if err != nil {
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

func getWeightHistoryByUser(context *gin.Context) {
	userID := context.GetUint("userID")

	var weightHistory []models.WeightHistory

	err := models.GetWeightHistoryByUser(userID, &weightHistory)

	if err != nil {
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

func registerWeightHistoryRoutes(server *gin.Engine) {
	group := server.Group("/weight-histories")

	group.Use(middlewares.Authenticate)
	group.POST("/user", createWeightHistory)
	group.GET("/user", getWeightHistoryByUser)
}
