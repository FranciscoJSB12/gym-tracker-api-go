package controllers

import (
	"net/http"

	"github.com/FranciscoJSB12/gym-tracker-api-go/db"
	"github.com/FranciscoJSB12/gym-tracker-api-go/dtos"
	"github.com/FranciscoJSB12/gym-tracker-api-go/models"
	"github.com/FranciscoJSB12/gym-tracker-api-go/utils"
	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Could not parse request data",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong while signing you up!",
			Ok:      false,
			Data:    nil,
		})
		return
	}
	user.Password = hashedPassword

	if result := db.DB.Create(&user); result.Error != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong while signing you up!",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	utils.Respond(context, &dtos.Response{
		Status:  http.StatusCreated,
		Message: "User signed up successfully",
		Ok:      true,
		Data:    nil,
	})
}

func Login(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Could not parse request data",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	enteredPassword := user.Password
	enteredEmail := user.Email

	if result := db.DB.Where("email = ?", enteredEmail).First(&user); result.Error != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid credentials",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	if isPasswordValid := utils.CheckPasswordHash(enteredPassword, user.Password); !isPasswordValid {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid credentials",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, int64(user.ID))
	if err != nil {
		utils.Respond(context, &dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong while logging you in!",
			Ok:      false,
			Data:    nil,
		})
		return
	}

	utils.Respond(context, &dtos.Response{
		Status:  http.StatusOK,
		Message: "Login successful!",
		Ok:      true,
		Data:    gin.H{"token": token},
	})
}
