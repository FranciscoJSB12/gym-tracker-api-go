package controllers

import (
	"net/http"

	"github.com/FranciscoJSB12/gym-tracker-api-go/db"
	"github.com/FranciscoJSB12/gym-tracker-api-go/models"
	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	result := db.DB.Create(&user)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not sign up user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully", "id": user.ID})

}
