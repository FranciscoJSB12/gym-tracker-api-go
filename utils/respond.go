package utils

import (
	"github.com/FranciscoJSB12/gym-tracker-api-go/dtos"
	"github.com/gin-gonic/gin"
)

// Respond sends a standardized JSON response using a Response DTO.
func Respond(ctx *gin.Context, r *dtos.Response) {
	ctx.JSON(r.Status, gin.H{
		"message": r.Message,
		"ok":      r.Ok,
		"data":    r.Data,
	})
}
