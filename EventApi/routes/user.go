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

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	token, err := user.Validate()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user logged in", "data": gin.H{
		"id":    user.ID,
		"email": user.Email,
		"token": token,
	}})
}

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
	}

	err = user.Save()
	fmt.Print(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"data": gin.H{
			"Email": user.Email,
			"ID":    user.ID,
		},
	})
}

func getUsers(context *gin.Context) {
	var users []models.UserData
	err := models.GetAllUsers(&users)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "users fetched",
		"data":    users,
	})
}
