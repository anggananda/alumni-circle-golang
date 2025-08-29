package main

import (
	"alumni-circle-api/config"
	"alumni-circle-api/handlers"
	"alumni-circle-api/repositories"
	"alumni-circle-api/routes"
	"alumni-circle-api/services"

	"github.com/gin-gonic/gin"
  "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "alumni-circle-api/docs"
)

// @title Alumni Circle API
// @version 1.0
// @description API untuk aplikasi Alumni Circle
// @termsOfService http://swagger.io/terms/

// @contact.name Dwi Angga
// @contact.url http://github.com/username
// @contact.email angga@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

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
  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	router.Run(":8080")
}
