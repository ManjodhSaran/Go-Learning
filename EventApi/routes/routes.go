package routes

import (
	"eventapi.com/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisteredRoutes(server *gin.Engine) {

	api_v1 := server.Group("/api/v1")
	{

		// Authentication routes
		auth := api_v1.Group("/auth")
		{
			auth.POST("/signup", signup)
			auth.POST("/login", login)
		}

		// Public routes
		authenticated := api_v1.Group("/")
		authenticated.Use(middlewares.Auth)

		// Events routes
		events := authenticated.Group("/events")
		{
			events.GET("/", getEvents)
			events.GET("/:id", getEvent)
			events.POST("/", createEvents)
			events.PUT("/:id", updateEvent)
			events.DELETE("/:id", deleteEvent)
		}

		// Registering for events routes
		registerEvents := authenticated.Group("/events/:id/register")
		{
			registerEvents.GET("/", eventRegistration)
			registerEvents.POST("/", registerForEvent)
			registerEvents.DELETE("/", unregisterFromEvent)
		}

		// Users routes
		users := authenticated.Group("/users")
		{
			users.GET("/", getUsers)
		}
	}
}
