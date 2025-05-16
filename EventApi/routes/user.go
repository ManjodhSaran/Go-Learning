package routes

import (
	"fmt"
	"net/http"

	"eventapi.com/models"
	"github.com/gin-gonic/gin"
)

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	fmt.Print(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	token, err := user.Validate()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user logged in", "data": token})
}

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
