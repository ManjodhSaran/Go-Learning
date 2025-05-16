package routes

import (
	"net/http"
	"strconv"

	"eventapi.com/models"
	"github.com/gin-gonic/gin"
)

func eventRegistration(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	regs, err := models.GetDetailedRegistration(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event registrations", "error": err.Error()})
		return
	}
	// if len(regs) == 0 {
	// 	context.JSON(http.StatusNotFound, gin.H{"message": "No registrations found"})
	// 	return
	// }

	context.JSON(http.StatusOK, regs)
}

func registerForEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	_, err = models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	var eventRegistration models.EventRegistration
	eventRegistration.EventId = id

	userID := context.GetInt64("user_id")

	eventRegistration.UserId = userID

	isRegistered, err := models.CheckIsRegistered(userID, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not check registration status"})
		return
	}
	if isRegistered != 0 {
		context.JSON(http.StatusConflict, gin.H{"message": "User is already registered for this event", "data": eventRegistration})
		return
	}

	err = eventRegistration.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered for event", "data": eventRegistration})
}

func unregisterFromEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	_, err = models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	userID := context.GetInt64("user_id")
	isRegistered, err := models.CheckIsRegistered(userID, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not check registration status"})
		return
	}
	if isRegistered == 0 {
		context.JSON(http.StatusConflict, gin.H{"message": "User is not registered for this event"})
		return
	}

	err = models.DeleteRegistration(isRegistered)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not unregister from event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Unregistered from event"})
}
