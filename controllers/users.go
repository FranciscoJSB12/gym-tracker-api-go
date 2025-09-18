package controllers

import (
	"net/http"

	"github.com/FranciscoJSB12/gym-tracker-api-go/db"
	"github.com/FranciscoJSB12/gym-tracker-api-go/models"
	"github.com/FranciscoJSB12/gym-tracker-api-go/utils"
	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "ok": false})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not sign up user", "ok": false})
		return
	}

	user.Password = hashedPassword

	result := db.DB.Create(&user)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not sign up user", "ok": false})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully", "ok": true})
}

func Login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "ok": false})
		return
	}

	enteredPassword := user.Password
	enteredEmail := user.Email

	result := db.DB.Where("email = ?", enteredEmail).First(&user)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not log user in", "ok": false})
		return
	}

	isPasswordValid := utils.CheckPasswordHash(enteredPassword, user.Password)

	if !isPasswordValid {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not log user in", "ok": false})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not log user in", "ok": false})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "ok": true, "token": token})
}
