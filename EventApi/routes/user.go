package routes

import (
	"net/http"

	"eventapi.com/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	newUser := models.User{
		Email: user.Email,
		ID:    user.ID,
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created", "data": newUser})
}
