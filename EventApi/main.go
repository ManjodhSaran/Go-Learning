package main

import (
	"net/http"
	"strconv"

	"eventapi.com/db"
	"eventapi.com/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvents)

	server.Run(":8080")
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
	}

	event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	if len(events) == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "No events found"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
	}

	event.ID = 1
	event.UserId = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "data": event})
}
