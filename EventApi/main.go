package main

import (
	"eventapi.com/db"
	"eventapi.com/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisteredRoutes(server)
	server.Run(":8080")
}
