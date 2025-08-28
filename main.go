package main

import (
	"alumni-circle-api/config"
	"alumni-circle-api/handlers"
	"alumni-circle-api/repositories"
	"alumni-circle-api/routes"
	"alumni-circle-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.ConnectDB()

	alumniRepo := repositories.NewAlumniMySQLRepository(config.DB)
	alumniService := services.NewAlumniService(alumniRepo)
	alumniHandler := handlers.NewAlumniHandler(alumniService)

	eventRepo := repositories.NewEventMySQLRepository(config.DB)
	eventService := services.NewEventService(eventRepo)
	eventHandler := handlers.NewEventHandler(eventService)

	discussionRepo := repositories.NewDiscussionMySQLRepository(config.DB)
	discussionService := services.NewDiscussionService(discussionRepo)
	discussionHandler := handlers.NewDiscussionHandler(discussionService)

	routes.SetUpRoutes(router, alumniHandler, eventHandler, discussionHandler)

	router.Run(":8080")
}
