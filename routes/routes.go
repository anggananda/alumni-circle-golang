package routes

import (
	"alumni-circle-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine, alumniHandler *handlers.AlumniHandler, eventHandler *handlers.EventHandler, discussionHandler *handlers.DiscussionHandler) {
	rg := router.Group("/api/v1")
	setUpAuthRoutes(rg, alumniHandler)
	setUpPrivateRoutes(rg, alumniHandler, eventHandler, discussionHandler)
}

func setUpAuthRoutes(rg *gin.RouterGroup, alumniHandler *handlers.AlumniHandler) {
	auth := rg.Group("/auth")
	{
		auth.POST("/register", alumniHandler.Register)
		auth.POST("/login", alumniHandler.Login)
	}
}

func setUpPrivateRoutes(rg *gin.RouterGroup, alumniHandler *handlers.AlumniHandler, eventHandler *handlers.EventHandler, discussionHandler *handlers.DiscussionHandler) {
	private := rg.Group("/private")

	setUpAlumniRoutes(private, alumniHandler)
	setUpEventRoutes(private, eventHandler)
	setUpDiscussionRoutes(private, discussionHandler)
}

func setUpAlumniRoutes(rg *gin.RouterGroup, alumniHandler *handlers.AlumniHandler) {
	rg.GET("/alumni", alumniHandler.GetAllAlumni)
	rg.GET("/alumni/:IDAlumni", alumniHandler.GetAlumniByID)
	rg.DELETE("/alumni/:IDAlumni", alumniHandler.DeleteAlumni)
}

func setUpEventRoutes(rg *gin.RouterGroup, eventHandler *handlers.EventHandler) {
	rg.GET("/event", eventHandler.GetAllEvent)
	rg.GET("/event/category/:idKategori", eventHandler.GetEventByCategory)
	rg.GET("/event/:idEvent", eventHandler.GetEventByID)
}

func setUpDiscussionRoutes(rg *gin.RouterGroup, discussionHandler *handlers.DiscussionHandler) {
	rg.GET("/discussion", discussionHandler.GetAllDiscussion)
	rg.GET("/discussion/:idDiskusi", discussionHandler.GetDiscussionByID)
}
